package types

type IProduct struct {
    Name   string `json:"name" validate:"required"`
    Price int `json:"price" validate:"required"`
}