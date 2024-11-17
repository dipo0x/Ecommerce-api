package middleware

import  (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"ecommerce-api/types"
)

var Validator = validator.New()

func ValidateStruct(model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var errors []*types.IError

		if err := c.BodyParser(model); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  400,
				"success": false,
				"error":   "Invalid request payload",
			})
		}

		err := Validator.Struct(model)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				var el types.IError
				el.Field = err.Field()
				el.Tag = err.Tag()
				el.Value = err.Param()
				errors = append(errors, &el)
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  400,
				"success": false,
				"error": errors,
			})
		}

		return c.Next()
	}
}