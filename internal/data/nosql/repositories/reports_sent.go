package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportsSent interface {
	Insert(ctx context.Context, report primitive.ObjectID) error
	Remove(ctx context.Context, report primitive.ObjectID) error
}

type reportsSent struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (r *reportsSent) Insert(ctx context.Context, report primitive.ObjectID) error {
	update := bson.M{
		"$addToSet": bson.M{
			"report_sent": report,
		},
	}
	_, err := r.collection.UpdateOne(ctx, r.filters, update)
	if err != nil {
		return fmt.Errorf("failed to insert report: %w", err)
	}
	return nil
}

func (r *reportsSent) Remove(ctx context.Context, report primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"report_sent": report,
		},
	}
	_, err := r.collection.UpdateOne(ctx, r.filters, update)
	if err != nil {
		return fmt.Errorf("failed to delete report: %w", err)
	}
	return nil
}
