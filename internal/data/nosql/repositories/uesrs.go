package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	New() Users
	Insert(ctx context.Context, user models.User) (*models.User, error)
	DeleteOne(ctx context.Context) error
	DeleteMany(ctx context.Context) (int64, error)
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.User, error)
	Get(ctx context.Context) (*models.User, error)

	Filter(filters map[string]any) Users
	FilterCoincidence(filters map[string]any) Users

	UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error)
	UpdateMany(ctx context.Context, fields map[string]any) (int64, error)

	Ideas() Ideas
	Projects() Projects
	Organizations() Organizations
	ReportsSent() ReportsSent
	ReportsReceived() ReportsReceived

	SortBy(field string, ascending bool) Users
	Limit(limit int64) Users
	Skip(skip int64) Users
}

type users struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewUsers(uri, dbName, collectionName string) (Users, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	coll := database.Collection(collectionName)

	return &users{
		client:     client,
		database:   database,
		collection: coll,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (u *users) New() Users {
	return &users{
		client:     u.client,
		database:   u.database,
		collection: u.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (u *users) Insert(ctx context.Context, user models.User) (*models.User, error) {
	user.CreatedAt = time.Now()

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return &user, nil
}

func (u *users) DeleteMany(ctx context.Context) (int64, error) {
	result, err := u.collection.DeleteMany(ctx, u.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to delete users: %w", err)
	}
	return result.DeletedCount, nil
}

func (u *users) DeleteOne(ctx context.Context) error {
	_, err := u.collection.DeleteOne(ctx, u.filters)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (u *users) Count(ctx context.Context) (int64, error) {
	count, err := u.collection.CountDocuments(ctx, u.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}
	return count, nil
}

func (u *users) Select(ctx context.Context) ([]models.User, error) {
	options := options.Find().SetSort(u.sort).SetLimit(u.limit).SetSkip(u.skip)
	cursor, err := u.collection.Find(ctx, u.filters, options)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("failed to decode users: %w", err)
	}
	return users, nil
}

func (u *users) Get(ctx context.Context) (*models.User, error) {
	var user models.User
	err := u.collection.FindOne(ctx, u.filters).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}

func (u *users) Filter(filters map[string]any) Users {
	var validFilters = map[string]bool{
		"_id":      true,
		"username": true,
		"role":     true,
		"ban_id":   true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		u.filters[field] = value
	}
	return u
}

func (u *users) FilterCoincidence(filters map[string]any) Users {
	var validFilters = map[string]bool{
		"username": true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}

		strValue, ok := value.(string)
		if !ok || strValue == "" {
			continue
		}

		u.filters[field] = bson.M{
			"$regex":   fmt.Sprintf(".*%s.*", strValue),
			"$options": "i",
		}
	}

	return u
}

func (u *users) UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	validFields := map[string]bool{
		"username": true,
		"role":     true,
		"avatar":   true,
		"ban_id":   true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	updateFields["updated_at"] = time.Now()

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no valid fields to update")
	}

	if u.filters == nil || u.filters["_id"] == nil {
		return nil, errors.New("team filters are empty or team ID is not set")
	}

	filter := bson.M{"_id": u.filters["_id"]}
	update := bson.M{"$set": updateFields}

	result, err := u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no team found with the given criteria")
	}

	var user models.User
	err = u.collection.FindOne(ctx, u.filters).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}

func (u *users) UpdateMany(ctx context.Context, fields map[string]any) (int64, error) {
	if len(fields) == 0 {
		return 0, fmt.Errorf("no fields to update")
	}

	validFields := map[string]bool{
		"username": true,
		"role":     true,
		"avatar":   true,
		"ban_id":   true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	updateFields["updated_at"] = time.Now()

	if len(updateFields) == 0 {
		return 0, fmt.Errorf("no valid fields to update")
	}

	if u.filters == nil || u.filters["_id"] == nil {
		return 0, errors.New("team filters are empty or team ID is not set")
	}

	filter := bson.M{"_id": u.filters["_id"]}
	update := bson.M{"$set": updateFields}

	result, err := u.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update team: %w", err)
	}

	return result.ModifiedCount, nil
}

func (u *users) Ideas() Ideas {
	return &ideas{
		client:     u.client,
		database:   u.database,
		collection: u.collection,
		filters:    u.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (u *users) Projects() Projects {
	return &projects{
		client:     u.client,
		database:   u.database,
		collection: u.collection,
		filters:    u.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (u *users) Organizations() Organizations {
	return &organizations{
		client:     u.client,
		database:   u.database,
		collection: u.collection,
		filters:    u.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (u *users) ReportsSent() ReportsSent {
	return &reportsSent{
		client:     u.client,
		database:   u.database,
		collection: u.collection,
		filters:    u.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (u *users) ReportsReceived() ReportsReceived {
	return &reportsReceived{
		client:     u.client,
		database:   u.database,
		collection: u.collection,
		filters:    u.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (u *users) SortBy(field string, ascending bool) Users {
	order := 1
	if !ascending {
		order = -1
	}
	u.sort = append(u.sort, bson.E{Key: field, Value: order})
	return u
}

func (u *users) Limit(limit int64) Users {
	u.limit = limit
	return u
}

func (u *users) Skip(skip int64) Users {
	u.skip = skip
	return u
}
