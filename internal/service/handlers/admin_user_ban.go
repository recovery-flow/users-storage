package handlers

import (
	"fmt"
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
	"github.com/sirupsen/logrus"
)

func AdminUserBan(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUserBan(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	adminID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	userID := req.Data.Id
	term := req.Data.Attributes.Term
	sort := req.Data.Attributes.Sort
	desc := req.Data.Attributes.Desc

	filter := make(map[string]any)
	filter["_id"] = userID

	_, err = server.MongoDB.Users.New().Filter(filter).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	ban := make(map[string]any)

	endDate := time.Now().UTC().AddDate(0, 0, int(term))
	sortOfBan := []models.BanSort{models.CommentsBan, models.ActivityBan, models.PermanentBan}

	validSort := false
	for _, s := range sortOfBan {
		if string(s) == sort {
			ban["sort"] = sort
			validSort = true
			break
		}
	}

	if !validSort {
		log.Errorf("Invalid sort of ban: %v", sort)
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("invalid sort of ban: %v", sort))...)
		return
	}

	ban["banned"] = true
	ban["start"] = time.Now().UTC()
	ban["end"] = endDate
	ban["desc"] = desc
	ban["initiator_id"] = adminID

	err = server.MongoDB.Users.New().Filter(filter).Accessibility().UpdateOne(r.Context(), ban)
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
}
