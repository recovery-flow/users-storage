package handlers

import (
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service"
	"github.com/recovery-flow/users-storage/internal/service/domain"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Config *config.Config
	Domain domain.Domain
	Log    *logrus.Logger
}

func NewHandler(svc *service.Service) (*Handler, error) {
	return &Handler{
		Config: svc.Config,
		Domain: svc.Domain,
		Log:    svc.Log,
	}, nil
}
