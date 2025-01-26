package handlers

import (
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
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

	publicID := "avatars/" + userID.String()
	_, err = server.Storage.Upload.Destroy(r.Context(), uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		log.Errorf("Failed to delete avatar from Cloudinary: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to delete avatar"))
		return
	}

	filter := map[string]any{
		"_id": userID,
	}
	update := map[string]any{
		"avatar": nil,
	}

	_, err = server.MongoDB.Users.New().Filter(filter).UpdateOne(r.Context(), update)
	if err != nil {
		log.Errorf("Failed to update user record in database: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to update user record"))
		return
	}

	httpkit.Render(w, http.StatusOK)
}
