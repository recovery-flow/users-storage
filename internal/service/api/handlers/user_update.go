package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/api/requests"
	"github.com/recovery-flow/users-storage/internal/service/api/responses"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) UserUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.WithError(err).Error("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewUserUpdate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	fields := map[string]any{
		"username":      req.Data.Attributes.Username,
		"title_name":    req.Data.Attributes.TitleName,
		"speciality":    req.Data.Attributes.Speciality,
		"position":      req.Data.Attributes.Position,
		"city":          req.Data.Attributes.City,
		"country":       req.Data.Attributes.Country,
		"date_of_birth": req.Data.Attributes.DateOfBirth,
	}

	stmt := make(map[string]any)
	for key, value := range fields {
		if value != nil {
			stmt[key] = value
		}
	}

	user, err := server.MongoDB.Users.New().FilterStrict(map[string]any{
		"_id": userID,
	}).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Errorf("Failed to update username")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
