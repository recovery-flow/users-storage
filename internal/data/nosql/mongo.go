package nosql

import (
	"fmt"

	"github.com/recovery-flow/users-storage/internal/data/nosql/repositories"
)

type Repo struct {
	Users repositories.Users
	Teams repositories.Teams
}

func NewRepositoryNoSql(uri, dbName string) (*Repo, error) {
	usersRepo, err := repositories.NewUsers(uri, dbName, "users")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize users repository: %w", err)
	}

	teamsRepo, err := repositories.NewTeams(uri, dbName, "teams")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize teams repository: %w", err)
	}

	return &Repo{
		Teams: teamsRepo,
		Users: usersRepo,
	}, nil
}
