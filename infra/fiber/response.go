package infrafiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int         `json:"-"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload,omitempty"`
	Error      string      `json:"error,omitempty"`
	ErrorCode  string      `json:"error_code,omitempty"`
}

func BadRequest(ctx *fiber.Ctx, resp Response) error {
	resp.StatusCode = http.StatusBadRequest
	return OutputJson(ctx, resp)
}

func OutputJson(ctx *fiber.Ctx, resp Response) error {
	return ctx.Status(resp.StatusCode).JSON(resp)
}
