package infra

import (
	"github.com/recovery-flow/rerabbit"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/infra/repositories"
	"github.com/sirupsen/logrus"
)

type Infra struct {
	Users repositories.Users

	Rabbit rerabbit.RabbitBroker
}

func NewInfra(cfg *config.Config, log *logrus.Logger) (*Infra, error) {
	usersRepo, err := repositories.NewUsers(cfg, log)
	if err != nil {
		log.WithError(err).Fatal("failed to create users repository")
		return nil, err
	}
	eve, err := rerabbit.NewBroker(cfg.Rabbit.URL)
	if err != nil {
		log.WithError(err).Fatal("failed to create rabbit broker")
		return nil, err
	}

	return &Infra{
		Users:  usersRepo,
		Rabbit: eve,
	}, nil
}
