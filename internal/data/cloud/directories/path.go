package directories

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type User interface {
	SetAvatar(ctx context.Context, file multipart.File, Id uuid.UUID) (*uploader.UploadResult, error)
	DeleteAvatar(ctx context.Context, Id uuid.UUID) (*uploader.DestroyResult, error)
}

type cloud struct {
	Storage *cloudinary.Cloudinary
}

func NewCloud(cloudName, APIKey, APISecret string) (User, error) {
	// URL connect format: cloudinary://<API_KEY>:<API_SECRET>@<CLOUD_NAME>
	cld, err := cloudinary.NewFromParams(cloudName, APIKey, APISecret)
	if err != nil {
		return nil, err
	}
	return &cloud{
		Storage: cld,
	}, nil
}

func (c *cloud) SetAvatar(ctx context.Context, file multipart.File, Id uuid.UUID) (*uploader.UploadResult, error) {
	yes := true
	res, err := c.Storage.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:       "users/avatars",
		PublicID:     Id.String(),
		Overwrite:    &yes,
		ResourceType: "image",
	})
	return res, err
}

func (c *cloud) DeleteAvatar(ctx context.Context, Id uuid.UUID) (*uploader.DestroyResult, error) {
	res, err := c.Storage.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: Id.String(),
	})
	return res, err
}
