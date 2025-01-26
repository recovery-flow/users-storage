package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/recovery-flow/users-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
)

func AdminUserUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUserUpdate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	username := req.Data.Attributes.Username
	description := req.Data.Attributes.Description
	role := req.Data.Attributes.Role

	userId, err := uuid.Parse(req.Data.Id)
	if err != nil {
		log.WithError(err).Errorf("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	filter := make(map[string]any)
	filter["_id"] = userId

	_, err = server.MongoDB.Users.New().Filter(filter).Get(r.Context())
	if err != nil {
		log.WithError(err).Errorf("Failed to update username")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	stmt := map[string]any{}

	if username != nil {
		stmt["username"] = username
	}
	if description != nil {
		stmt["description"] = description
	}
	if role != nil {
		stmt["role"] = role
	}

	user, err := server.MongoDB.Users.New().Filter(filter).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Errorf("Failed to update username")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
