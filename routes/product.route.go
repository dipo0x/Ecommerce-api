package routes

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"ecommerce-api/types"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
    router.Post("/create", middleware.ValidateStruct(new(types.IProduct)), middleware.ValidateToken(), controller.CreateProduct)
	router.Get("/get", middleware.ValidateToken(), controller.GetProduct)
}