package types

import (
	"github.com/google/uuid"
)
type IAuth struct {
	ID  uuid.UUID `bson:"_id,omitempty" json:"id"`
	UserId   uuid.UUID `json:"userId" validate:"required"`
    Password   string `json:"password" validate:"required"`
}