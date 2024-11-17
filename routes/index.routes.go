package routes

import (
	"github.com/gofiber/fiber/v2"
    "ecommerce-api/controller"
)

func IndexRoutes(router fiber.Router) {
    router.Get("/", controller.Index)
}