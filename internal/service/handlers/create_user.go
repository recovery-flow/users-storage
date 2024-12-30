package handlers

import (
	"net/http"

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

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	user, err := Server.Databaser.Users.Crete(r, userID, req.Data.Attributes.Username)
	if err != nil {
		log.Errorf("Failed to create user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user))
}

func NewUserResponse(user dbcore.User) resources.User {
	return resources.User{
		Data: resources.UserData{
			Type: "user",
			Attributes: resources.UserDataAttributes{
				Id:       user.ID.String(),
				Username: user.Username,
				Title:    user.Title.String,
				Status:   user.Status.String,
				Avatar:   user.Avatar.String,
				Bio:      user.Bio.String,
			},
		},
	}
}
