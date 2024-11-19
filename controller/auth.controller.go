package controller

import (
	"context"
	"ecommerce-api/config"
	"ecommerce-api/helpers"
	"ecommerce-api/models"
	"ecommerce-api/types"
	"ecommerce-api/utils"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.M{"username": user.Username}
	err := config.MongoDatabase.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "User already exists")
	}

	password, _ = utils.HashPassword(password)

	auth.ID = uuid.New()
	auth.Password = password
	auth.CreatedAt = time.Now()
	auth.UpdatedAt = time.Now()

	_, err = config.MongoDatabase.Collection("auths").InsertOne(context.Background(), auth)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusInternalServerError, "Failed to save auth details")
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.AuthId = auth.ID

	_, err = config.MongoDatabase.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusInternalServerError, "Failed to save user details")
	}

	update := bson.M{
		"$set": bson.M{
			"userId": user.ID,
		},
	}

	filter = bson.M{"_id": auth.ID}

	_, err = config.MongoDatabase.Collection("auths").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Failed to create user")
	}

	return helpers.RespondWithSuccess(c, fiber.StatusCreated, user)
}

func LoginUser(c *fiber.Ctx) error {
	var auth types.IAuth
	var requestAuth types.IAuth
	var user models.User

	if err := c.BodyParser(&requestAuth); err != nil {
		println(err.Error())
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	err := config.MongoDatabase.Collection("users").FindOne(context.Background(), bson.M{"username": requestAuth.Username}).Decode(&user)
	if err != nil{
		return helpers.RespondWithError(c, fiber.StatusNotFound, "User not found")
	}

	_ = config.MongoDatabase.Collection("auths").FindOne(context.Background(), bson.M{"userId": user.ID}).Decode(&auth)

    err = utils.CheckPasswordHash(auth.Password, requestAuth.Password)
	if err != nil {
        return helpers.RespondWithError(c, fiber.StatusUnauthorized, "Invalid credentials")
    }
	token := utils.GenerateToken(user.ID.String())

	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	
	return helpers.RespondWithSuccess(c, fiber.StatusOK, data)
}