package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/sirupsen/logrus"
)

func AdminDeleteAvatar(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, err = server.Cloud.User.DeleteAvatar(r.Context(), userID)
	if err != nil {
		log.WithError(err).Errorf("Failed to delete avatar from Cloudinary")
		httpkit.RenderErr(w, problems.InternalError("Failed to delete avatar"))
		return
	}

	httpkit.Render(w, http.StatusOK)
}
