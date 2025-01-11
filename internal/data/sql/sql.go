package sql

import (
	"database/sql"

	"github.com/recovery-flow/users-storage/internal/data/sql/repositories"
	"github.com/recovery-flow/users-storage/internal/data/sql/repositories/sqlcore"
)

type Repo struct {
	Users repositories.Users
}

func NewDBConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepoSQL(url string) (*Repo, error) {
	db, err := NewDBConnection(url)
	if err != nil {
		return nil, err
	}
	queries := sqlcore.New(db)
	return &Repo{
		Users: repositories.NewUsers(queries),
	}, nil
}
