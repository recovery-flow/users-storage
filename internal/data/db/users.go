package db

import (
	"database/sql"
	"net/http"

	"github.com/cifra-city/users-storage/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Users interface {
	Crete(r *http.Request, id uuid.UUID, username string) (dbcore.User, error)

	Get(r *http.Request, id uuid.UUID) (dbcore.User, error)
	GetByUsername(r *http.Request, username string) (dbcore.User, error)

	UpdateUsername(r *http.Request, id uuid.UUID, username string) (dbcore.User, error)
	UpdateTitle(r *http.Request, id uuid.UUID, title *string) (dbcore.User, error)
	UpdateAvatar(r *http.Request, id uuid.UUID, avatar *string) (dbcore.User, error)
	UpdateStatus(r *http.Request, id uuid.UUID, status *string) (dbcore.User, error)
	UpdateBio(r *http.Request, id uuid.UUID, bio *string) (dbcore.User, error)
	UpdateFull(r *http.Request, id uuid.UUID, username string, title *string, avatar *string, status *string, bio *string) (dbcore.User, error)
}

type users struct {
	queries *dbcore.Queries
}

func StmtNullString(s *string) sql.NullString {
	var stmt sql.NullString
	if s != nil {
		stmt.String = *s
		stmt.Valid = true
	}
	return stmt
}

func (u *users) Crete(r *http.Request, id uuid.UUID, username string) (dbcore.User, error) {
	return u.queries.CreateUser(r.Context(), dbcore.CreateUserParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) Get(r *http.Request, id uuid.UUID) (dbcore.User, error) {
	return u.queries.GetUser(r.Context(), id)
}

func (u *users) GetByUsername(r *http.Request, username string) (dbcore.User, error) {
	return u.queries.GetUserByUsername(r.Context(), username)
}

func (u *users) UpdateUsername(r *http.Request, id uuid.UUID, username string) (dbcore.User, error) {
	return u.queries.UpdateUsername(r.Context(), dbcore.UpdateUsernameParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) UpdateTitle(r *http.Request, id uuid.UUID, title *string) (dbcore.User, error) {
	return u.queries.UpdateTitle(r.Context(), dbcore.UpdateTitleParams{
		ID:    id,
		Title: StmtNullString(title),
	})
}

func (u *users) UpdateAvatar(r *http.Request, id uuid.UUID, avatar *string) (dbcore.User, error) {
	return u.queries.UpdateAvatar(r.Context(), dbcore.UpdateAvatarParams{
		ID:     id,
		Avatar: StmtNullString(avatar),
	})
}

func (u *users) UpdateStatus(r *http.Request, id uuid.UUID, status *string) (dbcore.User, error) {
	return u.queries.UpdateStatus(r.Context(), dbcore.UpdateStatusParams{
		ID:     id,
		Status: StmtNullString(status),
	})
}

func (u *users) UpdateBio(r *http.Request, id uuid.UUID, bio *string) (dbcore.User, error) {
	return u.queries.UpdateBio(r.Context(), dbcore.UpdateBioParams{
		ID:  id,
		Bio: StmtNullString(bio),
	})
}

func (u *users) UpdateFull(r *http.Request, id uuid.UUID, username string, title *string, avatar *string, status *string, bio *string) (dbcore.User, error) {
	return u.queries.UpdateFullUser(r.Context(), dbcore.UpdateFullUserParams{
		ID:       id,
		Username: username,
		Title:    StmtNullString(title),
		Avatar:   StmtNullString(avatar),
		Status:   StmtNullString(status),
		Bio:      StmtNullString(bio),
	})
}
