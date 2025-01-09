package repositories

import (
	"context"
	"database/sql"

	sqlcore2 "github.com/cifra-city/users-storage/internal/data/sql/repositories/sqlcore"
	"github.com/google/uuid"
)

type Users interface {
	Crete(ctx context.Context, id uuid.UUID, username string) (sqlcore2.User, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore2.User, error)
	GetByUsername(ctx context.Context, username string) (sqlcore2.User, error)

	UpdateUsername(ctx context.Context, id uuid.UUID, username string) (sqlcore2.User, error)
	UpdateTitle(ctx context.Context, id uuid.UUID, title *string) (sqlcore2.User, error)
	UpdateAvatar(ctx context.Context, id uuid.UUID, avatar *string) (sqlcore2.User, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status *string) (sqlcore2.User, error)
	UpdateBio(ctx context.Context, id uuid.UUID, bio *string) (sqlcore2.User, error)
	UpdateCity(ctx context.Context, id uuid.UUID, city uuid.UUID) (sqlcore2.User, error)

	UpdateFull(ctx context.Context, id uuid.UUID, username *string, title *string, avatar *string, status *string, bio *string, city *uuid.UUID) (sqlcore2.User, error)

	Search(ctx context.Context, text *string, limit int, offset int) ([]sqlcore2.User, error)
}

type users struct {
	queries *sqlcore2.Queries
}

func NewUsers(queries *sqlcore2.Queries) Users {
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

func (u *users) Crete(ctx context.Context, id uuid.UUID, username string) (sqlcore2.User, error) {
	return u.queries.CreateUser(ctx, sqlcore2.CreateUserParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) Get(ctx context.Context, id uuid.UUID) (sqlcore2.User, error) {
	return u.queries.GetUser(ctx, id)
}

func (u *users) GetByUsername(ctx context.Context, username string) (sqlcore2.User, error) {
	return u.queries.GetUserByUsername(ctx, username)
}

func (u *users) UpdateUsername(ctx context.Context, id uuid.UUID, username string) (sqlcore2.User, error) {
	return u.queries.UpdateUsername(ctx, sqlcore2.UpdateUsernameParams{
		ID:       id,
		Username: username,
	})
}

func (u *users) UpdateTitle(ctx context.Context, id uuid.UUID, title *string) (sqlcore2.User, error) {
	return u.queries.UpdateTitle(ctx, sqlcore2.UpdateTitleParams{
		ID:    id,
		Title: StmtNullString(title),
	})
}

func (u *users) UpdateAvatar(ctx context.Context, id uuid.UUID, avatar *string) (sqlcore2.User, error) {
	return u.queries.UpdateAvatar(ctx, sqlcore2.UpdateAvatarParams{
		ID:     id,
		Avatar: StmtNullString(avatar),
	})
}

func (u *users) UpdateStatus(ctx context.Context, id uuid.UUID, status *string) (sqlcore2.User, error) {
	return u.queries.UpdateStatus(ctx, sqlcore2.UpdateStatusParams{
		ID:     id,
		Status: StmtNullString(status),
	})
}

func (u *users) UpdateBio(ctx context.Context, id uuid.UUID, bio *string) (sqlcore2.User, error) {
	return u.queries.UpdateBio(ctx, sqlcore2.UpdateBioParams{
		ID:  id,
		Bio: StmtNullString(bio),
	})
}
func (u *users) UpdateCity(ctx context.Context, id uuid.UUID, city uuid.UUID) (sqlcore2.User, error) {
	return u.queries.UpdateCity(ctx, sqlcore2.UpdateCityParams{
		ID:   id,
		City: StmtNullUUID(&city),
	})
}

func (u *users) UpdateFull(ctx context.Context, id uuid.UUID, username *string, title *string, avatar *string, status *string, bio *string, city *uuid.UUID) (sqlcore2.User, error) {

	return u.queries.UpdateFullUser(ctx, sqlcore2.UpdateFullUserParams{
		ID:       id,
		Username: *username,
		Title:    StmtNullString(title),
		Avatar:   StmtNullString(avatar),
		Status:   StmtNullString(status),
		Bio:      StmtNullString(bio),
		City:     StmtNullUUID(city),
	})
}

func (u *users) Search(ctx context.Context, text *string, limit int, offset int) ([]sqlcore2.User, error) {
	res, err := u.queries.SearchUsers(ctx, sqlcore2.SearchUsersParams{
		Column1: StmtNullString(text),
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}

	var usersCol []sqlcore2.User
	for _, el := range res {
		usersCol = append(usersCol, sqlcore2.User{
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
