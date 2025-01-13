package handlers

import (
	"net/http"
	"time"

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

func TeamCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewTeamCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	name := req.Data.Attributes.Name
	descr := req.Data.Attributes.Description

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	admin := models.Member{
		UserId:      userID,
		Role:        roles.RoleTeamOwner,
		Description: "Team Owner",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	teamId := uuid.New()
	err = server.MongoDB.Teams.Insert(r.Context(), models.Team{
		ID:          teamId,
		Name:        name,
		Description: descr,
		Members:     []models.Member{admin}, // Исправлено на срез
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		log.Errorf("Failed to create team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	team, err := server.MongoDB.Teams.FilterById(teamId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to create team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewTeamResponse(team, resources.TeamCreateType))
}
