package handlers

import (
	"net/http"

	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/service/domain"
	"github.com/recovery-flow/users-storage/internal/service/infra/repositories/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handlers) UsersFilter(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	// Инициализируем карту фильтров
	filters := make(map[string]mongodb.QueryFilter)

	// Если передан параметр "username", можно задать тип поиска.
	if username := q.Get("username"); username != "" {
		matchType := q.Get("username_match")
		if matchType == "exact" {
			filters["username"] = mongodb.QueryFilter{
				Type:   "strict",
				Method: "eq",
				Value:  username,
			}
		} else {
			filters["username"] = mongodb.QueryFilter{
				Type:   "soft",
				Method: "regex",
				Value:  username,
			}
		}
	}

	if role := q.Get("role"); role != "" {
		filters["role"] = mongodb.QueryFilter{
			Type:   "strict",
			Method: "eq",
			Value:  role,
		}
	}

	if verified := q.Get("verified"); verified != "" {
		filters["verified"] = mongodb.QueryFilter{
			Type:   "strict",
			Method: "eq",
			Value:  verified,
		}
	}

	if id := q.Get("id"); id != "" {
		filters["id"] = mongodb.QueryFilter{
			Type:   "strict",
			Method: "eq",
			Value:  id,
		}
	}

	if title := q.Get("title"); title != "" {
		matchType := q.Get("title_match")
		if matchType == "exact" {
			filters["title"] = mongodb.QueryFilter{
				Type:   "strict",
				Method: "eq",
				Value:  title,
			}
		}
		if matchType == "soft" {
			filters["title"] = mongodb.QueryFilter{
				Type:   "soft",
				Method: "regex",
				Value:  title,
			}
		}
	}

	if speciality := q.Get("speciality"); speciality != "" {
		filters["speciality"] = mongodb.QueryFilter{
			Type:   "strict",
			Method: "eq",
			Value:  speciality,
		}
	}

	if position := q.Get("position"); position != "" {
		filters["position"] = mongodb.QueryFilter{
			Type:   "strict",
			Method: "eq",
			Value:  position,
		}
	}

	if city := q.Get("city"); city != "" {
		filters["city"] = mongodb.QueryFilter{
			Type:   "strict",
			Method: "eq",
			Value:  city,
		}
	}

	if country := q.Get("country"); country != "" {
		filters["country"] = mongodb.QueryFilter{
			Type:   "strict",
			Method: "eq",
			Value:  country,
		}
	}

	// Обработка date_of_birth
	if dob := q.Get("date_of_birth"); dob != "" {
		filters["date_of_birth"] = mongodb.QueryFilter{
			Type:   "date",
			Method: "eq",
			Value:  dob,
		}
	} else {
		dobFrom := q.Get("date_of_birth_from")
		dobTo := q.Get("date_of_birth_to")
		if dobFrom != "" || dobTo != "" {
			rangeQuery := bson.M{}
			if dobFrom != "" {
				rangeQuery["$gte"] = dobFrom
			}
			if dobTo != "" {
				rangeQuery["$lte"] = dobTo
			}
			filters["date_of_birth"] = mongodb.QueryFilter{
				Type:   "date",
				Method: "range",
				Value:  rangeQuery,
			}
		}
	}

	// Обработка updated_at
	if updatedAt := q.Get("updated_at"); updatedAt != "" {
		filters["updated_at"] = mongodb.QueryFilter{
			Type:   "date",
			Method: "eq",
			Value:  updatedAt,
		}
	} else {
		updatedFrom := q.Get("updated_at_from")
		updatedTo := q.Get("updated_at_to")
		if updatedFrom != "" || updatedTo != "" {
			rangeQuery := bson.M{}
			if updatedFrom != "" {
				rangeQuery["$gte"] = updatedFrom
			}
			if updatedTo != "" {
				rangeQuery["$lte"] = updatedTo
			}
			filters["updated_at"] = mongodb.QueryFilter{
				Type:   "date",
				Method: "range",
				Value:  rangeQuery,
			}
		}
	}

	if createdAt := q.Get("created_at"); createdAt != "" {
		filters["created_at"] = mongodb.QueryFilter{
			Type:   "date",
			Method: "eq",
			Value:  createdAt,
		}
	} else {
		createdFrom := q.Get("created_at_from")
		createdTo := q.Get("created_at_to")
		if createdFrom != "" || createdTo != "" {
			rangeQuery := bson.M{}
			if createdFrom != "" {
				rangeQuery["$gte"] = createdFrom
			}
			if createdTo != "" {
				rangeQuery["$lte"] = createdTo
			}
			filters["created_at"] = mongodb.QueryFilter{
				Type:   "date",
				Method: "range",
				Value:  rangeQuery,
			}
		}
	}

	rq := domain.RequestQuery{
		Filters: filters,
	}

	response, err := h.Domain.SelectUsers(r.Context(), rq)
	if err != nil {
		h.Log.WithError(err).Error("Failed to get users")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, response)
}
