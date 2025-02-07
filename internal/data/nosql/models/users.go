package models

import (
	"github.com/google/uuid"
	"github.com/recovery-flow/roles"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       uuid.UUID      `bson:"_id"                        json:"id"`
	Username string         `bson:"username"                   json:"username"`
	Role     roles.UserRole `bson:"role"                       json:"role"`
	Type     UserTypes      `bson:"type"                       json:"type"`
	Verified bool           `bson:"verified"                   json:"verified"`

	TitleName  *string `bson:"title_name,omitempty"                 json:"title_name,omitempty"`
	Speciality *string `bson:"speciality,omitempty"           json:"speciality,omitempty"`
	City       *string `bson:"city,omitempty"                 json:"city,omitempty"`
	Country    *string `bson:"country,omitempty"              json:"country,omitempty"`

	Level  int32 `bson:"level"  json:"level"`
	Points int32 `bson:"points" json:"points"`

	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
}

type UserTypes string

const (
	UserTypeUser UserTypes = "user"
	UserTypeOrg  UserTypes = "organization"
	UserTypeGov  UserTypes = "government"
)
