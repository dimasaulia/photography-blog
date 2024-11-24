package home_routes

import (
	cont "blog/home/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupHomeRoutes(r fiber.Router, cont cont.HomeControllers) {
	r.Get("/", cont.Home)
}
