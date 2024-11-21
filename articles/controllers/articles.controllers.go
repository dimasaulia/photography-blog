package articles_controllers

import "github.com/gofiber/fiber/v2"

type ArticleControllers interface {
	GetDetailArticle(c *fiber.Ctx) error
}
