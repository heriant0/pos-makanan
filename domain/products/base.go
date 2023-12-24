package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/pos-makanan/internal/middlewares"
	"github.com/jmoiron/sqlx"
)

func InitRouter(router fiber.Router, db *sqlx.DB) {
	repository := newRepository(db)
	service := newService(repository)
	handler := newHandler(service)

	products := router.Group("products")
	{
		products.Post("/", middlewares.AuthMiddlware, handler.create)
		// produts.Get("/", middlewares.AuthMiddlware, handler.getProfile)
		// produts.Put("/", middlewares.AuthMiddlware, handler.update)
	}
}
