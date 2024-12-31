package handlers

import (
	"net/http"

	"errors"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/data/db/dbcore"
	"github.com/cifra-city/users-storage/internal/service/requests"
	"github.com/cifra-city/users-storage/resources"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateUse(r)
	if err != nil {
		logrus.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	username := req.Data.Attributes.Username
	if username == "" && len(username) < 3 && len(username) > 20 {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("username is required"))...)
		return
	}

	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := server.Logger

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	if server.Databaser.Users == nil {
		log.Warn("Users database not found")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	_, err = server.Databaser.Users.GetByUsername(r, username)
	if err == nil {
		log.Debugf("Username already exists: %v", username)
		httpkit.RenderErr(w, problems.Conflict("Username already exists"))
		return
	}
	user, err := server.Databaser.Users.Crete(r, userID, username)
	if err != nil {
		log.Errorf("Failed to create user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	httpkit.Render(w, NewUserResponse(user, requests.UserCreateType))
}

func NewUserResponse(user dbcore.User, typeOperation string) resources.User {
	var title, status, avatar, bio string
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
			},
		},
	}
}
