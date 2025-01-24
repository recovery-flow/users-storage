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

type ideaAddUser struct {
	Event  string `json:"event"`
	UserId string `json:"user_id"`
	IdeaId string `json:"idea_id"`
}

func IdeaAddUser(ctx context.Context, body []byte) error {
	var event ideaAddUser
	err := json.Unmarshal(body, &event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event body: %w", err)
	}

	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
		return err
	}
	log := server.Logger

	filter := make(map[string]any)
	filter["_id"] = event.UserId

	Id, err := primitive.ObjectIDFromHex(event.IdeaId)
	if err != nil {
		log.WithError(err).Error("Failed to parse idea id")
		return err
	}

	err = server.MongoDB.Users.New().Filter(filter).Ideas().Insert(ctx, Id)
	if err != nil {
		log.WithError(err).Error("Failed to add user to idea")
		return err
	}
	return nil
}
