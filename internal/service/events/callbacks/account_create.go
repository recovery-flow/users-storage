package callbacks

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/sirupsen/logrus"
)

type AccountCreated struct {
	Event     string    `json:"event"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Timestamp time.Time `json:"timestamp"`
}

var adjectives = []string{
	"Happy", "Lazy", "Brave", "Cool", "Smart", "Funky", "Shiny", "Swift",
	"Bright", "Witty", "Bold", "Chill", "Lucky", "Kind", "Wild", "Silly",
}
var nouns = []string{
	"Tiger", "Falcon", "Panda", "Bear", "Eagle", "Wolf", "Shark", "Dragon",
	"Phoenix", "Raven", "Lion", "Hawk", "Dolphin", "Fox", "Rabbit", "Otter",
}

func GenerateUsername() string {
	rand.Seed(time.Now().UnixNano())

	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	optionalNumber := rand.Intn(1000)

	username := fmt.Sprintf("%s%s%d", adjective, noun, optionalNumber)

	if len(username) > 20 {
		username = fmt.Sprintf("%s%s", adjective, noun)
		if len(username) > 20 {
			username = username[:20]
		}
	}

	return username
}

func CreateAccount(ctx context.Context, body []byte) error {
	var event AccountCreated
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

	username := GenerateUsername()

	userID, err := uuid.Parse(event.UserID)
	if err != nil {
		log.WithError(err).Errorf("failed to parse user id")
		return err
	}
	log.Infof("Role: %s", event.Role)

	role, err := roles.StringToRoleUser(event.Role)
	if err != nil {
		log.WithError(err).Errorf("failed to parse role")
		return err
	}

	_, err = server.MongoDB.Users.Insert(ctx, models.User{
		ID:        userID,
		Username:  username,
		Avatar:    "",
		Role:      role,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.WithError(err).Errorf("error creating user: %v", err)
		return err
	}

	log.Infof("Account created: %s", userID)
	return nil
}
