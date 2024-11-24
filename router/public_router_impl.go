package router

import (
	articles_controllers "blog/articles/controllers"
	articles_routes "blog/articles/routes"
	home_controllers "blog/home/controllers"
	home_routes "blog/home/routes"

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
	homeControllers := home_controllers.NewHomeControllers()
	articleControllers := articles_controllers.NewArticlesControllers()

	home_routes.SetupHomeRoutes(r.App.Group("/"), homeControllers)
	articles_routes.SetupArticlesRoutes(r.App.Group("/articles"), articleControllers)
}
