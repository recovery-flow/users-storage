package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/resources"
	"github.com/sirupsen/logrus"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
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

	user, err := server.MongoDB.Users.FilterByUsername(username).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user))
}

func NewUserResponse(user models.User) resources.User {
	return resources.User{
		Data: resources.UserData{
			Id:   user.ID.String(),
			Type: resources.UserType,
			Attributes: resources.UserDataAttributes{
				Username:  user.Username,
				Avatar:    "",
				Role:      user.Role,
				CreatedAt: user.CreatedAt,
			},
		},
	}
}
