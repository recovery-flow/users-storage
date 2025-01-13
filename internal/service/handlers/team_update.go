package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/recovery-flow/users-storage/internal/service/roles"
	"github.com/recovery-flow/users-storage/resources"
	"github.com/sirupsen/logrus"
)

func TeamUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewTeamUpdate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	teamId, err := uuid.Parse(chi.URLParam(r, "team_id"))
	if err != nil {
		log.Errorf("Failed to parse team ID: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	name := req.Data.Attributes.Name
	description := req.Data.Attributes.Description

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	team, err := server.MongoDB.Teams.FilterById(teamId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	cond := false
	for _, member := range team.Members {
		if member.ID == userID {
			if roles.CompareRolesTeam(member.Role, roles.RoleTeamModer) != -1 {
				cond = true
				break
			}
		}
	}

	if !cond {
		log.Warn("User is not authorized to update team")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	stmt := map[string]any{}

	if name != nil {
		stmt["name"] = name
	}
	if description != nil {
		stmt["description"] = description
	}

	err = server.MongoDB.Teams.FilterById(teamId).Update(r.Context(), stmt)
	if err != nil {
		log.Errorf("Failed to update team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	team, err = server.MongoDB.Teams.FilterById(teamId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewTeamResponse(team, resources.TeamType))
}
