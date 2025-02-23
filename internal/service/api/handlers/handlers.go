package handlers

import (
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service"
	"github.com/recovery-flow/users-storage/internal/service/domain"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Config *config.Config
	Domain domain.Domain
	Log    *logrus.Logger
}

func NewHandlers(svc *service.Service) (*Handlers, error) {
	return &Handlers{
		Config: svc.Config,
		Domain: svc.Domain,
		Log:    svc.Log,
	}, nil
}
