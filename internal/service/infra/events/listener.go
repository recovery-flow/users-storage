package events

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/infra/events/callbacks"
	"github.com/sirupsen/logrus"
)

const (
	accountQ         = "account"
	accountCreateKey = "account.create"
)

func Listener(ctx context.Context) {
	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVICE)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to get server from context")
	}

	type QueueConfig struct {
		QueueName  string
		RoutingKey string
		Callback   func(context.Context, []byte) error
	}

	queues := []QueueConfig{
		{accountQ, accountCreateKey, callbacks.CreateAccount},
	}

	for _, q := range queues {
		if err := server.Broker.Listen(ctx, server.Logger, q.QueueName, q.RoutingKey, q.Callback); err != nil {
			logrus.WithError(err).Fatalf("Listener encountered an error for queue %s with key %s", q.QueueName, q.RoutingKey)
		}
	}
}
