package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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
		logrus.WithError(err).Error("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError())
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
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}
	if userID.String() != req.Data.Id {
		log.WithError(err).Errorf("User ID does not match request user ID")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"id": validation.NewError("id", "User ID does not match request user_id"),
		})...)
		return
	}

	filter := make(map[string]any)
	filter["_id"] = userID

	_, err = server.MongoDB.Users.New().Filter(filter).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to update username")
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
		log.WithError(err).Error("Failed to update username")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
