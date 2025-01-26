package callbacks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reportUserToUser struct {
	Event           string `json:"event"`
	OffensiveUserId string `json:"offensive_user_id"`
	DefenciveUserId string `json:"defencive_user_id"`
}

func ReportUserToUser(ctx context.Context, body []byte) error {
	var event reportUserToUser
	err := json.Unmarshal(body, &event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event body: %w", err)
	}

	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to get server from context")
		return err
	}
	log := server.Logger

	filter := make(map[string]any)
	filter["_id"] = event.OffensiveUserId

	offensiveId, err := primitive.ObjectIDFromHex(event.OffensiveUserId)
	if err != nil {
		log.WithError(err).Error("Failed to parse offensive id")
		return err
	}
	defenciveId, err := primitive.ObjectIDFromHex(event.DefenciveUserId)
	if err != nil {
		log.WithError(err).Error("Failed to parse defencive id")
		return err
	}

	err = server.MongoDB.Users.New().Filter(filter).ReportsSent().Insert(ctx, defenciveId)
	if err != nil {
		log.WithError(err).Error("Failed to add SentReport to user")
		return err
	}

	filter["_id"] = event.DefenciveUserId
	err = server.MongoDB.Users.New().Filter(filter).ReportsReceived().Insert(ctx, offensiveId)
	if err != nil {
		log.WithError(err).Error("Failed to add ReceivedReport to user")
		return err
	}
	return nil
}
