package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/resources"
	"github.com/sirupsen/logrus"
)

func GetTeam(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := server.Logger

	teamId, err := uuid.Parse(chi.URLParam(r, "team_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("team_id is invalid"))...)
		return
	}

	team, err := server.MongoDB.Teams.FilterById(teamId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewTeamResponse(team, resources.TeamType))
}
