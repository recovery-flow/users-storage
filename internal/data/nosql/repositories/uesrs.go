package repositories

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	FilterStrict(filters map[string]any) Users
	FilterSoft(filters map[string]any) Users
	FilterNumber(filters map[string]any, greater bool) Users
	FilterDate(filters map[string]any, after bool) Users

	UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error)
	UpdateMany(ctx context.Context, fields map[string]any) (int64, error)

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
	user.CreatedAt = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

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
	opts := options.Find().SetSort(u.sort).SetLimit(u.limit).SetSkip(u.skip)
	cursor, err := u.collection.Find(ctx, u.filters, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	var usrs []models.User
	if err = cursor.All(ctx, &usrs); err != nil {
		return nil, fmt.Errorf("failed to decode users: %w", err)
	}
	return usrs, nil
}

func (u *users) Get(ctx context.Context) (*models.User, error) {
	var user models.User
	err := u.collection.FindOne(ctx, u.filters).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}

func (u *users) FilterStrict(filters map[string]any) Users {
	var validFilters = map[string]bool{
		"_id":        true,
		"username":   true,
		"role":       true,
		"type":       true,
		"verified":   true,
		"title_name": true,
		"speciality": true,
		"city":       true,
		"country":    true,
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

func (u *users) FilterSoft(filters map[string]any) Users {
	var validFilters = map[string]bool{
		"username":   true,
		"title_name": true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
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

func (u *users) FilterDate(filters map[string]any, after bool) Users {
	validDateFields := map[string]bool{
		"updated_at": true,
		"closed_at":  true,
	}

	var op string
	if after {
		op = "$gte"
	} else {
		op = "$lte"
	}

	for field, value := range filters {
		if !validDateFields[field] {
			continue
		}
		if value == nil {
			continue
		}

		var t time.Time
		switch val := value.(type) {
		case time.Time:
			t = val
		case *time.Time:
			t = *val
		case string:
			parsed, err := time.Parse(time.RFC3339, val)
			if err != nil {
				continue
			}
			t = parsed
		default:
			continue
		}

		u.filters[field] = bson.M{op: t}
	}

	return u
}

func (u *users) FilterNumber(filters map[string]any, greater bool) Users {
	validCountFields := map[string]bool{
		"level":  true,
		"points": true,
	}

	op := "$lte"
	if greater {
		op = "$gte"
	}

	for field, value := range filters {
		if !validCountFields[field] {
			continue
		}
		if value == nil {
			continue
		}

		var n int64
		switch val := value.(type) {
		case int:
			n = int64(val)
		case int64:
			n = val
		case float64:
			n = int64(val)
		case string:
			parsed, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				continue
			}
			n = parsed
		default:
			continue
		}
		u.filters[field] = bson.M{op: n}
	}

	return u
}

func (u *users) UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	validFields := map[string]bool{
		"username":   true,
		"role":       true,
		"type":       true,
		"verified":   true,
		"title_name": true,
		"speciality": true,
		"city":       true,
		"country":    true,
		"level":      true,
		"points":     true,
		"created_at": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no valid fields to update")
	}

	updateFields["updated_at"] = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	_, err := u.collection.UpdateOne(ctx, u.filters, bson.M{"$set": updateFields})
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %w", err)
	}

	return u.Get(ctx)
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

	updateFields["updated_at"] = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	if len(updateFields) == 0 {
		return 0, fmt.Errorf("no valid fields to update")
	}

	if u.filters == nil || u.filters["_id"] == nil {
		return 0, errors.New("team filters are empty or team ID is not set")
	}

	result, err := u.collection.UpdateMany(ctx, bson.M{"_id": u.filters["_id"]}, bson.M{"$set": updateFields})
	if err != nil {
		return 0, fmt.Errorf("failed to update team: %w", err)
	}

	return result.ModifiedCount, nil
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
