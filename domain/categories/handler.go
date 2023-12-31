package categories

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) list(ctx *fiber.Ctx) error {
	categoryList, err := h.svc.GetAll(ctx.UserContext())
	if err != nil {
		log.Error(fmt.Errorf("error handler - list: %w", err))

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message":       "unknown error",
			"error_code":    "999999",
			"error_message": err.Error(),
		})
	}

	var result = []CategoryResponse{}

	for _, category := range categoryList {
		result = append(result, category.parseToCategoryResponse())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "get all category success",
		"payload": result,
	})
}
