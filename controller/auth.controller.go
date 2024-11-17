package controller

import (
	"context"
	"ecommerce-api/config"
	"ecommerce-api/models"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"ecommerce-api/helpers"
	"ecommerce-api/utils"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	var auth models.Auth

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	password, ok := body["password"].(string)
	if !ok || password == "" {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Password is required")
	}

	if err := c.BodyParser(&user); err != nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Invalid user data")
	}
	userCollection := config.MongoDatabase.Collection("users")

	filter := bson.M{"username": user.Username}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "User already exists")
	}

	password, _ = utils.HashPassword(password)

	auth.ID = uuid.New()
	auth.Password = password
	auth.CreatedAt = time.Now()
	auth.UpdatedAt = time.Now()

	authCollection := config.MongoDatabase.Collection("auths")
	_, err = authCollection.InsertOne(context.Background(), auth)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusInternalServerError, "Failed to save auth details")
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.AuthId = auth.ID

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusInternalServerError, "Failed to save user details")
	}

	update := bson.M{
		"$set": bson.M{
			"userId": user.ID,
		},
	}

	filter = bson.M{"_id": auth.ID}

	_, err = authCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Failed to create user")
	}

	return helpers.RespondWithSuccess(c, fiber.StatusCreated, user)
}
