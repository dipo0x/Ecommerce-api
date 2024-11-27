package models

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	Name   string `bson:"name" json:"name" validate:"required"`
    Price int `bson:"price" json:"price" validate:"required"`
	Quantity int `bson:"quantity" json:"quantity" validate:"required"`
	ImgSrc   string `bson:"imgSrc" json:"imgSrc" validate:"required"`
	Description   string `bson:"description" json:"description" validate:"required"`
	Category   string `bson:"category" json:"category" validate:"required"`
    OwnerId  uuid.UUID `bson:"ownerId" json:"ownerId" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}