package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/users-storage/internal/data/nosql/models"
	"github.com/recovery-flow/users-storage/internal/service/roles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Teams интерфейс для работы с коллекцией команд
type Teams interface {
	New() Teams
	Insert(ctx context.Context, team models.Team) error
	Delete(ctx context.Context) (int64, error)
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Team, error)
	Get(ctx context.Context) (models.Team, error)

	FilterById(id uuid.UUID) Teams
	FilterByMemberId(memberId uuid.UUID) Teams
	FilterByMemberUserId(memberUserId uuid.UUID) Teams

	UpdateName(ctx context.Context, name string) (int64, error)
	UpdateDescription(ctx context.Context, description string) (int64, error)

	AddMember(ctx context.Context, teamId, userId uuid.UUID, role roles.TeamRole, description string) (models.Team, error) // Добавить участника
	DeleteMember(ctx context.Context, teamId, userId uuid.UUID) (models.Team, error)                                       // Удалить участника
	UpdateMember(ctx context.Context, teamId, userId uuid.UUID, role roles.TeamRole, description string) (int64, error)    // Обновить участника
	SelectMembers(ctx context.Context, teamId uuid.UUID) ([]models.Member, error)                                          // Получить всех участников команды
	GetMember(ctx context.Context, teamId, userId uuid.UUID) (models.Member, error)

	SortBy(field string, ascending bool) Teams
	Limit(limit int64) Teams
	Skip(skip int64) Teams
}

type teams struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

// NewTeams конструктор для Teams
func NewTeams(uri, dbName, collectionName string) (Teams, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &teams{
		client:     client,
		database:   database,
		collection: collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (t *teams) New() Teams {
	return &teams{
		client:     t.client,
		database:   t.database,
		collection: t.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (t *teams) Insert(ctx context.Context, team models.Team) error {
	team.ID = uuid.New()
	team.CreatedAt = time.Now()
	team.UpdatedAt = time.Now()

	_, err := t.collection.InsertOne(ctx, team)
	if err != nil {
		return fmt.Errorf("failed to insert team: %w", err)
	}
	return nil
}

func (t *teams) Delete(ctx context.Context) (int64, error) {
	result, err := t.collection.DeleteMany(ctx, t.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to delete teams: %w", err)
	}
	return result.DeletedCount, nil
}

func (t *teams) Count(ctx context.Context) (int64, error) {
	count, err := t.collection.CountDocuments(ctx, t.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count teams: %w", err)
	}
	return count, nil
}

func (t *teams) Select(ctx context.Context) ([]models.Team, error) {
	options := options.Find().SetSort(t.sort).SetLimit(t.limit).SetSkip(t.skip)
	cursor, err := t.collection.Find(ctx, t.filters, options)
	if err != nil {
		return nil, fmt.Errorf("failed to find teams: %w", err)
	}
	defer cursor.Close(ctx)

	var teams []models.Team
	if err := cursor.All(ctx, &teams); err != nil {
		return nil, fmt.Errorf("failed to decode teams: %w", err)
	}
	return teams, nil
}

func (t *teams) Get(ctx context.Context) (models.Team, error) {
	var team models.Team
	err := t.collection.FindOne(ctx, t.filters).Decode(&team)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Team{}, nil
		}
		return models.Team{}, fmt.Errorf("failed to find team: %w", err)
	}
	return team, nil
}

func (t *teams) FilterById(id uuid.UUID) Teams {
	t.filters["_id"] = id
	return t
}

func (t *teams) FilterByMemberId(memberId uuid.UUID) Teams {
	t.filters["members._id"] = memberId
	return t
}

func (t *teams) FilterByMemberUserId(memberUserId uuid.UUID) Teams {
	t.filters["members.user_id"] = memberUserId
	return t
}

func (t *teams) UpdateName(ctx context.Context, name string) (int64, error) {
	filter := t.filters
	update := bson.M{
		"$set": bson.M{"name": name, "updated_at": time.Now()},
	}

	result, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update name: %w", err)
	}
	return result.ModifiedCount, nil
}

func (t *teams) UpdateDescription(ctx context.Context, description string) (int64, error) {
	filter := t.filters
	update := bson.M{
		"$set": bson.M{"description": description, "updated_at": time.Now()},
	}

	result, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update description: %w", err)
	}
	return result.ModifiedCount, nil
}

func (t *teams) SortBy(field string, ascending bool) Teams {
	order := 1
	if !ascending {
		order = -1
	}
	t.sort = append(t.sort, bson.E{Key: field, Value: order})
	return t
}

func (t *teams) Skip(skip int64) Teams {
	t.skip = skip
	return t
}

func (t *teams) Limit(limit int64) Teams {
	t.limit = limit
	return t
}
