package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

type Users struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewUsers(uri, dbName, collectionName string) (*Users, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	coll := database.Collection(collectionName)

	return &Users{
		client:     client,
		database:   database,
		collection: coll,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (u *Users) New() *Users {
	return &Users{
		client:     u.client,
		database:   u.database,
		collection: u.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (u *Users) Insert(ctx context.Context, user models.User) (*models.User, error) {
	user.CreatedAt = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return &user, nil
}

func (u *Users) DeleteMany(ctx context.Context) (int64, error) {
	result, err := u.collection.DeleteMany(ctx, u.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to delete Users: %w", err)
	}
	return result.DeletedCount, nil
}

func (u *Users) DeleteOne(ctx context.Context) error {
	_, err := u.collection.DeleteOne(ctx, u.filters)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (u *Users) Count(ctx context.Context) (int64, error) {
	return u.collection.CountDocuments(ctx, u.filters)
}

func (u *Users) Select(ctx context.Context) ([]models.User, error) {
	opts := options.Find().SetSort(u.sort).SetLimit(u.limit).SetSkip(u.skip)
	cursor, err := u.collection.Find(ctx, u.filters, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find Users: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var usrs []models.User
	if err = cursor.All(ctx, &usrs); err != nil {
		return nil, fmt.Errorf("failed to decode Users: %w", err)
	}
	return usrs, nil
}

func (u *Users) Get(ctx context.Context) (*models.User, error) {
	var user models.User
	err := u.collection.FindOne(ctx, u.filters).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}

func (u *Users) Filter(filters map[string]models.QueryFilter) *Users {
	strictFilters := make(map[string]any)
	softFilters := make(map[string]any)
	dateFilters := make(map[string]any)
	numFilters := make(map[string]any)

	for key, qf := range filters {
		switch qf.Type {
		case "strict":
			if qf.Method == "" || qf.Method == "eq" {
				strictFilters[key] = qf.Value
			} else {
				strictFilters[key] = bson.M{"$" + qf.Method: qf.Value}
			}
		case "soft":
			softFilters[key] = bson.M{"$regex": qf.Value, "$options": "i"}
		case "date":
			dateFilters[key] = bson.M{"$" + qf.Method: qf.Value}
		case "num":
			numFilters[key] = bson.M{"$" + qf.Method: qf.Value}
		default:
			strictFilters[key] = qf.Value
		}
	}

	if len(strictFilters) > 0 {
		u.FilterStrict(strictFilters)
	}
	if len(softFilters) > 0 {
		u.FilterSoft(softFilters)
	}
	if len(dateFilters) > 0 {
		u.FilterDate(dateFilters, true)
	}
	if len(numFilters) > 0 {
		u.FilterStrict(numFilters)
	}
	return u
}

func (u *Users) FilterStrict(filters map[string]any) *Users {
	var validFilters = map[string]bool{
		"_id":           true,
		"username":      true,
		"role":          true,
		"verified":      true,
		"speciality":    true,
		"position":      true,
		"city":          true,
		"country":       true,
		"date_of_birth": true,
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

func (u *Users) FilterSoft(filters map[string]any) *Users {
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

		if strVal, ok := value.(string); ok && strVal != "" {
			u.filters[field] = bson.M{
				"$regex":   fmt.Sprintf(".*%s.*", strVal),
				"$options": "i",
			}
		}
	}

	return u
}

func (u *Users) FilterDate(filters map[string]any, after bool) *Users {
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

func (u *Users) UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error) {
	validFields := map[string]bool{
		"username":      true,
		"role":          true,
		"type":          true,
		"verified":      true,
		"avatar":        true,
		"title_name":    true,
		"speciality":    true,
		"position":      true,
		"city":          true,
		"country":       true,
		"date_of_birth": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	updateFields["updated_at"] = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated models.User
	err := u.collection.FindOneAndUpdate(ctx, u.filters, bson.M{"$set": updateFields}, opts).Decode(&updated)
	if err != nil {
		return nil, fmt.Errorf("failed to update document: %w", err)
	}

	for key, value := range updateFields {
		if _, exists := u.filters[key]; exists {
			u.filters[key] = value
		}
	}

	return &updated, nil
}

func (u *Users) UpdateMany(ctx context.Context, fields map[string]any) ([]models.User, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	validFields := map[string]bool{
		"role":          true,
		"type":          true,
		"verified":      true,
		"avatar":        true,
		"title_name":    true,
		"speciality":    true,
		"position":      true,
		"city":          true,
		"country":       true,
		"date_of_birth": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	updateFields["updated_at"] = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	if u.filters == nil || len(u.filters) == 0 {
		return nil, errors.New("filters are empty, cannot determine which documents to update")
	}

	_, err := u.collection.UpdateMany(ctx, u.filters, bson.M{"$set": updateFields})
	if err != nil {
		return nil, fmt.Errorf("failed to update documents: %w", err)
	}

	// Обновляем карту фильтров: если обновляемые поля присутствуют в фильтрах, заменяем их значениями из updateFields
	for key, value := range updateFields {
		if _, exists := u.filters[key]; exists {
			u.filters[key] = value
		}
	}

	updatedDocs, err := u.Limit(u.limit).Select(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated documents: %w", err)
	}

	return updatedDocs, nil
}

func (u *Users) Limit(limit int64) *Users {
	u.limit = limit
	return u
}

func (u *Users) Skip(skip int64) *Users {
	u.skip = skip
	return u
}

func (u *Users) SortBy(field string, ascending bool) *Users {
	order := 1
	if !ascending {
		order = -1
	}
	u.sort = append(u.sort, bson.E{Key: field, Value: order})
	return u
}
