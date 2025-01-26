package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/recovery-flow/users-storage/internal/service/responses"
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
	if userID.String() != req.Data.Id {
		log.Errorf("User ID does not match request user ID")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("user_id does not match request user_id"))...)
		return
	}

	filter := make(map[string]any)
	filter["_id"] = userID

	_, err = server.MongoDB.Users.New().Filter(filter).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	stmt := make(map[string]any)

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
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
