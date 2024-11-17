package types

import (
	"github.com/google/uuid"
)
type IUser struct {
	ID  uuid.UUID `bson:"_id,omitempty" json:"id"`
    username   string `json:"username" validate:"required"`
    fName string `json:"fName" validate:"required"`
    lName  string `json:"lName" validate:"required"`
	authId   uuid.UUID `json:"authId" validate:"required"`
}