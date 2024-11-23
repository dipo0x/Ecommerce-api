package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"ecommerce-api/config"
	"ecommerce-api/models"
	"ecommerce-api/helpers"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return helpers.RespondWithError(c, fiber.StatusUnauthorized, "Missing Authorization")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return helpers.RespondWithError(c, fiber.StatusUnauthorized, "Invalid token format")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			
			return []byte(config.Config("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			return helpers.RespondWithError(c, fiber.StatusUnauthorized, "Invalid token")
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["userId"].(string)
			parsedID, _ := uuid.Parse(userId)
			err = config.MongoDatabase.Collection("users").FindOne(context.Background(), bson.M{"_id": parsedID}).Decode(&user)
			if err == nil {
				c.Locals("user", user)
			}
		} else {
			return helpers.RespondWithError(c, fiber.StatusUnauthorized, "Invalid token claims")
		}
		return c.Next()
	}
}
