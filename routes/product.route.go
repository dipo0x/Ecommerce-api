package routes

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
    router.Post("/create", middleware.ValidateToken(), controller.CreateProduct)
}