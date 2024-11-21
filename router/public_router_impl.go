package router

import (
	articles_controllers "blog/articles/controllers"
	articles_routes "blog/articles/routes"

	"github.com/gofiber/fiber/v2"
)

type PublicRoutersImpl struct {
	App *fiber.App
}

func NewRouters(app *fiber.App) PublicRouters {
	return &PublicRoutersImpl{
		App: app,
	}
}

func (r PublicRoutersImpl) Setup() {
	articleControllers := articles_controllers.NewArticlesControllers()
	articles_routes.SetupArticlesRoutes(r.App.Group("/articles"), articleControllers)
}
