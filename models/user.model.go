package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	Username   string `bson:"username" json:"username" validate:"required"`
    FName string `bson:"fName" json:"fName" validate:"required"`
    LName  string `bson:"lName" json:"lName" validate:"required"`
	AuthId   uuid.UUID `bson:"authId" json:"authId" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
