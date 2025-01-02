package db

import (
	"context"
	"database/sql"

	"github.com/cifra-city/users-storage/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Users interface {
	Crete(ctx context.Context, id uuid.UUID, username string) (dbcore.User, error)

	Get(ctx context.Context, id uuid.UUID) (dbcore.User, error)
	GetByUsername(ctx context.Context, username string) (dbcore.User, error)

	UpdateUsername(ctx context.Context, id uuid.UUID, username string) (dbcore.User, error)
	UpdateTitle(ctx context.Context, id uuid.UUID, title *string) (dbcore.User, error)
	UpdateAvatar(ctx context.Context, id uuid.UUID, avatar *string) (dbcore.User, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status *string) (dbcore.User, error)
	UpdateBio(ctx context.Context, id uuid.UUID, bio *string) (dbcore.User, error)
	UpdateCity(ctx context.Context, id uuid.UUID, city uuid.UUID) (dbcore.User, error)

	UpdateFull(ctx context.Context, id uuid.UUID, username *string, title *string, avatar *string, status *string, bio *string, city *uuid.UUID) (dbcore.User, error)

	Search(ctx context.Context, text *string, limit int, offset int) ([]dbcore.User, error)
}

type users struct {
	queries *dbcore.Queries
}

func NewUsers(queries *dbcore.Queries) Users {
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

func (u *users) Crete(ctx context.Context, id uuid.UUID, username string) (dbcore.User, error) {
	return u.queries.CreateUser(ctx, dbcore.CreateUserParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) Get(ctx context.Context, id uuid.UUID) (dbcore.User, error) {
	return u.queries.GetUser(ctx, id)
}

func (u *users) GetByUsername(ctx context.Context, username string) (dbcore.User, error) {
	return u.queries.GetUserByUsername(ctx, username)
}

func (u *users) UpdateUsername(ctx context.Context, id uuid.UUID, username string) (dbcore.User, error) {
	return u.queries.UpdateUsername(ctx, dbcore.UpdateUsernameParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) UpdateTitle(ctx context.Context, id uuid.UUID, title *string) (dbcore.User, error) {
	return u.queries.UpdateTitle(ctx, dbcore.UpdateTitleParams{
		ID:    id,
		Title: StmtNullString(title),
	})
}

func (u *users) UpdateAvatar(ctx context.Context, id uuid.UUID, avatar *string) (dbcore.User, error) {
	return u.queries.UpdateAvatar(ctx, dbcore.UpdateAvatarParams{
		ID:     id,
		Avatar: StmtNullString(avatar),
	})
}

func (u *users) UpdateStatus(ctx context.Context, id uuid.UUID, status *string) (dbcore.User, error) {
	return u.queries.UpdateStatus(ctx, dbcore.UpdateStatusParams{
		ID:     id,
		Status: StmtNullString(status),
	})
}

func (u *users) UpdateBio(ctx context.Context, id uuid.UUID, bio *string) (dbcore.User, error) {
	return u.queries.UpdateBio(ctx, dbcore.UpdateBioParams{
		ID:  id,
		Bio: StmtNullString(bio),
	})
}
func (u *users) UpdateCity(ctx context.Context, id uuid.UUID, city uuid.UUID) (dbcore.User, error) {
	return u.queries.UpdateCity(ctx, dbcore.UpdateCityParams{
		ID:   id,
		City: StmtNullUUID(&city),
	})
}

func (u *users) UpdateFull(ctx context.Context, id uuid.UUID, username *string, title *string, avatar *string, status *string, bio *string, city *uuid.UUID) (dbcore.User, error) {

	return u.queries.UpdateFullUser(ctx, dbcore.UpdateFullUserParams{
		ID:       id,
		Username: *username,
		Title:    StmtNullString(title),
		Avatar:   StmtNullString(avatar),
		Status:   StmtNullString(status),
		Bio:      StmtNullString(bio),
		City:     StmtNullUUID(city),
	})
}

func (u *users) Search(ctx context.Context, text *string, limit int, offset int) ([]dbcore.User, error) {
	res, err := u.queries.SearchUsers(ctx, dbcore.SearchUsersParams{
		Column1: StmtNullString(text),
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}

	var usersCol []dbcore.User
	for _, el := range res {
		usersCol = append(usersCol, dbcore.User{
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
