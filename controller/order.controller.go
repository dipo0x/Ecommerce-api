package controller

import (
	"context"
	"ecommerce-api/config"
	"ecommerce-api/helpers"
	"ecommerce-api/models"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v81"
    "github.com/stripe/stripe-go/v81/checkout/session"
	"go.mongodb.org/mongo-driver/bson"
)
func PlaceOrder(c *fiber.Ctx) error {
	var product models.Product
	productID, _ := uuid.Parse(c.Params("productId"))
	quantity, err := strconv.Atoi(c.Params("quantity"))

	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Invalid product ID")
	}
	filter := bson.M{"_id": productID}

	err = config.MongoDatabase.Collection("products").FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusNotFound, "Product not found")
	}
	if(product.Quantity < quantity){
		return helpers.RespondWithError(c, fiber.StatusBadRequest, "Quantity ordered is greater than available stock")
	}

	stripe.Key = config.Config("STRIPE_KEY")

	///directly from stripe doc
	params := &stripe.CheckoutSessionParams{
	LineItems: []*stripe.CheckoutSessionLineItemParams{
		&stripe.CheckoutSessionLineItemParams{
		PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
			UnitAmount: stripe.Int64(0),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
			Name: stripe.String("Free t-shirt"),
			},
			Currency: stripe.String(string(stripe.CurrencyUSD)),
		},
		Quantity: stripe.Int64(1),
		},
	},
	Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	SuccessURL: stripe.String("https://example.com/success"),
	CancelURL: stripe.String("https://example.com/cancel"),
	};
	result, _ := session.New(params);


	////discontinued this project because stripe won't let me build without setting up my business account
	return helpers.RespondWithSuccess(c, fiber.StatusOK, result)
}
