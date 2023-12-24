package orders

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	infrafiber "github.com/heriant0/pos-makanan/infra/fiber"
	"github.com/heriant0/pos-makanan/internal/config"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{svc}
}

func (h handler) createOrder(ctx *fiber.Ctx) (err error) {
	req := createOrderRequest{}
	ctx.BodyParser(&req)

	res, err := h.svc.createOrder(ctx.Context(), req)
	if err != nil {
		ctx.JSON(infrafiber.Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Error:      err.Error(),
		})
		return
	}

	ctx.JSON(infrafiber.Response{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "order success",
		Payload:    res,
	})

	return
}

func (h handler) webhook(ctx *fiber.Ctx) (err error) {
	headerXCallback := ctx.GetReqHeaders()["X-Callback-Token"]
	callbackToken := config.Cfg.Payment.CallbackToken

	if headerXCallback != nil && len(headerXCallback) == 0 {
		ctx.JSON(infrafiber.Response{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Error:      "Request is invalid",
		})
		return
	}

	if headerXCallback[0] != callbackToken {
		ctx.JSON(infrafiber.Response{
			StatusCode: http.StatusForbidden,
			Success:    false,
			Error:      "You are not allowed",
		})
		return
	}

	req := updateOrderRequest{}
	ctx.BodyParser(&req)

	err = h.svc.changeOrderStatus(ctx.Context(), req)
	return
}
