package merchants

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
	var req = MerchantRequest{}
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
			Error:     InvalidRole.Error(),
			ErrorCode: "40000",
		})
	}

	err := h.svc.create(ctx.UserContext(), req, userId.(int))
	if err != nil {
		errorCode := "40000"
		switch err {
		case NameIsRequired:
			errorCode = "40001"
		case AddressIsRequired:
			errorCode = "40002"
		case PhoneNumberIsEmpty:
			errorCode = "40003"
		case PhoneNumberLength:
			errorCode = "40004"
		case ImageUrlIsRequird:
			errorCode = "40005"
		case CityIsRequired:
			errorCode = "40006"
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
		Message:    "create merchants success",
	})
}

func (h handler) getProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id")

	res, err := h.svc.getProfile(ctx.UserContext(), userId.(int))
	if err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40401",
		})
	}

	return infrafiber.OutputJson(ctx, infrafiber.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "get merchant success",
		Payload: MerchantResponse{
			Name:        res.Name,
			Address:     res.Address,
			PhoneNumber: res.PhoneNumber,
			City:        res.City,
			ImageUrl:    res.ImageUrl,
		},
	})
}

func (h handler) update(ctx *fiber.Ctx) error {
	var req = MerchantRequest{}
	userId := ctx.Locals("user_id")

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40000",
		})
	}

	err := h.svc.update(ctx.UserContext(), req, userId.(int))
	if err != nil {
		errorCode := "40000"
		switch err {
		case NameIsRequired:
			errorCode = "40001"
		case AddressIsRequired:
			errorCode = "40002"
		case PhoneNumberIsEmpty:
			errorCode = "40003"
		case PhoneNumberLength:
			errorCode = "40004"
		case ImageUrlIsRequird:
			errorCode = "40005"
		case CityIsRequired:
			errorCode = "40006"
		}

		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: errorCode,
		})
	}

	return infrafiber.OutputJson(ctx, infrafiber.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "update merchant success",
	})
}
