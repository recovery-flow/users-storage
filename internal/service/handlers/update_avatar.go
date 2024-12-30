package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/service/requests"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	// Получаем запрос
	req, err := requests.NewUpdateAvatarRequest(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	defer req.File.Close()

	service, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		service.Logger.Warn("UserID not found in context")
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
	uploadResult, err := service.Storage.Upload.Upload(r.Context(), req.File, uploadParams)
	if err != nil {
		service.Logger.Errorf("Failed to upload avatar to Cloudinary: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to upload avatar"))
		return
	}

	// Сохраняем URL аватара в базе данных
	_, err = service.Databaser.Users.UpdateAvatar(r, userID, &uploadResult.SecureURL)
	if err != nil {
		service.Logger.Errorf("Failed to update avatar URL in database: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to save avatar"))
		return
	}

	response := map[string]string{"avatar_url": uploadResult.SecureURL}
	httpkit.Render(w, response)
}
