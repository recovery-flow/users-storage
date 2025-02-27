package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/tokens/identity"
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	client   *redis.Client
	lifeTime time.Duration
}

func NewUsers(client *redis.Client, lifetime time.Duration) *Users {
	return &Users{
		client:   client,
		lifeTime: lifetime,
	}
}

func (u *Users) Add(ctx context.Context, user models.User) error {
	IdKey := fmt.Sprintf("user:id:%s", user.ID)
	nameKey := fmt.Sprintf("user:username:%s", user.Username)

	data := map[string]interface{}{
		"username":   user.Username,
		"role":       string(user.Role),
		"verified":   user.Verified,
		"created_at": user.CreatedAt.Time().UTC(),
	}
	if user.Avatar != nil {
		data["avatar"] = *user.Avatar
	}
	if user.TitleName != nil {
		data["title_name"] = *user.TitleName
	}
	if user.Speciality != nil {
		data["speciality"] = *user.Speciality
	}
	if user.Position != nil {
		data["position"] = *user.Position
	}
	if user.City != nil {
		data["city"] = *user.City
	}
	if user.Country != nil {
		data["country"] = *user.Country
	}
	if user.DateOfBirth != nil {
		data["date_of_birth"] = user.DateOfBirth.Time().UTC()
	}
	if user.UpdatedAt != nil {
		data["updated_at"] = user.UpdatedAt.Time().UTC()
	}

	err := u.client.HSet(ctx, IdKey, data).Err()
	if err != nil {
		return fmt.Errorf("error adding account to Redis: %w", err)
	}

	err = u.client.Set(ctx, nameKey, user.ID.String(), 0).Err()
	if err != nil {
		return fmt.Errorf("error creating email index: %w", err)
	}

	if u.lifeTime > 0 {
		keys := []string{IdKey, nameKey}
		for _, key := range keys {
			_ = u.client.Expire(ctx, key, u.lifeTime).Err()
		}
	}

	return nil
}

func (u *Users) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	IdKey := fmt.Sprintf("user:id:%s", userID)
	vals, err := u.client.HGetAll(ctx, IdKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting user from Redis: %w", err)
	}

	if len(vals) == 0 {
		return nil, fmt.Errorf("user not found, id=%s", userID)
	}

	return parseUser(userID, vals)
}

func (u *Users) GetByUsername(ctx context.Context, name string) (*models.User, error) {
	nameKey := fmt.Sprintf("user:username:%s", name)

	userID, err := u.client.Get(ctx, nameKey).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting userID by name: %w", err)
	}

	ID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("error parsing userID: %w", err)
	}

	return u.GetByID(ctx, ID)
}

func (u *Users) DeleteByID(ctx context.Context, userID uuid.UUID) error {
	key := fmt.Sprintf("user:id:%s", userID.String())

	exists, err := u.client.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("error checking user existence in Redis: %w", err)
	}

	if exists == 0 {
		return nil
	}

	username, err := u.client.HGet(ctx, key, "username").Result()
	if err != nil {
		return fmt.Errorf("error getting username: %w", err)
	}

	err = u.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	err = u.client.Del(ctx, fmt.Sprintf("user:username:%s", username)).Err()
	if err != nil {
		return fmt.Errorf("error deleting username index: %w", err)
	}

	return nil
}

func (u *Users) DeleteByUsername(ctx context.Context, name string) error {
	nameKey := fmt.Sprintf("user:username:%s", name)

	userID, err := u.client.Get(ctx, nameKey).Result()
	if err != nil {
		return fmt.Errorf("error getting userID by name: %w", err)
	}

	ID, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("error parsing userID: %w", err)
	}

	return u.DeleteByID(ctx, ID)
}

func parseUser(userID uuid.UUID, vals map[string]string) (*models.User, error) {
	role, err := identity.ParseIdentityType(vals["role"])
	if err != nil {
		return nil, fmt.Errorf("error parsing role: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339, vals["created_at"])
	if err != nil {
		return nil, fmt.Errorf("error parsing created_at: %w", err)
	}

	var updatedAt *primitive.DateTime
	if val, ok := vals["updated_at"]; ok {
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return nil, fmt.Errorf("error parsing updated_at: %w", err)
		}
		updated := primitive.NewDateTimeFromTime(t)
		updatedAt = &updated
	}

	var dateOfBirth *primitive.DateTime
	if val, ok := vals["date_of_birth"]; ok {
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return nil, fmt.Errorf("error parsing date_of_birth: %w", err)
		}
		dob := primitive.NewDateTimeFromTime(t)
		dateOfBirth = &dob
	}

	user := &models.User{
		ID:          userID,
		Username:    vals["username"],
		Role:        role,
		Verified:    vals["verified"] == "true",
		CreatedAt:   primitive.NewDateTimeFromTime(createdAt),
		UpdatedAt:   updatedAt,
		DateOfBirth: dateOfBirth,
	}

	optionalFields := map[string]**string{
		"avatar":     &user.Avatar,
		"title_name": &user.TitleName,
		"speciality": &user.Speciality,
		"position":   &user.Position,
		"city":       &user.City,
		"country":    &user.Country,
	}

	for key, field := range optionalFields {
		if val, ok := vals[key]; ok {
			*field = &val
		}
	}

	return user, nil
}
