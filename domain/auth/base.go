package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/pos-makanan/internal/middlewares"
	"github.com/jmoiron/sqlx"
)

func InitRouter(router fiber.Router, db *sqlx.DB) {
	repository := newRepository(db)
	service := newService(repository)
	handler := newHandler(service)

	auth := router.Group("auth")
	{
		auth.Post("/register", handler.register)
		auth.Post("/login", handler.login)
		auth.Patch("/role", middlewares.AuthMiddlware, handler.updateRole)
	}
}
