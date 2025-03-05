package callbacks

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/users-storage/internal/service"
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/recovery-flow/users-storage/internal/service/infra/events/evebody"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func AccountCreate(ctx context.Context, svc *service.Service, event evebody.AccountCreated) error {
	username := GenerateUsername()

	userID, err := uuid.Parse(event.AccountID)
	if err != nil {
		svc.Log.WithError(err).Errorf("failed to parse user id")
		return err
	}

	_, err = svc.Domain.CreateUser(ctx, models.User{
		ID:        userID,
		Username:  username,
		Verified:  false,
		CreatedAt: primitive.NewDateTimeFromTime(event.Timestamp),
	})
	if err != nil {
		svc.Log.WithError(err).Errorf("error creating user: %v", err)
		return err
	}

	svc.Log.Infof("Account created: %s", userID)
	return nil
}
