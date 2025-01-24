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

type projectAddUser struct {
	Event     string `json:"event"`
	UserId    string `json:"user_id"`
	ProjectId string `json:"project_id"`
}

func ProjectAddUser(ctx context.Context, body []byte) error {
	var event projectAddUser
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

	prjId, err := primitive.ObjectIDFromHex(event.ProjectId)
	if err != nil {
		log.WithError(err).Error("Failed to parse project id")
		return err
	}

	err = server.MongoDB.Users.New().Filter(filter).Projects().Insert(ctx, prjId)
	if err != nil {
		log.WithError(err).Error("Failed to add user to organization")
		return err
	}
	return nil
}
