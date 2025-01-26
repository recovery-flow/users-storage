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

type unbanUser struct {
	Event  string `json:"event"`
	UserID string `json:"user_id"`
}

func UnbanUser(ctx context.Context, body []byte) error {
	var event unbanUser
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

	userId, err := primitive.ObjectIDFromHex(event.UserID)
	if err != nil {
		log.WithError(err).Error("Failed to parse user id")
		return err
	}

	filter := make(map[string]any)
	filter["_id"] = userId

	stmt := make(map[string]any)
	stmt["ban_id"] = nil

	_, err = server.MongoDB.Users.New().Filter(filter).UpdateOne(ctx, stmt)
	if err != nil {
		log.WithError(err).Error("Failed to delete Ban from user")
		return err
	}

	return nil
}
