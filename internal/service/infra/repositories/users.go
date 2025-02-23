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
)

// Users описывает интерфейс репозитория для работы с пользователями.
type Users interface {
	New(ctx context.Context) Users

	Create(ctx context.Context, user models.User) (*models.User, error)

	Get(ctx context.Context) (*models.User, error)
	Select(ctx context.Context) ([]models.User, error)

	Filter(ctx context.Context, filter map[string]any) Users

	UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error)
	UpdateMany(ctx context.Context, fields map[string]any) (*models.User, error)

	Delete(ctx context.Context, userID string) error

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

// NewUsers создаёт новый репозиторий, инициализируя Redis и MongoDB.
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

// New возвращает новый репозиторий с очищенным набором фильтров.
func (u *users) New(ctx context.Context) Users {
	return &users{
		redis:         u.redis,
		mongo:         u.mongo,
		filters:       make(map[string]any),
		sort:          "",
		sortAscending: true,
		limit:         0,
		skip:          0,
		log:           u.log,
	}
}

// Create создаёт пользователя в MongoDB, а затем добавляет его в кэш.
func (u *users) Create(ctx context.Context, user models.User) (*models.User, error) {
	newUser, err := u.mongo.Insert(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user in mongo")
	}
	if err := u.redis.Add(ctx, *newUser); err != nil {
		u.log.Warn("failed to add user to cache: ", err)
	}
	return newUser, nil
}

// Get пытается сначала получить пользователя из кэша (по _id или username).
// Если данные не найдены, происходит обращение к MongoDB и запись результата в кэш.
func (u *users) Get(ctx context.Context) (*models.User, error) {
	// Если в фильтрах указан _id, пытаемся получить данные по нему из Redis.
	if id, ok := u.filters["_id"]; ok {
		var uid uuid.UUID
		switch v := id.(type) {
		case uuid.UUID:
			uid = v
		case string:
			var err error
			uid, err = uuid.Parse(v)
			if err != nil {
				u.log.Warn("failed to parse _id: ", err)
			}
		}
		if uid != uuid.Nil {
			user, err := u.redis.GetByID(ctx, uid)
			if err == nil {
				return user, nil
			}
			u.log.Info("user not found in cache by _id: ", err)
		}
	}
	// Если указан username, пытаемся получить пользователя по username.
	if username, ok := u.filters["username"]; ok {
		if name, ok := username.(string); ok && name != "" {
			user, err := u.redis.GetByUsername(ctx, name)
			if err == nil {
				return user, nil
			}
			u.log.Info("user not found in cache by username: ", err)
		}
	}
	// Если в кэше ничего не найдено, запрашиваем данные из MongoDB.
	m := u.mongo.New()
	// Применяем фильтры. Здесь делим фильтр на строгие, мягкие и по датам.
	strictFilters := make(map[string]any)
	softFilters := make(map[string]any)
	dateFilters := make(map[string]any)
	for k, v := range u.filters {
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
		m = m.FilterStrict(strictFilters)
	}
	if len(softFilters) > 0 {
		m = m.FilterSoft(softFilters)
	}
	if len(dateFilters) > 0 {
		// Здесь по умолчанию считаем фильтр "после" (after) true.
		m = m.FilterDate(dateFilters, true)
	}
	user, err := m.Get(ctx)
	if err != nil {
		return nil, err
	}
	// Если пользователь получен из Mongo, сохраняем его в кэш.
	if err := u.redis.Add(ctx, *user); err != nil {
		u.log.Warn("failed to add user to cache after mongo fetch: ", err)
	}
	return user, nil
}

// Select возвращает список пользователей по заданным фильтрам, сортировке и лимиту.
// Для выборки множества объектов кэширование не применяется.
func (u *users) Select(ctx context.Context) ([]models.User, error) {
	m := u.mongo.New()
	strictFilters := make(map[string]any)
	softFilters := make(map[string]any)
	dateFilters := make(map[string]any)
	for k, v := range u.filters {
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
		m = m.FilterStrict(strictFilters)
	}
	if len(softFilters) > 0 {
		m = m.FilterSoft(softFilters)
	}
	if len(dateFilters) > 0 {
		m = m.FilterDate(dateFilters, true)
	}
	if u.limit > 0 {
		m = m.Limit(u.limit)
	}
	if u.skip > 0 {
		m = m.Skip(u.skip)
	}
	if u.sort != "" {
		m = m.SortBy(u.sort, u.sortAscending)
	}
	return m.Select(ctx)
}

// Filter добавляет переданные фильтры в репозиторий и возвращает обновлённый интерфейс.
func (u *users) Filter(ctx context.Context, filter map[string]any) Users {
	for k, v := range filter {
		u.filters[k] = v
	}
	return u
}

// UpdateOne обновляет одного пользователя в MongoDB по заданным полям и обновляет кэш.
func (u *users) UpdateOne(ctx context.Context, fields map[string]any) (*models.User, error) {
	user, err := u.mongo.UpdateOne(ctx, fields)
	if err != nil {
		return nil, err
	}
	if err := u.redis.Add(ctx, *user); err != nil {
		u.log.Warn("failed to update cache for user: ", err)
	}
	return user, nil
}

// UpdateMany поддерживается только если фильтр содержит _id (т.е. обновляется один пользователь).
func (u *users) UpdateMany(ctx context.Context, fields map[string]any) (*models.User, error) {
	if _, exists := u.filters["_id"]; !exists {
		return nil, errors.New("UpdateMany is only supported for a single user with a specified _id")
	}
	user, err := u.mongo.UpdateOne(ctx, fields)
	if err != nil {
		return nil, err
	}
	if err := u.redis.Add(ctx, *user); err != nil {
		u.log.Warn("failed to update cache for user: ", err)
	}
	return user, nil
}

// Delete удаляет пользователя из MongoDB (с учётом фильтра) и удаляет его запись из кэша.
func (u *users) Delete(ctx context.Context, userID string) error {
	u.filters["_id"] = userID
	if err := u.mongo.DeleteOne(ctx); err != nil {
		return err
	}
	if err := u.redis.Delete(ctx, userID); err != nil {
		u.log.Warn("failed to delete user from cache: ", err)
	}
	return nil
}

// Sort задаёт поле сортировки и порядок (по возрастанию/убыванию) и возвращает обновлённый репозиторий.
func (u *users) Sort(field string, ascending bool) Users {
	u.sort = field
	u.sortAscending = ascending
	return u
}

// Limit задаёт ограничение на количество записей.
func (u *users) Limit(limit int64) Users {
	u.limit = limit
	return u
}

// Skip задаёт количество записей, которые необходимо пропустить.
func (u *users) Skip(skip int64) Users {
	u.skip = skip
	return u
}
