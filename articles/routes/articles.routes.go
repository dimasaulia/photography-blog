package articles_routes

import (
	cont "blog/articles/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupArticlesRoutes(r fiber.Router, cont cont.ArticleControllers) {
	r.Get("/list", cont.GetAllArticles)
	r.Get("/:slug", cont.GetDetailArticle)
}
