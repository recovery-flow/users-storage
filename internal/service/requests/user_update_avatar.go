package requests

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

type UpdateAvatarRequest struct {
	File   multipart.File
	Header *multipart.FileHeader
}

func NewUploadImage(r *http.Request) (UpdateAvatarRequest, error) {
	file, header, err := r.FormFile("avatar")
	if err != nil {
		return UpdateAvatarRequest{}, fmt.Errorf("failed to retrieve file from form-data")
	}

	if header.Size > 5*1024*1024 {
		err := file.Close()
		if err != nil {
			return UpdateAvatarRequest{}, err
		}
		return UpdateAvatarRequest{}, fmt.Errorf("file size exceeds the 5MB limit")
	}

	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		err := file.Close()
		if err != nil {
			return UpdateAvatarRequest{}, err
		}
		return UpdateAvatarRequest{}, fmt.Errorf("invalid file format, only JPEG and PNG are allowed")
	}

	return UpdateAvatarRequest{File: file, Header: header}, nil
}
