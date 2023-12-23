package middlewares

import (
	"github.com/gofiber/fiber/v2"
	infrafiber "github.com/heriant0/pos-makanan/infra/fiber"
	"github.com/heriant0/pos-makanan/utility"
)

func AuthMiddlware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authentication")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// verify token
	claims, err := utility.VerifyToken(token)
	if err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40101",
		})
	}

	payload := claims["payload"].(map[string]interface{})
	id := payload["id"].(float64)
	role := payload["role"].(string)

	ctx.Locals("user_id", int(id))
	ctx.Locals("user_role", role)

	return ctx.Next()
}
