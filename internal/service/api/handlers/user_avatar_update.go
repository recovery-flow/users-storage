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

func (h *Handlers) UserUpdateAvatar(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUploadImage(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	cloudRest, err := server.Cloud.User.SetAvatar(r.Context(), req.File, userID)
	if err != nil {
		log.WithError(err).Errorf("Failed to upload avatar to Cloudinary")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	res, err := server.MongoDB.Users.New().FilterStrict(map[string]any{
		"id": userID,
	}).UpdateOne(r.Context(), map[string]any{"avatar": cloudRest.SecureURL})
	if err != nil {
		log.WithError(err).Errorf("Failed to update avatar")
		httpkit.RenderErr(w, problems.InternalError("Failed to update avatar"))
		return
	}

	httpkit.Render(w, responses.User(*res))
}
