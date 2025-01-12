package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `bson:"_id"              json:"id"`
	Username  string    `bson:"username"         json:"username"`
	Role      string    `bson:"role"             json:"role"`
	Avatar    string    `bson:"avatar,omitempty" json:"avatar,omitempty"`
	CreatedAt time.Time `bson:"created_at"       json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"       json:"updated_at"`
}
