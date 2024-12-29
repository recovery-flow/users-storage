package db

import (
	"database/sql"

	"github.com/cifra-city/users-storage/internal/data/db/dbcore"
)

type Databaser struct {
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

func NewDatabaser(url string) (*Databaser, error) {
	db, err := NewDBConnection(url)
	if err != nil {
		return nil, err
	}
	_ = dbcore.New(db)
	return &Databaser{}, nil
}
