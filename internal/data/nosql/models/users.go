package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/roles"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              uuid.UUID            `bson:"_id"                        json:"id"`
	Username        string               `bson:"username"                   json:"username"`
	Role            roles.UserRole       `bson:"role"                       json:"role"`
	Avatar          string               `bson:"avatar,omitempty"           json:"avatar,omitempty"`
	Organizations   []primitive.ObjectID `bson:"organizations,omitempty"    json:"organizations,omitempty"`
	Projects        []primitive.ObjectID `bson:"projects,omitempty"         json:"projects,omitempty"`
	Ideas           []primitive.ObjectID `bson:"ideas,omitempty"            json:"ideas,omitempty"`
	ReportsSent     []primitive.ObjectID `bson:"reports_sent,omitempty"     json:"reports_sent,omitempty"`
	ReportsReceived []primitive.ObjectID `bson:"reports_received,omitempty" json:"reports_received,omitempty"`
	Banned          *Accessibility       `bson:"banned,omitempty"           json:"banned,omitempty"`

	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt time.Time `bson:"created_at"           json:"created_at"`
}

type Accessibility struct {
	Banned      bool       `bson:"banned"                 json:"banned"`
	Start       *time.Time `bson:"start,omitempty"        json:"start,omitempty"`
	End         *time.Time `bson:"end,omitempty"          json:"end,omitempty"`
	Sort        *BanSort   `bson:"sort,omitempty"         json:"sort,omitempty"`
	Desc        string     `bson:"desc,omitempty"         json:"desc,omitempty"`
	InitiatorID *uuid.UUID `bson:"initiator_id,omitempty" json:"initiator_id,omitempty"`
}

type BanSort string

const (
	CommentsBan  BanSort = "comments"
	ActivityBan  BanSort = "activity"
	PermanentBan BanSort = "permanent"
)
