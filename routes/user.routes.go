package routes

import (
	"github.com/gofiber/fiber/v2"
    "ecommerce-api/controller"
)

func UserRoutes(router fiber.Router) {
    router.Get("/get", controller.GetUser)
}