package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Projects interface {
	Insert(ctx context.Context, project primitive.ObjectID) error
	Remove(ctx context.Context, project primitive.ObjectID) error
}

type projects struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (p *projects) Insert(ctx context.Context, project primitive.ObjectID) error {
	update := bson.M{
		"$addToSet": bson.M{
			"projects": project,
		},
	}
	_, err := p.collection.UpdateOne(ctx, p.filters, update)
	if err != nil {
		return fmt.Errorf("failed to insert project: %w", err)
	}
	return nil
}

func (p *projects) Remove(ctx context.Context, project primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"projects": project,
		},
	}
	_, err := p.collection.UpdateOne(ctx, p.filters, update)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}
