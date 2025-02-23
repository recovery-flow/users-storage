package domain

import (
	"github.com/recovery-flow/users-storage/internal/service/infra"
	"github.com/sirupsen/logrus"
)

type Domain interface {
}

type domain struct {
	Infra *infra.Infra
	log   *logrus.Logger
}
