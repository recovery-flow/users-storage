package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func UsersFilter(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	queryParams := r.URL.Query()

	resp := server.MongoDB.Users.New()

	filterStrict := make(map[string]any)
	strictParams := []string{"type", "role", "verified", "speciality", "city", "country", "ban_status"}
	for _, param := range strictParams {
		if value := queryParams.Get(param); value != "" {
			filterStrict[param] = value
		}
	}

	resp = resp.FilterStrict(filterStrict)

	filterSoft := make(map[string]any)
	softParams := []string{"username", "title_name"}
	for _, param := range softParams {
		if value := queryParams.Get(param); value != "" {
			filterSoft[param] = value
		}
	}

	resp = resp.FilterSoft(filterSoft)

	more := 0
	filterNumbers := make(map[string]any)
	numberParams := []string{"level", "points", "method_int_sort"}
	for _, param := range numberParams {
		if value := queryParams.Get(param); value != "" {
			if param == "method_int_sort" {
				method := queryParams.Get("method_int_sort")
				switch method {
				case "more":
					more = 1
				case "less":
					more = -1
				default:
					continue
				}
			}
			if parsedValue, err := strconv.Atoi(value); err == nil {
				filterNumbers[param] = parsedValue
			}
		}
	}

	switch more {
	case 0:
		resp = resp.FilterStrict(filterNumbers)
	case 1:
		resp = resp.FilterNumber(filterNumbers, true)
	default:
		resp = resp.FilterNumber(filterNumbers, false)
	}

	after := false
	filterDate := make(map[string]any)
	dateParams := []string{"created_at", "updated_at", "method_date_sort"}
	for _, param := range dateParams {
		if value := queryParams.Get(param); value != "" {
			if param == "method_date_sort" {
				method := queryParams.Get("method_date_sort")
				switch method {
				case "after":
					after = true
				default:
					continue
				}
			}
			filterDate[param] = value
		}
	}

	resp.FilterDate(filterDate, after)

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

	limit := int64(pageSize)
	skip := int64((pageNumber - 1) * pageSize)

	log.Infof("FilterStrict: %v", filterStrict)
	log.Infof("FilterSoft: %v", filterSoft)

	users, err := resp.Limit(limit).Skip(skip).Select(r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.WithError(err).Errorf("Failed to get user")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	totalUsers, err := resp.Count(r.Context())
	if err != nil {
		log.WithError(err).Errorf("Failed to count users")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	baseURL := "./public/users/search"
	response := responses.NewUsersCollectionResponse(users, baseURL, queryParams, totalUsers, int64(pageSize), int64(pageNumber))

	httpkit.Render(w, response)
}
