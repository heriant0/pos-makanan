package middlewares

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func LogMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		clientIP := ctx.IP()
		method := ctx.Method()
		url := ctx.Path()

		log.WithFields(log.Fields{
			"client_id": clientIP,
			"url":       url,
			"method":    method,
		}).Info("http request")

		return ctx.Next()
	}
}
