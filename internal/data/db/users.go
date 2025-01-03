package db

import (
	"context"
	"database/sql"

	"github.com/cifra-city/users-storage/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Users interface {
	Crete(ctx context.Context, id uuid.UUID, username string) (sqlcore.User, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore.User, error)
	GetByUsername(ctx context.Context, username string) (sqlcore.User, error)

	UpdateUsername(ctx context.Context, id uuid.UUID, username string) (sqlcore.User, error)
	UpdateTitle(ctx context.Context, id uuid.UUID, title *string) (sqlcore.User, error)
	UpdateAvatar(ctx context.Context, id uuid.UUID, avatar *string) (sqlcore.User, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status *string) (sqlcore.User, error)
	UpdateBio(ctx context.Context, id uuid.UUID, bio *string) (sqlcore.User, error)
	UpdateCity(ctx context.Context, id uuid.UUID, city uuid.UUID) (sqlcore.User, error)

	UpdateFull(ctx context.Context, id uuid.UUID, username *string, title *string, avatar *string, status *string, bio *string, city *uuid.UUID) (sqlcore.User, error)

	Search(ctx context.Context, text *string, limit int, offset int) ([]sqlcore.User, error)
}

type users struct {
	queries *sqlcore.Queries
}

func NewUsers(queries *sqlcore.Queries) Users {
	return &users{queries: queries}
}

func StmtNullString(s *string) sql.NullString {
	var stmt sql.NullString
	if s != nil {
		stmt.String = *s
		stmt.Valid = true
	}
	return stmt
}

func StmtNullUUID(s *uuid.UUID) uuid.NullUUID {
	var stmt uuid.NullUUID
	if s != nil {
		stmt.UUID = *s
		stmt.Valid = true
	}
	return stmt
}

func (u *users) Crete(ctx context.Context, id uuid.UUID, username string) (sqlcore.User, error) {
	return u.queries.CreateUser(ctx, sqlcore.CreateUserParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) Get(ctx context.Context, id uuid.UUID) (sqlcore.User, error) {
	return u.queries.GetUser(ctx, id)
}

func (u *users) GetByUsername(ctx context.Context, username string) (sqlcore.User, error) {
	return u.queries.GetUserByUsername(ctx, username)
}

func (u *users) UpdateUsername(ctx context.Context, id uuid.UUID, username string) (sqlcore.User, error) {
	return u.queries.UpdateUsername(ctx, sqlcore.UpdateUsernameParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) UpdateTitle(ctx context.Context, id uuid.UUID, title *string) (sqlcore.User, error) {
	return u.queries.UpdateTitle(ctx, sqlcore.UpdateTitleParams{
		ID:    id,
		Title: StmtNullString(title),
	})
}

func (u *users) UpdateAvatar(ctx context.Context, id uuid.UUID, avatar *string) (sqlcore.User, error) {
	return u.queries.UpdateAvatar(ctx, sqlcore.UpdateAvatarParams{
		ID:     id,
		Avatar: StmtNullString(avatar),
	})
}

func (u *users) UpdateStatus(ctx context.Context, id uuid.UUID, status *string) (sqlcore.User, error) {
	return u.queries.UpdateStatus(ctx, sqlcore.UpdateStatusParams{
		ID:     id,
		Status: StmtNullString(status),
	})
}

func (u *users) UpdateBio(ctx context.Context, id uuid.UUID, bio *string) (sqlcore.User, error) {
	return u.queries.UpdateBio(ctx, sqlcore.UpdateBioParams{
		ID:  id,
		Bio: StmtNullString(bio),
	})
}
func (u *users) UpdateCity(ctx context.Context, id uuid.UUID, city uuid.UUID) (sqlcore.User, error) {
	return u.queries.UpdateCity(ctx, sqlcore.UpdateCityParams{
		ID:   id,
		City: StmtNullUUID(&city),
	})
}

func (u *users) UpdateFull(ctx context.Context, id uuid.UUID, username *string, title *string, avatar *string, status *string, bio *string, city *uuid.UUID) (sqlcore.User, error) {

	return u.queries.UpdateFullUser(ctx, sqlcore.UpdateFullUserParams{
		ID:       id,
		Username: *username,
		Title:    StmtNullString(title),
		Avatar:   StmtNullString(avatar),
		Status:   StmtNullString(status),
		Bio:      StmtNullString(bio),
		City:     StmtNullUUID(city),
	})
}

func (u *users) Search(ctx context.Context, text *string, limit int, offset int) ([]sqlcore.User, error) {
	res, err := u.queries.SearchUsers(ctx, sqlcore.SearchUsersParams{
		Column1: StmtNullString(text),
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}

	var usersCol []sqlcore.User
	for _, el := range res {
		usersCol = append(usersCol, sqlcore.User{
			ID:       el.ID,
			Username: el.Username,
			Title:    el.Title,
			Status:   el.Status,
			Avatar:   el.Avatar,
			Bio:      el.Bio,
		})
	}

	return usersCol, nil
}
