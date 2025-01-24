package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Organizations interface {
	Insert(ctx context.Context, organization primitive.ObjectID) error
	Delete(ctx context.Context, organization primitive.ObjectID) error
}

type organizations struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (o *organizations) Insert(ctx context.Context, organization primitive.ObjectID) error {
	update := bson.M{
		"$addToSet": bson.M{
			"organizations": organization,
		},
	}
	_, err := o.collection.UpdateOne(ctx, o.filters, update)
	if err != nil {
		return fmt.Errorf("failed to insert organization: %w", err)
	}
	return nil
}

func (o *organizations) Delete(ctx context.Context, organization primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"organizations": organization,
		},
	}
	_, err := o.collection.UpdateOne(ctx, o.filters, update)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}
	return nil
}
