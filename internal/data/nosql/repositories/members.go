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
	Insert(ctx context.Context, member models.Member) error
	Delete(ctx context.Context) (int64, error)
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Member, error)
	Get() (models.Member, error)

	FilterById(id uuid.UUID) Members
	FilterByUserId(userId uuid.UUID) Members
	FilterByRole(role roles.TeamRole) Members
	FilterByCreatedAt(from, to time.Time) Members

	Update(ctx context.Context, fields map[string]any) error

	SortBy(field string, ascending bool) Members
	Limit(limit int64) Members
	Skip(skip int64) Members
}

type members struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (m *members) Insert(ctx context.Context, member models.Member) error {
	member.ID = uuid.New()
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()

	_, err := m.collection.InsertOne(ctx, bson.M{"_id": member.ID, "members": member})
	if err != nil {
		return fmt.Errorf("failed to add member: %w", err)
	}
	return nil
}

func (m *members) Delete(ctx context.Context) (int64, error) {
	result, err := m.collection.DeleteMany(ctx, m.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to delete member: %w", err)
	}
	return result.DeletedCount, nil
}

func (m *members) Count(ctx context.Context) (int64, error) {
	count, err := m.collection.CountDocuments(ctx, m.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count members: %w", err)
	}
	return count, nil
}

func (m *members) Select(ctx context.Context) ([]models.Member, error) {
	param := options.Find().SetSort(m.sort).SetLimit(m.limit).SetSkip(m.skip)
	cursor, err := m.collection.Find(ctx, m.filters, param)
	if err != nil {
		return nil, fmt.Errorf("failed to find teams: %w", err)
	}
	defer cursor.Close(ctx)

	var members []models.Member
	if err := cursor.All(ctx, &members); err != nil {
		return nil, fmt.Errorf("failed to decode teams: %w", err)
	}
	return members, nil
}

func (m *members) Get() (models.Member, error) {
	var member models.Member
	err := m.collection.FindOne(context.Background(), m.filters).Decode(&member)
	if err != nil {
		return models.Member{}, fmt.Errorf("failed to get member: %w", err)
	}
	return member, nil
}

func (m *members) FilterById(id uuid.UUID) Members {
	m.filters["members"].(bson.M)["$elemMatch"].(bson.M)["_id"] = id
	return m
}

func (m *members) FilterByUserId(userId uuid.UUID) Members {
	m.filters["members"].(bson.M)["$elemMatch"].(bson.M)["user_id"] = userId
	return m
}

func (m *members) FilterByRole(role roles.TeamRole) Members {
	m.filters["members"].(bson.M)["$elemMatch"].(bson.M)["role"] = role
	return m
}

func (m *members) FilterByCreatedAt(from, to time.Time) Members {
	if m.filters["members"] == nil {
		m.filters["members"] = bson.M{
			"$elemMatch": bson.M{
				"created_at": bson.M{"$gte": from, "$lte": to},
			},
		}
	} else {
		elemMatch := m.filters["members"].(bson.M)["$elemMatch"].(bson.M)
		elemMatch["created_at"] = bson.M{"$gte": from, "$lte": to}
	}
	return m
}

func (m *members) Update(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	validFields := map[string]bool{
		"role":        true,
		"description": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields["members.$."+key] = value
		}
	}

	updateFields["members.$.updated_at"] = time.Now()

	if len(updateFields) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	filter := bson.M{
		"_id": m.filters["_id"],
	}

	if m.filters["members.user_id"] != nil {
		filter["members.user_id"] = m.filters["members.user_id"]
	}

	update := bson.M{"$set": updateFields}

	result, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update member: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no member found with the given criteria")
	}
	return nil
}

func (m *members) SortBy(field string, ascending bool) Members {
	order := 1
	if !ascending {
		order = -1
	}

	m.sort = bson.D{{field, order}}
	return m
}

func (m *members) Limit(limit int64) Members {
	m.limit = limit
	return m
}

func (m *members) Skip(skip int64) Members {
	m.skip = skip
	return m
}
