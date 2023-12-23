package users

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
	var req = UserRequest{}
	userId := ctx.Locals("user_id")
	userRole := ctx.Locals("user_role")

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.BadRequest(ctx, infrafiber.Response{
			Message:   "bad request",
			Error:     err.Error(),
			ErrorCode: "40000",
		})
	}
	if userRole != "user" {
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
		case GenderIsRequired:
			errorCode = "40001"
		case GenderIsInvalid:
			errorCode = "40002"
		case PhoneNumberIsEmpty:
			errorCode = "40003"
		case PhoneNumberLength:
			errorCode = "40004"
		case NameIsRequired:
			errorCode = "40005"
		case AddressIsRequired:
			errorCode = "40006"
		case DateOfBirthIsRequired:
			errorCode = "40007"
		case DateOfBirthIsInvalid:
			errorCode = "40008"
		case ImageUrlIsRequird:
			errorCode = "40009"
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
		Message:    "create user success",
	})
}

func (h handler) getProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id")

	user, err := h.svc.getProfile(ctx.UserContext(), userId.(int))
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
		Message:    "get user success",
		Payload: UserResponse{
			Name:        user.Name,
			DateOfBirth: user.DateOfBirth,
			PhoneNumber: user.PhoneNumber,
			Gender:      user.Gender,
			Address:     user.Address,
			ImageUrl:    user.ImageUrl,
		},
	})
}

func (h handler) update(ctx *fiber.Ctx) error {
	var req = UserRequest{}
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
		case GenderIsRequired:
			errorCode = "40001"
		case GenderIsInvalid:
			errorCode = "40002"
		case PhoneNumberIsEmpty:
			errorCode = "40003"
		case PhoneNumberLength:
			errorCode = "40004"
		case NameIsRequired:
			errorCode = "40005"
		case AddressIsRequired:
			errorCode = "40006"
		case DateOfBirthIsRequired:
			errorCode = "40007"
		case DateOfBirthIsInvalid:
			errorCode = "40008"
		case ImageUrlIsRequird:
			errorCode = "40009"
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
		Message:    "update role success",
	})
}
