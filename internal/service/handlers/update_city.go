package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/service/requests"
	"github.com/cifra-city/users-storage/resources"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func UpdateCity(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateBio(r)
	if err != nil {
		logrus.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	city := req.Data.Attributes.City
	cityID, err := uuid.Parse(*city)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
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

	user, err := server.Databaser.Users.UpdateCity(r.Context(), userID, cityID)
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user, resources.UserUpdateType))
}
