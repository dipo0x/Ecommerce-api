package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	username   string `bson:"username" json:"username" validate:"required"`
    fName string `bson:"fName" json:"fName" validate:"required"`
    lName  string `bson:"lName" json:"lName" validate:"required"`
	authId   uuid.UUID `bson:"authId" json:"authId" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
