package main

import (
	"ecommerce-api/config"
	"ecommerce-api/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main () {

	app := fiber.New()
	
	routes.IndexRoutes(app.Group("/"))
	routes.AuthRoutes(app.Group("/auth"))
	routes.ProductRoutes(app.Group("/product"))
	routes.OrderRoutes(app.Group("/order"))
	routes.UserRoutes(app.Group("/user"))

	err:= config.InitializeMongoDB(config.Config("MONGO_URI"), config.Config("MONGO_DATABASE"))

	if err != nil {
		defer config.DisconnectMongoDB()
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	port := config.Config("PORT")
	log.Fatal(app.Listen(port))
}