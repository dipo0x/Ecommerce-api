package routes

import (
	"github.com/gofiber/fiber/v2"
    "ecommerce-api/controller"
    "ecommerce-api/middleware"
	"ecommerce-api/types"
)

func AuthRoutes(router fiber.Router) {
    router.Post("/sign-up",  middleware.ValidateStruct(new(types.IUser)), controller.CreateUser)
    router.Post("/login",  middleware.ValidateStruct(new(types.ILogin)), controller.LoginUser)
}