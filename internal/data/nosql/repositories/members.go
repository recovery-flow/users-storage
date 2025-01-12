package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/internal/service/roles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Members interface {
	Add(ctx context.Context, userId uuid.UUID, role roles.TeamRole, description string) error
	Delete(ctx context.Context, userId uuid.UUID) error
	Update(ctx context.Context, userId uuid.UUID, role roles.TeamRole, description string) error
	Select(ctx context.Context) ([]models.Member, error)
	Get(ctx context.Context, userId uuid.UUID) (models.Member, error)
}

type members struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	teamId     uuid.UUID // Привязка к текущей команде
}

func (t *members) AddMember(ctx context.Context, teamId, userId uuid.UUID, role roles.TeamRole, description string) (models.Team, error) {
	filter := bson.M{"_id": teamId}
	update := bson.M{
		"$push": bson.M{
			"members": models.Member{
				ID:          uuid.New(),
				UserId:      userId,
				Role:        role,
				Description: description,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
		"$set": bson.M{"updated_at": time.Now()},
	}

	result := t.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var team models.Team
	if err := result.Decode(&team); err != nil {
		return models.Team{}, fmt.Errorf("failed to add member: %w", err)
	}
	return team, nil
}

func (t *members) DeleteMember(ctx context.Context, teamId, userId uuid.UUID) (models.Team, error) {
	filter := bson.M{"_id": teamId}
	update := bson.M{
		"$pull": bson.M{
			"members": bson.M{"user_id": userId},
		},
		"$set": bson.M{"updated_at": time.Now()},
	}

	result := t.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var team models.Team
	if err := result.Decode(&team); err != nil {
		return models.Team{}, fmt.Errorf("failed to delete member: %w", err)
	}
	return team, nil
}

func (t *members) UpdateMember(ctx context.Context, teamId, userId uuid.UUID, role roles.TeamRole, description string) (int64, error) {
	filter := bson.M{"_id": teamId, "members.user_id": userId}
	update := bson.M{
		"$set": bson.M{
			"members.$.role":        role,
			"members.$.description": description,
			"members.$.updated_at":  time.Now(),
			"updated_at":            time.Now(),
		},
	}

	result, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update member: %w", err)
	}
	return result.ModifiedCount, nil
}

func (t *members) SelectMembers(ctx context.Context, teamId uuid.UUID) ([]models.Member, error) {
	filter := bson.M{"_id": teamId}
	var team models.Team
	err := t.collection.FindOne(ctx, filter).Decode(&team)
	if err != nil {
		return nil, fmt.Errorf("failed to find team: %w", err)
	}
	return team.Members, nil
}

func (t *members) GetMember(ctx context.Context, teamId, userId uuid.UUID) (models.Member, error) {
	filter := bson.M{"_id": teamId}
	var team models.Team
	err := t.collection.FindOne(ctx, filter).Decode(&team)
	if err != nil {
		return models.Member{}, fmt.Errorf("failed to find team: %w", err)
	}

	for _, member := range team.Members {
		if member.UserId == userId {
			return member, nil
		}
	}
	return models.Member{}, fmt.Errorf("member not found in team")
}
