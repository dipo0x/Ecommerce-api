package routes

import (
	"ecommerce-api/controller"
	"ecommerce-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router fiber.Router) {
	router.Get("/place-order/:productId/:quantity", middleware.ValidateToken(), controller.PlaceOrder)
}