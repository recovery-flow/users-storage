package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/sirupsen/logrus"
)

func UserDeleteAvatar(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	_, err = server.Cloud.User.DeleteAvatar(r.Context(), userID)
	if err != nil {
		log.WithError(err).Errorf("Failed to delete avatar from Cloudinary")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	_, err = server.MongoDB.Users.New().FilterStrict(map[string]any{
		"id": userID,
	}).UpdateOne(r.Context(), map[string]any{"avatar": nil})
	if err != nil {
		log.WithError(err).Errorf("Failed to update avatar")
		httpkit.RenderErr(w, problems.InternalError("Failed to update avatar"))
		return
	}

	httpkit.Render(w, http.StatusOK)
}
