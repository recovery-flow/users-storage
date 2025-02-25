package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/service/api/responses"
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
)

func (h *Handlers) UserGet(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := h.Domain.GetUser(r.Context(), models.RequestQuery{
		Filters: map[string]models.QueryFilter{"_id": {Type: "strict", Method: "$eq", Value: userID}},
	})
	if err != nil {
		h.Log.WithError(err).Errorf("Failed to get user")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
