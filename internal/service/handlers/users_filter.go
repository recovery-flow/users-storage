package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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

	filterStrict := make(map[string]any)
	strictParams := []string{"type", "role", "verified", "speciality", "city", "country", "ban_status"}
	for _, param := range strictParams {
		if value := queryParams.Get(param); value != "" {
			filterStrict[param] = value
		}
	}

	// Мягкие фильтры (поиск по совпадению)
	filterSoft := make(map[string]any)
	softParams := []string{"username", "title_name"}
	for _, param := range softParams {
		if value := queryParams.Get(param); value != "" {
			filterSoft[param] = bson.M{"$regex": value, "$options": "i"} // Поиск без учета регистра
		}
	}

	// Числовые фильтры (конвертируем в int)
	more := false
	filterNumbers := make(map[string]any)
	numberParams := []string{"level", "points", "more"}
	for _, param := range numberParams {
		if value := queryParams.Get(param); value != "" {
			if param == "more" {
				more = true
				continue
			}
			if parsedValue, err := strconv.Atoi(value); err == nil {
				filterNumbers[param] = parsedValue
			}
		}
	}

	// Фильтр по дате
	after := false
	filterDate := make(map[string]any)
	dateParams := []string{"created_at", "updated_at", "after"}
	for _, param := range dateParams {
		if value := queryParams.Get(param); value != "" {
			if param == "after" {
				after = true
				continue
			}
			if parsedTime, err := time.Parse(time.RFC3339, value); err == nil {
				filterDate[param] = bson.M{"$gte": parsedTime}
			}
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

	limit := int64(pageSize)
	skip := int64((pageNumber - 1) * pageSize)

	users, err := server.MongoDB.Users.New().
		FilterStrict(filterStrict).
		FilterSoft(filterSoft).
		FilterNumber(filterNumbers, more).
		FilterDate(filterDate, after).
		Limit(limit).Skip(skip).Select(r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.WithError(err).Errorf("Failed to get user")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	totalUsers, err := server.MongoDB.Users.New().
		FilterStrict(filterStrict).
		FilterSoft(filterSoft).
		FilterNumber(filterNumbers, more).
		FilterDate(filterDate, after).
		Count(r.Context())
	if err != nil {
		log.WithError(err).Errorf("Failed to count users")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	baseURL := "./public/users/search"
	response := responses.NewUsersCollectionResponse(users, baseURL, queryParams, totalUsers, int64(pageSize), int64(pageNumber))

	httpkit.Render(w, response)
}
