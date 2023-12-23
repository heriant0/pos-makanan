package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/heriant0/pos-makanan/utility"
)

func AuthMiddlware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authentication")
	fmt.Println("🚀 ~ file: auth_middleware.go ~ line 10 ~ funcAuthMiddlware ~ token : ", token)
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// verify token
	claims, err := utility.VerifyToken(token)
	fmt.Println("🚀 ~ file: auth_middleware.go ~ line 19 ~ funcAuthMiddlware ~ claims : ", claims)
	if err != nil {
		// return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		// 	"error": err.Error(),
		// })
		return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	ctx.Locals("user_id", claims)

	return ctx.Next()
}
