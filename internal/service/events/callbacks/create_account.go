package callbacks

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AccountCreated struct {
	Event     string    `json:"event"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
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
		logrus.Fatalf("failed to get server from context: %v", err)
		return err
	}
	log := server.Logger

	username := GenerateUsername()

	userID, err := uuid.Parse(event.UserID)
	if err != nil {
		log.Errorf("failed to parse user id: %v", err)
		return err
	}

	user, err := server.Databaser.Users.Crete(ctx, userID, username)
	if err != nil {
		log.Errorf("error creating user: %v", err)
		return err
	}

	log.Infof("Account created: %s", user.ID)
	return nil
}
