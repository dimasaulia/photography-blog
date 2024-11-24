package home_controllers

import "github.com/gofiber/fiber/v2"

type HomeControllers interface {
	Home(c *fiber.Ctx) error
}
