package articles_controllers

import (
	"blog/provider/converter"
	wi "blog/provider/interface"
	cu "blog/utility"

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
	css, _ := cu.LoadCss(*cu.BaseCss)

	data := wi.WebInterface{
		Markdown: html,
		Styles:   []string{css},
	}

	return c.Render("detail.article", data)
}
