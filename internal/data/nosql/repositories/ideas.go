package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Ideas interface {
	Insert(ctx context.Context, idea primitive.ObjectID) error
	Remove(ctx context.Context, idea primitive.ObjectID) error
}

type ideas struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (i *ideas) Insert(ctx context.Context, idea primitive.ObjectID) error {
	update := bson.M{
		"$addToSet": bson.M{
			"ideas": idea,
		},
	}
	_, err := i.collection.UpdateOne(ctx, i.filters, update)
	if err != nil {
		return fmt.Errorf("failed to insert idea: %w", err)
	}
	return nil
}

func (i *ideas) Remove(ctx context.Context, idea primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"ideas": idea,
		},
	}
	_, err := i.collection.UpdateOne(ctx, i.filters, update)
	if err != nil {
		return fmt.Errorf("failed to delete idea: %w", err)
	}
	return nil
}
