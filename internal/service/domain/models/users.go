package models

import (
	"github.com/google/uuid"
	"github.com/recovery-flow/tokens/identity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       uuid.UUID        `bson:"_id"      json:"id"`
	Username string           `bson:"username" json:"username"`
	Role     identity.IdnType `bson:"role"     json:"role"`
	Verified bool             `bson:"verified" json:"verified"`

	Avatar      *string             `bson:"avatar,omitempty"        json:"avatar,omitempty"`
	TitleName   *string             `bson:"title_name,omitempty"    json:"title_name,omitempty"`
	Speciality  *string             `bson:"speciality,omitempty"    json:"speciality,omitempty"`
	Position    *string             `bson:"position,omitempty"      json:"position,omitempty"`
	City        *string             `bson:"city,omitempty"          json:"city,omitempty"`
	Country     *string             `bson:"country,omitempty"       json:"country,omitempty"`
	DateOfBirth *primitive.DateTime `bson:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`

	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
}
