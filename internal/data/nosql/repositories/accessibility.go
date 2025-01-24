package repositories

import (
	"context"
	"fmt"

	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Accessibility interface {
	Get(ctx context.Context) (*models.Accessibility, error)
}

type accessibility struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (a *accessibility) Get(ctx context.Context) (*models.Accessibility, error) {
	var access models.Accessibility
	err := a.collection.FindOne(ctx, a.filters).Decode(&access)
	if err != nil {
		return nil, fmt.Errorf("failed to find accessability: %w", err)
	}
	return &access, nil
}

func (a *accessibility) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	if a.filters == nil || len(a.filters) == 0 {
		return fmt.Errorf("no filters set for accessability update")
	}

	validFields := map[string]bool{
		"banned":       true,
		"start":        true,
		"end":          true,
		"sort":         true,
		"initiator_id": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] && value != nil {
			updateFields["accessibility."+key] = value
		}
	}

	if len(updateFields) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	update := bson.M{"$set": updateFields}
	_, err := a.collection.UpdateOne(ctx, a.filters, update)
	if err != nil {
		return fmt.Errorf("failed to update accessibility: %w", err)
	}

	return nil
}
