package routes

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
    router.Get("/get", middleware.ValidateToken(), controller.GetUser)
}