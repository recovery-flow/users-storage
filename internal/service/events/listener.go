package events

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/events/callbacks"
	"github.com/sirupsen/logrus"
)

const (
	accountQ         = "account"
	accountCreateKey = "account.create"
)

func Listener(ctx context.Context) {
	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, accountQ, accountCreateKey, callbacks.CreateAccount)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}
}
