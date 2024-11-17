package controller

import (
	"context"
	"ecommerce-api/config"
	"ecommerce-api/models"
	"time"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	var auth models.Auth

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"error": "Invalid request payload"})
	}

	password := c.FormValue("password")
	auth.ID = uuid.New()
	auth.Password = password

	authCollection := config.MongoDatabase.Collection("auths")
	_, err := authCollection.InsertOne(context.Background(), auth)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"error":  "Failed to save auth details",
		})
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.AuthId = auth.ID

	userCollection := config.MongoDatabase.Collection("users")
	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"success": false,
			"error":  "Failed to save user details",
		})
	}

	update := bson.M{
        "$set": bson.M{
            "userId":  user.ID,
        },
    }

	filter := bson.M{"_id": auth.ID}

	_, err = authCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"success": false,
			"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": 201,
		"success": true,
		"data": user,
	})
}
