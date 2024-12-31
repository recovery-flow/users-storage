package requests

import (
	"errors"
	"mime/multipart"
	"net/http"
)

// UpdateAvatarRequest описывает структуру для данных запроса
type UpdateAvatarRequest struct {
	File   multipart.File
	Header *multipart.FileHeader
}

// NewUpdateAvatarRequest извлекает файл из HTTP-запроса
func NewUpdateAvatarRequest(r *http.Request) (UpdateAvatarRequest, error) {
	// Проверяем Content-Type запроса
	//if err := validation.Validate(r.Header.Get("Content-Type"), validation.Required, validation.In("multipart/form-data")); err != nil {
	//	return UpdateAvatarRequest{}, errors.New("invalid content type, expected multipart/form-data")
	//}

	// Извлекаем файл из поля "avatar"
	file, header, err := r.FormFile("avatar")
	if err != nil {
		return UpdateAvatarRequest{}, errors.New("failed to retrieve file from form-data")
	}

	// Проверяем размер файла (например, 5 МБ)
	if header.Size > 5*1024*1024 {
		file.Close()
		return UpdateAvatarRequest{}, errors.New("file size exceeds the 5MB limit")
	}

	// Проверяем MIME-тип файла
	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		file.Close()
		return UpdateAvatarRequest{}, errors.New("invalid file format, only JPEG and PNG are allowed")
	}

	return UpdateAvatarRequest{File: file, Header: header}, nil
}
