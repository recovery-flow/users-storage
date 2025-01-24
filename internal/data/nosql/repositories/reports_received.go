package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportsReceived interface {
	Insert(ctx context.Context, report primitive.ObjectID) error
	Delete(ctx context.Context, report primitive.ObjectID) error
}

type reportsReceived struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (r *reportsReceived) Insert(ctx context.Context, reportreportreport primitive.ObjectID) error {
	update := bson.M{
		"$addToSet": bson.M{
			"report_received": reportreportreport,
		},
	}
	_, err := r.collection.UpdateOne(ctx, r.filters, update)
	if err != nil {
		return fmt.Errorf("failed to insert report: %w", err)
	}
	return nil
}

func (r *reportsReceived) Delete(ctx context.Context, report primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"report_received": report,
		},
	}
	_, err := r.collection.UpdateOne(ctx, r.filters, update)
	if err != nil {
		return fmt.Errorf("failed to delete report: %w", err)
	}
	return nil
}
