package categories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func InitRouter(router fiber.Router, db *sqlx.DB) {
	repository := newRespository(db)
	service := newService(repository)
	handler := newHandler(service)

	category := router.Group("categories")
	{
		category.Get("/", handler.list)
	}
}
