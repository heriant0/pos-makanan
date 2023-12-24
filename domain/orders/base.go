package orders

import (
	"github.com/gofiber/fiber/v2"
	paymentgateway "github.com/heriant0/pos-makanan/external/payment-gateway"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(router fiber.Router, pg *sqlx.DB, mongo *mongo.Client, xendit paymentgateway.Xendit) {
	orderRepo := newRespository(mongo)
	paymentRepo := newPaymentRepo(xendit)
	productRepo := newProductRepo(pg)

	svc := newService(orderRepo, paymentRepo, productRepo)
	handler := newHandler(svc)

	r := router.Group("orders")
	{
		r.Post("/", handler.createOrder)
		r.Post("/webhook", handler.webhook)
	}
}
