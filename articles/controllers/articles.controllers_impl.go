package articles_controllers

import (
	"blog/provider/converter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ArticleControllersImpl struct{}

func NewArticlesControllers() ArticleControllers {
	return &ArticleControllersImpl{}
}

func (article *ArticleControllersImpl) GetDetailArticle(c *fiber.Ctx) error {
	articleName := c.Params("slug")
	articleName += ".md"
	html, err := converter.MarkdownConverter(articleName)
	if err != nil {
		log.Error("Gagal melakukan konversi markdown ke html. Detail: " + err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal melakukan konversi")
	}

	return c.Render("detail.article", fiber.Map{"embed": html})
}
