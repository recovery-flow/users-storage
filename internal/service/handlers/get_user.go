package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/go-chi/chi"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	user, err := Server.Databaser.Users.GetByUsername(r, username)
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(user))
}
