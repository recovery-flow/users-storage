package handlers

import (
	"errors"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func UsersSearch(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	log := server.Logger
	queryParams := r.URL.Query()

	filter := make(map[string]any)
	for key, values := range queryParams {
		if len(values) > 0 && key != "page[size]" && key != "page[number]" {
			filter[key] = values[0]
		}
	}

	pageSize := 10
	pageNumber := 1

	if size := queryParams.Get("page[size]"); size != "" {
		if parsedSize, err := strconv.Atoi(size); err == nil && parsedSize > 0 {
			pageSize = parsedSize
		}
	}

	if number := queryParams.Get("page[number]"); number != "" {
		if parsedNumber, err := strconv.Atoi(number); err == nil && parsedNumber > 0 {
			pageNumber = parsedNumber
		}
	}

	if len(filter) == 0 {
		httpkit.RenderErr(w, problems.BadRequest(
			validation.Errors{
				"query": errors.New("query parameters are required"),
			})...)
		return
	}

	limit := int64(pageSize)
	skip := int64((pageNumber - 1) * pageSize)

	users, err := server.MongoDB.Users.New().FilterCoincidence(filter).Limit(limit).Skip(skip).Select(r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.WithError(err).Errorf("Failed to get user")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	totalUsers, err := server.MongoDB.Users.New().Filter(filter).Count(r.Context())
	if err != nil {
		log.WithError(err).Errorf("Failed to count users")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	baseURL := "./public/users/search"
	response := responses.NewUsersCollectionResponse(users, baseURL, queryParams, totalUsers, int64(pageSize), int64(pageNumber))

	httpkit.Render(w, response)
}
