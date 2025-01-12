package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/internal/service/roles"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	New() Users
	Insert(ctx context.Context, user models.User) error
	Delete(ctx context.Context) (int64, error)
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.User, error)
	Get(ctx context.Context) (models.User, error)

	FilterById(id uuid.UUID) Users
	FilterByUsername(username string) Users

	Update(ctx context.Context, username string, role roles.UserRole, avatar string) (int64, error)
	UpdateUsername(ctx context.Context, username string) (int64, error)
	UpdateRole(ctx context.Context, role roles.UserRole) (int64, error)
	UpdateAvatar(ctx context.Context, avatar string) (int64, error)

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

func (u *users) Insert(ctx context.Context, user models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (u *users) Delete(ctx context.Context) (int64, error) {
	result, err := u.collection.DeleteMany(ctx, u.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to delete users: %w", err)
	}
	return result.DeletedCount, nil
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

func (u *users) Get(ctx context.Context) (models.User, error) {
	var user models.User
	err := u.collection.FindOne(ctx, u.filters).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, nil
		}
		return models.User{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}

func (u *users) FilterById(id uuid.UUID) Users {
	u.filters["_id"] = id
	return u
}

func (u *users) FilterByUsername(username string) Users {
	u.filters["username"] = username
	return u
}

func (u *users) Update(ctx context.Context, username string, role roles.UserRole, avatar string) (int64, error) {
	update := bson.M{
		"$set": bson.M{
			"username":   username,
			"role":       role,
			"avatar":     avatar,
			"updated_at": time.Now(),
		},
	}

	result, err := u.collection.UpdateMany(ctx, u.filters, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update users: %w", err)
	}
	return result.ModifiedCount, nil
}

func (u *users) UpdateUsername(ctx context.Context, username string) (int64, error) {
	update := bson.M{"$set": bson.M{"username": username, "updated_at": time.Now()}}
	result, err := u.collection.UpdateMany(ctx, u.filters, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update username: %w", err)
	}
	return result.ModifiedCount, nil
}

func (u *users) UpdateRole(ctx context.Context, role roles.UserRole) (int64, error) {
	update := bson.M{"$set": bson.M{"username": role, "role": time.Now()}}
	result, err := u.collection.UpdateMany(ctx, u.filters, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update username: %w", err)
	}
	return result.ModifiedCount, nil
}

func (u *users) UpdateAvatar(ctx context.Context, avatar string) (int64, error) {
	update := bson.M{"$set": bson.M{"avatar": avatar, "updated_at": time.Now()}}
	result, err := u.collection.UpdateMany(ctx, u.filters, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update username: %w", err)
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
