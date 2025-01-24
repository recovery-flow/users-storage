package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              uuid.UUID            `bson:"_id"                        json:"id"`
	Username        string               `bson:"username"                   json:"username"`
	Role            string               `bson:"role"                       json:"role"`
	Avatar          string               `bson:"avatar,omitempty"           json:"avatar,omitempty"`
	Organizations   []primitive.ObjectID `bson:"organizations,omitempty"    json:"organizations,omitempty"`
	Projects        []primitive.ObjectID `bson:"projects,omitempty"         json:"projects,omitempty"`
	Ideas           []primitive.ObjectID `bson:"ideas,omitempty"            json:"ideas,omitempty"`
	ReportsSent     []primitive.ObjectID `bson:"reports_sent,omitempty"     json:"reports_sent,omitempty"`
	ReportsReceived []primitive.ObjectID `bson:"reports_received,omitempty" json:"reports_received,omitempty"`
	BanStatus       bool                 `bson:"ban_status"                 json:"ban_status"`

	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt time.Time `bson:"created_at"           json:"created_at"`
}
