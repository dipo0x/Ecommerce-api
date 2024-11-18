package types

import (
	"github.com/google/uuid"
)
type IUser struct {
	ID  uuid.UUID `bson:"_id,omitempty" json:"id"`
    Username   string `json:"username" validate:"required"`
    FName string `json:"fName" validate:"required"`
    LName  string `json:"lName" validate:"required"`
}