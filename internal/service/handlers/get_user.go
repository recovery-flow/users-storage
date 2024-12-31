package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/service/requests"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	logrus.Infof("Request path: %s, username param: %s", r.URL.Path, username)

	if username == "" {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("username is required"))...)
		return
	}

	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := server.Logger
	log.Infof("Getting user: %v", username)

	user, err := server.Databaser.Users.GetByUsername(r, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.Errorf("Failed to get user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user, requests.UserGetType))
}
