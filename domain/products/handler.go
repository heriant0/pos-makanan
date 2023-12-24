package products

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

func (h handler) create(ctx *fiber.Ctx) error {
	var req = ProductRequest{}
	userId := ctx.Locals("user_id")
	userRole := ctx.Locals("user_role")

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40000",
		})
	}

	if userRole != "merchant" {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     ErrInvalidRole.Error(),
			ErrorCode: "40102",
		})
	}

	err := h.svc.create(ctx.UserContext(), req, userId.(int))
	if err != nil {
		errorCode := "40000"
		switch err {
		case ErrPriceIsRequired:
			errorCode = "40001"
		case ErrPriceIsInvalid:
			errorCode = "40002"
		case ErrStockIsRequired:
			errorCode = "40003"
		case ErrStockIsInvalid:
			errorCode = "40004"
		case ErrNameIsRequired:
			errorCode = "40005"
		case ErrDescriptionIsRequierd:
			errorCode = "40006"
		case ErrImageUrlIsRequierd:
			errorCode = "40007"
		case ErrCategoryIdIsRequired:
			errorCode = "40009"
		case ErrCategoryIdIsNotFound:
			errorCode = "40010"
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
		Message:    "create product success",
	})
}
