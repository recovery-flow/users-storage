package models

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID          uuid.UUID `bson:"_id"    json:"id"`
	Name        string    `bson:"name"   json:"name"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Members     []Member  `bson:"members" json:"members"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
