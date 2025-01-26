package events

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/events/callbacks"
	"github.com/sirupsen/logrus"
)

const (
	accountQ               = "account"
	accountCreateKey       = "account.create"
	OrganizationQ          = "organization"
	OrganizationAddUser    = "organization.add_user"
	OrganizationRemoveUser = "organization.remove_user"
	ProjectQ               = "project"
	ProjectAddUser         = "project.add_user"
	ProjectRemoveUser      = "project.remove_user"
	IdeaQ                  = "idea"
	IdeaAddUser            = "idea.add_user"
	IdeaRemoveUser         = "idea.remove_user"
	ReportQ                = "report"
	ReportSentCreate       = "report.user_to_user"
	BanQ                   = "ban"
	BanCreateKey           = "ban.user"
	UnbanCreateKey         = "unban.user"
)

func Listener(ctx context.Context) {
	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
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
		{OrganizationQ, OrganizationAddUser, callbacks.OrganizationAddUser},
		{OrganizationQ, OrganizationRemoveUser, callbacks.OrganizationRemoveUser},
		{ProjectQ, ProjectAddUser, callbacks.ProjectAddUser},
		{ProjectQ, ProjectRemoveUser, callbacks.ProjectRemoveUser},
		{IdeaQ, IdeaAddUser, callbacks.IdeaAddUser},
		{IdeaQ, IdeaRemoveUser, callbacks.IdeaRemoveUser},
		{ReportQ, ReportSentCreate, callbacks.ReportUserToUser},
		{BanQ, BanCreateKey, callbacks.BanUser},
		{BanQ, UnbanCreateKey, callbacks.UnbanUser},
	}

	for _, q := range queues {
		if err := server.Broker.Listen(ctx, server.Logger, q.QueueName, q.RoutingKey, q.Callback); err != nil {
			logrus.WithError(err).Fatalf("Listener encountered an error for queue %s with key %s", q.QueueName, q.RoutingKey)
		}
	}
}
