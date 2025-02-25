package domain

import (
	"context"

	"github.com/recovery-flow/users-storage/internal/service/domain/models"
	"github.com/recovery-flow/users-storage/internal/service/infra"
	"github.com/sirupsen/logrus"
)

type Domain interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)

	GetUser(ctx context.Context, query models.RequestQuery) (*models.User, error)
	SelectUsers(ctx context.Context, query models.RequestQuery) ([]models.User, error)
	UpdateUser(ctx context.Context, query models.RequestQuery, updateFields map[string]interface{}) (*models.User, error)
	UpdateUsers(ctx context.Context, query models.RequestQuery, updateFields map[string]interface{}) ([]models.User, error)
}

type domain struct {
	Infra *infra.Infra
	log   *logrus.Logger
}

func NewDomain(infra *infra.Infra, log *logrus.Logger) (Domain, error) {
	return &domain{
		Infra: infra,
		log:   log,
	}, nil
}

func (d *domain) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return d.Infra.Users.Create(ctx, user)
}

func (d *domain) GetUser(ctx context.Context, query models.RequestQuery) (*models.User, error) {
	return d.Infra.Users.New().Filter(query.Filters).Get(ctx)
}

func (d *domain) SelectUsers(ctx context.Context, query models.RequestQuery) ([]models.User, error) {
	return d.Infra.Users.New().Filter(query.Filters).
		Skip(query.Offset).Limit(query.Limit).Sort(query.SortField, query.SortAscending).Select(ctx)
}

func (d *domain) UpdateUser(ctx context.Context, query models.RequestQuery, updateFields map[string]interface{}) (*models.User, error) {
	return d.Infra.Users.New().Filter(query.Filters).UpdateOne(ctx, updateFields)
}

func (d *domain) UpdateUsers(ctx context.Context, query models.RequestQuery, updateFields map[string]interface{}) ([]models.User, error) {
	return d.Infra.Users.New().Filter(query.Filters).UpdateMany(ctx, updateFields)
}
