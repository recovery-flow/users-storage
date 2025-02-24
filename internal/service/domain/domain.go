package domain

import (
	"context"
	"fmt"

	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/recovery-flow/users-storage/internal/service/infra"
	"github.com/sirupsen/logrus"
)

type RequestQuery struct {
	Filters       map[string]interface{} // Ключ-значение для фильтрации (например, {"username": "john"})
	SortField     string                 // Поле сортировки (например, "created_at")
	SortAscending bool                   // Направление сортировки: true - по возрастанию, false - по убыванию
	Limit         int64                  // Количество записей для выборки
	Offset        int64                  // Смещение для пагинации
}

type Domain interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)

	GetUser(ctx context.Context, query RequestQuery) (*models.User, error)
	QueryUsers(ctx context.Context, query RequestQuery) ([]models.User, error)
	DeleteUser(ctx context.Context, query RequestQuery) error
	UpdateUser(ctx context.Context, query RequestQuery, updateFields map[string]interface{}) ([]models.User, error)
}

type domain struct {
	Infra *infra.Infra
	log   *logrus.Logger
}

func (d *domain) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return d.Infra.Users.Create(ctx, user)
}

func (d *domain) GetUser(ctx context.Context, query RequestQuery) (*models.User, error) {
	return d.Infra.Users.New().Filter(query.Filters).Get(ctx)
}

func (d *domain) QueryUsers(ctx context.Context, query RequestQuery) ([]models.User, error) {
	repo := d.Infra.Users.New()
	if query.Filters != nil {
		repo = repo.Filter(query.Filters)
	}
	if query.SortField != "" {
		repo = repo.Sort(query.SortField, query.SortAscending)
	}
	if query.Limit != 0 {
		repo = repo.Limit(query.Limit)
	}
	if query.Offset != 0 {
		repo = repo.Skip(query.Offset)
	}

	return repo.Select(ctx)
}

func (d *domain) UpdateUser(ctx context.Context, query RequestQuery, updateFields map[string]interface{}) (*models.User, error) {
	repo := d.Infra.Users.New(ctx)
	if query.Filters != nil {
		repo = repo.Filter(ctx, query.Filters)
	}

	return repo.UpdateOne(ctx, updateFields)
}

func (d *domain) UpdateUsers(ctx context.Context, query RequestQuery, updateFields map[string]interface{}) ([]models.User, error) {
	repo := d.Infra.Users.New(ctx)
	if query.Filters != nil {
		repo = repo.Filter(ctx, query.Filters)
	}

	return repo.UpdateMany(ctx, updateFields)
}

func (d *domain) DeleteUser(ctx context.Context, query RequestQuery) error {
	if id, ok := query.Filters["_id"]; ok {
		if userID, ok := id.(string); ok {
			return d.Infra.Users.Delete(ctx, userID)
		}
		return fmt.Errorf("DeleteUser: _id filter имеет неверный тип")
	}
	return fmt.Errorf("DeleteUser: требуется фильтр по _id")
}
