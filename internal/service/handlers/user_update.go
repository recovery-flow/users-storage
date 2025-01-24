package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/reponses"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/sirupsen/logrus"
)

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
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

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	filter := make(map[string]any)
	filter["_id"] = userID

	_, err = server.MongoDB.Users.Filter(filter).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
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

	user, err := server.MongoDB.Users.Filter(filter).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, reponses.NewUserResponse(*user))
}
