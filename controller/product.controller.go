package controller

import (
	"context"
	"ecommerce-api/config"
	"ecommerce-api/helpers"
	"ecommerce-api/models"
	"ecommerce-api/types"
	"time"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	var requestProduct types.IProduct

	if err := c.BodyParser(&requestProduct); err != nil {
		println(err.Error())
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	filter := bson.M{"name": requestProduct.Name}
	err := config.MongoDatabase.Collection("products").FindOne(context.Background(), filter).Decode(&product)
	if err == nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Product with this title already exists")
	}
	
	product.ID = uuid.New()
	product.Name = requestProduct.Name
	product.Price = requestProduct.Price
	product.Quantity = requestProduct.Quantity
	product.ImgSrc = requestProduct.ImgSrc
	product.Description = requestProduct.Description
	product.Category = requestProduct.Category

	product.OwnerId = c.Locals("user").(models.User).ID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_ , err = config.MongoDatabase.Collection("products").InsertOne(context.Background(), product)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusInternalServerError, "Failed to save product details")
	}

	return helpers.RespondWithSuccess(c, fiber.StatusCreated, product)
}

func GetProduct(c *fiber.Ctx) error {
	var products []models.Product

	cursor, err := config.MongoDatabase.Collection("products").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatalf("Failed to find products: %v", err)
		return (err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product  models.Product 
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Failed to decode product: %v", err)
			continue
		}
		products = append(products, product)
	}
	
	return helpers.RespondWithSuccess(c, fiber.StatusCreated, products)
}