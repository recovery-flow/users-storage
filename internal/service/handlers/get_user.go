package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/data/sql/repositories/sqlcore"
	"github.com/cifra-city/users-storage/resources"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := server.Logger

	username := chi.URLParam(r, "username")
	if username == "" {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("username is required"))...)
		return
	}

	log.Infof("Getting user: %v", username)

	user, err := server.Databaser.Users.GetByUsername(r.Context(), username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.Errorf("Failed to get user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user, resources.UserGetType))
}

func NewUserResponse(user sqlcore.User, typeOperation string) resources.User {
	var title, status, avatar, bio string
	var city uuid.NullUUID
	if user.Title.Valid {
		title = user.Title.String
	}
	if user.Status.Valid {
		status = user.Status.String
	}
	if user.Avatar.Valid {
		avatar = user.Avatar.String
	}
	if user.Bio.Valid {
		bio = user.Bio.String
	}
	if user.City.Valid {
		city = user.City
	}
	return resources.User{
		Data: resources.UserData{
			Type: typeOperation,
			Attributes: resources.UserDataAttributes{
				Id:       user.ID.String(),
				Username: user.Username,
				Title:    title,
				Status:   status,
				Avatar:   avatar,
				Bio:      bio,
				City:     city.UUID.String(),
			},
		},
	}
}
