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

	OrganizationQ          = "organization"
	OrganizationAddUser    = "organization.add_user"
	OrganizationRemoveUser = "organization.remove_user"

	ProjectQ          = "project"
	ProjectAddUser    = "project.add_user"
	ProjectRemoveUser = "project.remove_user"

	IdeaQ          = "idea"
	IdeaAddUser    = "idea.add_user"
	IdeaRemoveUser = "idea.remove_user"

	ReportQ          = "report"
	ReportSentCreate = "report.user_to_user"

	BanQ           = "ban"
	BanCreateKey   = "ban.user"
	UnbanCreateKey = "unban.user"
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

	err = server.Broker.Listen(ctx, server.Logger, OrganizationQ, OrganizationAddUser, callbacks.OrganizationAddUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, OrganizationQ, OrganizationRemoveUser, callbacks.OrganizationRemoveUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, ProjectQ, ProjectAddUser, callbacks.ProjectAddUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, ProjectQ, ProjectRemoveUser, callbacks.ProjectRemoveUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, IdeaQ, IdeaAddUser, callbacks.IdeaAddUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, IdeaQ, IdeaRemoveUser, callbacks.IdeaRemoveUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, ReportQ, ReportSentCreate, callbacks.ReportUserToUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, BanQ, BanCreateKey, callbacks.BanUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, BanQ, UnbanCreateKey, callbacks.UnbanUser)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}
}
