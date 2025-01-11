package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/recovery-flow/users-storage/resources"
	"github.com/sirupsen/logrus"
)

func UpdateUserFull(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateUserFull(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	username := req.Data.Attributes.Username
	title := req.Data.Attributes.Title
	status := req.Data.Attributes.Status
	avatar := req.Data.Attributes.Avatar
	bio := req.Data.Attributes.Bio
	city := req.Data.Attributes.City

	cityID, err := uuid.Parse(*city)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	user, err := server.Databaser.Users.UpdateFull(r.Context(), userID, username, title, status, avatar, bio, &cityID)
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user, resources.UserUpdateType))
}
