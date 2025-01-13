package handlers

import (
	"errors"
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
	"github.com/sirupsen/logrus"
)

func MemberUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewMemberUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	teamReq := req.Data.Attributes.TeamId
	roleStr := req.Data.Attributes.Role
	description := req.Data.Attributes.Description
	userForUpdate := req.Data.Attributes.UserId

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	if teamReq != chi.URLParam(r, "team_id") {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("url and body conflict"))...)
		return
	}

	teamId, err := uuid.Parse(chi.URLParam(r, "team_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if userForUpdate != chi.URLParam(r, "user_id") {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("url and body conflict"))...)
		return
	}

	updateUserId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	memberId, err := uuid.Parse(chi.URLParam(r, "member_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	team, err := server.MongoDB.Teams.FilterById(teamId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	stmt := map[string]any{}

	if roleStr != nil {
		cond := false
		for _, member := range team.Members {
			if member.ID == userID {
				if roles.CompareRolesTeam(member.Role, roles.RoleTeamAdmin) != -1 {
					cond = true
					break
				}
			}
		}
		if !cond {
			httpkit.RenderErr(w, problems.Forbidden("You don't have permissions to remove member"))
			return
		}
		stmt["role"] = roles.TeamRole(*roleStr)
	}

	if description != nil {
		stmt["description"] = *description
	}

	teamMembers, err := server.MongoDB.Teams.FilterById(teamId).Members()
	if err != nil {
		log.Errorf("Failed to get team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = teamMembers.FilterById(memberId).FilterByUserId(updateUserId).Update(r.Context(), stmt)
	if err != nil {
		log.Errorf("Failed to update member: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	membersByTeam, err := server.MongoDB.Teams.Members()
	if err != nil {
		log.Errorf("Failed to get members: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	member, err := membersByTeam.FilterById(memberId).FilterByUserId(updateUserId).Get()
	if err != nil {
		log.Errorf("Failed to get member: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewMemberResponse(member))
}
