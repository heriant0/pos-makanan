package merchants

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/pos-makanan/internal/middlewares"
	"github.com/jmoiron/sqlx"
)

func InitRouter(router fiber.Router, db *sqlx.DB) {
	repository := newRepository(db)
	service := newService(repository)
	handler := newHandler(service)

	users := router.Group("merchants")
	{
		users.Post("/profile", middlewares.AuthMiddlware, handler.create)
		users.Get("/profile", middlewares.AuthMiddlware, handler.getProfile)
		users.Put("/profile", middlewares.AuthMiddlware, handler.update)
	}
}
