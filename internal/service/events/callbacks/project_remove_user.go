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

type projectRemoveUser struct {
	Event     string `json:"event"`
	UserId    string `json:"user_id"`
	ProjectId string `json:"project_id"`
}

func ProjectRemoveUser(ctx context.Context, body []byte) error {
	var event projectRemoveUser
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
	filter["_id"] = event.UserId

	prjId, err := primitive.ObjectIDFromHex(event.ProjectId)
	if err != nil {
		log.WithError(err).Error("Failed to parse project id")
		return err
	}

	err = server.MongoDB.Users.New().Filter(filter).Projects().Remove(ctx, prjId)
	if err != nil {
		log.WithError(err).Error("Failed to remove user from project")
		return err
	}
	return nil
}
