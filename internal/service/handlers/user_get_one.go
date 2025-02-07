package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/responses"
	"github.com/sirupsen/logrus"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to retrieve service configuration")
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	queryParams := r.URL.Query()
	filter := make(map[string]any)

	if userIdStr := queryParams.Get("user_id"); userIdStr != "" {
		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			log.WithError(err).Error("Invalid user_id format")
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"user_id": validation.Validate(userIdStr, validation.Required),
			})...)
			return
		}
		filter["_id"] = userId
	}

	if username := queryParams.Get("username"); username != "" {
		filter["username"] = username
	}

	user, err := server.MongoDB.Users.New().FilterStrict(filter).Get(r.Context())
	if err != nil {
		log.WithError(err).Errorf("Failed to get user")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.User(*user))
}
