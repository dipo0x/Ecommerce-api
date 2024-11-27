package types

type IProduct struct {
    Name   string `json:"name" validate:"required"`
    Price int `json:"price" validate:"required"`
    Quantity int `json:"quantity" validate:"required"`
    ImgSrc   string `json:"ImgSrc" validate:"required"`
    Description   string `json:"Description" validate:"required"`
    Category   string `json:"category" validate:"required"`
}