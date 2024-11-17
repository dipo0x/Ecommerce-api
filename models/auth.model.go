package models

import (
	"time"
	"github.com/google/uuid"
)

type Auth struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	UserId   uuid.UUID `bson:"userId" json:"userId" validate:"required"`
	Password   string `bson:"password" json:"password" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
