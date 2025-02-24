package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/service/api/responses"
	"github.com/recovery-flow/users-storage/internal/service/domain"
	"github.com/recovery-flow/users-storage/internal/service/infra/repositories/mongodb"
)

func (h *Handlers) UserGet(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := h.Domain.GetUser(r.Context(), domain.RequestQuery{
		Filters: map[string]mongodb.QueryFilter{"_id": {Type: "strict", Method: "$eq", Value: userId}},
	})
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to get user")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
