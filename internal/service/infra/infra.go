package infra

import (
	"github.com/recovery-flow/rerabbit"
	"github.com/recovery-flow/users-storage/internal/service/infra/repositories"
)

type Infra struct {
	Users repositories.Users

	Rabbit rerabbit.RabbitBroker
}
