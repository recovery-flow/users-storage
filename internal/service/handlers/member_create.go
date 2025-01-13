package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/recovery-flow/users-storage/internal/service/roles"
	"github.com/recovery-flow/users-storage/resources"
	"github.com/sirupsen/logrus"
)

func MemberCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
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

	req, err := requests.NewMemberCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	teamId, err := uuid.Parse(chi.URLParam(r, "team_id"))
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

	user := req.Data.Attributes.UserId
	roleStr := req.Data.Attributes.Role
	description := req.Data.Attributes.Description

	userId, err := uuid.Parse(user)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	role := roles.TeamRole(roleStr)

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
		log.Warn("User is not authorized to update team")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	membersByTeam, err := server.MongoDB.Teams.Members()
	if err != nil {
		log.Errorf("Failed to get members: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = membersByTeam.Insert(r.Context(), models.Member{
		ID:        uuid.New(),
		UserId:    userId,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Errorf("Failed to add member: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if description != nil {
		err = membersByTeam.FilterByUserId(userId).Update(r.Context(), map[string]interface{}{"description": description})
		if err != nil {
			log.Errorf("Failed to update member: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	membersByTeam, err = server.MongoDB.Teams.FilterById(teamId).Members()
	if err != nil {
		log.Errorf("Failed to get members: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	members, err := membersByTeam.FilterByUserId(userId).Get()
	if err != nil {
		log.Errorf("Failed to get member: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewMemberResponse(members))
}

func NewMemberResponse(member models.Member) resources.Member {
	return resources.Member{
		Data: resources.MemberData{
			Type: resources.MemberType,
			Id:   member.ID.String(),
			Attributes: resources.MemberDataAttributes{
				Role:        string(member.Role),
				Description: member.Description,
				CreatedAt:   member.CreatedAt,
			},
			Relationships: resources.NewMemberDataRelationships(),
		},
	}
}
