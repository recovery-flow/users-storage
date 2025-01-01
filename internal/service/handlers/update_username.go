package handlers

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/service/requests"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9._<>]+$`)

func UpdateUsername(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateUsername(r)
	if err != nil {
		logrus.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	username := req.Data.Attributes.Username

	if *username == "" && len(*username) < 3 && len(*username) > 20 {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("username is required and must be between 3 and 20 characters"))...)
		return
	}

	if !usernameRegex.MatchString(*username) {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("username can only contain letters, numbers, '.', '_', '<', and '>'"))...)
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

	user, err := server.Databaser.Users.UpdateUsername(r.Context(), userID, *username)
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user, requests.UserUpdateType))
}
