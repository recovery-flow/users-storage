package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/recovery-flow/users-storage/internal/service/infra/repositories/cache"
	"github.com/recovery-flow/users-storage/internal/service/infra/repositories/mongodb"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	New() Users
	Create(ctx context.Context, user models.User) (*models.User, error)

	Get(ctx context.Context) (*models.User, error)
	Select(ctx context.Context) ([]models.User, error)

	Filter(filter map[string]any) *mongodb.Users

	UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error)
	UpdateMany(ctx context.Context, fields map[string]any) ([]models.User, error)

	Sort(field string, ascending bool) Users
	Limit(limit int64) Users
	Skip(skip int64) Users
}

type users struct {
	redis *cache.Users
	mongo *mongodb.Users

	filters       map[string]any
	sort          string
	sortAscending bool
	limit         int64
	skip          int64

	log *logrus.Logger
}

func NewUsers(cfg *config.Config, log *logrus.Logger) (Users, error) {
	redisRepo := cache.NewUsers(
		redis.NewClient(&redis.Options{
			Addr:     cfg.Database.Redis.Addr,
			Password: cfg.Database.Redis.Password,
			DB:       cfg.Database.Redis.DB,
		}),
		time.Duration(cfg.Database.Redis.Lifetime)*time.Minute,
	)

	mongoRepo, err := mongodb.NewUsers(cfg.Database.Mongo.URI, cfg.Database.Mongo.Name, mongodb.UsersCollection)
	if err != nil {
		return nil, err
	}
	return &users{
		redis:         redisRepo,
		mongo:         mongoRepo,
		filters:       make(map[string]any),
		sort:          "",
		sortAscending: true,
		limit:         0,
		skip:          0,
		log:           log,
	}, nil
}

func (u *users) New() Users {
	return &users{
		redis:         u.redis,
		mongo:         u.mongo.New(),
		filters:       make(map[string]any),
		sort:          "",
		sortAscending: false,
		limit:         0,
		skip:          0,
		log:           u.log,
	}
}

func (u *users) Create(ctx context.Context, user models.User) (*models.User, error) {
	newUser, err := u.mongo.Insert(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user in mongo")
	}
	if err := u.redis.Add(ctx, *newUser); err != nil {
		u.log.WithField("redis", err).Errorf("failed to add user to cache")
	}
	return newUser, nil
}

func (u *users) Get(ctx context.Context) (*models.User, error) {
	user, err := u.Filter(u.filters).Get(ctx)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			if id, ok := u.filters["_id"].(string); ok {
				uid, err := uuid.Parse(id)
				if err == nil {
					err = u.redis.DeleteByID(ctx, uid)
					if err != nil && !errors.Is(err, redis.Nil) {
						u.log.WithField("redis", err).Warn("failed to delete user from cache by ID")
					}
				}
			}

			if username, ok := u.filters["username"].(string); ok {
				err = u.redis.DeleteByUsername(ctx, username)
				if err != nil && !errors.Is(err, redis.Nil) {
					u.log.WithField("redis", err).Warn("failed to delete user from cache by username")
				}
			}
			return nil, err
		}
	}

	if err := u.redis.Add(ctx, *user); err != nil {
		u.log.WithField("redis", err).Errorf("failed to add user to cache")
	}

	return user, nil
}

func (u *users) Select(ctx context.Context) ([]models.User, error) {
	res, err := u.Filter(u.filters).Limit(u.limit).Skip(u.skip).SortBy(u.sort, u.sortAscending).Select(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *users) UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error) {
	user, err := u.Filter(u.filters).UpdateOne(ctx, fields)
	if err != nil {
		return nil, err
	}

	if err := u.redis.Add(ctx, *user); err != nil {
		u.log.WithField("redis", err).Errorf("failed to add user to cache")
	}

	return user, nil
}

func (u *users) UpdateMany(ctx context.Context, fields map[string]any) ([]models.User, error) {
	sum, err := u.Filter(u.filters).UpdateMany(ctx, fields)
	if err != nil {
		return nil, err
	}

	go func(ctx context.Context, users []models.User) {
		for _, user := range sum {
			if err := u.redis.DeleteByID(ctx, user.ID); err != nil {
				if !errors.Is(err, redis.Nil) {
					u.log.WithField("redis", err).Errorf("failed to delete user from cache by ID")
				}
			}
		}
	}(ctx, sum)

	return sum, nil
}

func (u *users) Filter(filters map[string]any) *mongodb.Users {
	strictFilters := make(map[string]any)
	softFilters := make(map[string]any)
	dateFilters := make(map[string]any)
	for k, v := range filters {
		switch k {
		case "_id", "role", "verified", "speciality", "position", "city", "country", "date_of_birth":
			strictFilters[k] = v
		case "username":
			softFilters[k] = v
		case "updated_at", "closed_at":
			dateFilters[k] = v
		default:
			strictFilters[k] = v
		}
	}
	if len(strictFilters) > 0 {
		u.mongo.FilterStrict(strictFilters)
	}
	if len(softFilters) > 0 {
		u.mongo.FilterSoft(softFilters)
	}
	if len(dateFilters) > 0 {
		u.mongo.FilterDate(dateFilters, true)
	}
	if u.limit > 0 {
		u.mongo.Limit(u.limit)
	}
	if u.skip > 0 {
		u.mongo.Skip(u.skip)
	}
	if u.sort != "" {
		u.mongo.SortBy(u.sort, u.sortAscending)
	}

	return u.mongo
}

func (u *users) Sort(field string, ascending bool) Users {
	u.sort = field
	u.sortAscending = ascending
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
