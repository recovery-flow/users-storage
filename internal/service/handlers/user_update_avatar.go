package handlers

import (
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/sirupsen/logrus"
)

func UserUpdateAvatar(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateAvatarRequest(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	defer req.File.Close()

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	yes := true
	uploadParams := uploader.UploadParams{
		Folder:       "avatars",
		PublicID:     userID.String(),
		Overwrite:    &yes,
		ResourceType: "image",
	}
	uploadResult, err := server.Storage.Upload.Upload(r.Context(), req.File, uploadParams)
	if err != nil {
		log.Errorf("Failed to upload avatar to Cloudinary: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to upload avatar"))
		return
	}

	filter := make(map[string]any)
	filter["_id"] = userID

	stmt := map[string]any{
		"avatar": uploadResult.SecureURL,
	}

	_, err = server.MongoDB.Users.New().Filter(filter).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.Errorf("Failed to update avatar URL in database: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to save avatar"))
		return
	}

	httpkit.Render(w, http.StatusOK)
}
