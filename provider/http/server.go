package http

import (
	"github.com/gofiber/fiber/v2"
)

type HttpServer interface {
	Setup() *fiber.App
	StartListening(app *fiber.App)
}
