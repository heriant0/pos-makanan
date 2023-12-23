package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	infrafiber "github.com/heriant0/pos-makanan/infra/fiber"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) register(ctx *fiber.Ctx) error {
	var req = AuthRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40000",
		})
	}

	err := h.svc.register(ctx.UserContext(), req)
	if err != nil {
		errorCode := "40000"
		switch err {
		case EmailIsRequired:
			errorCode = "40001"
		case EmailIsInvalid:
			errorCode = "40003"
		case PasswordIsEmpty:
			errorCode = "40004"
		case PasswordLength:
			errorCode = "40005"
		}

		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: errorCode,
		})
	}

	return infrafiber.OutputJson(ctx, infrafiber.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "resgistration success",
	})
}

func (h handler) login(ctx *fiber.Ctx) error {
	var req = AuthRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40000",
		})
	}

	tokenAuth, err := h.svc.login(ctx.UserContext(), req)
	if err != nil {
		errorCode := "40000"
		switch err {
		case EmailIsRequired:
			errorCode = "40001"
		case EmailIsInvalid:
			errorCode = "40003"
		case PasswordIsEmpty:
			errorCode = "40004"
		case PasswordLength:
			errorCode = "40005"
		}

		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: errorCode,
		})
	}

	return infrafiber.OutputJson(ctx, infrafiber.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Payload:    tokenAuth,
		Message:    "login success",
	})
}

func (h handler) updateRole(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id")
	userRole := ctx.Locals("user_role")

	if userRole.(string) == "merchant" {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     "Unauthorized",
			ErrorCode: "40001",
		})
	}

	err := h.svc.update(ctx.UserContext(), userId.(int))
	if err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40001",
		})
	}

	return infrafiber.OutputJson(ctx, infrafiber.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "update role success",
	})
}
