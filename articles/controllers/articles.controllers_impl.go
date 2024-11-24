package articles_controllers

import (
	wi "blog/provider/interface"
	r "blog/provider/response"
	s "blog/provider/storage"
	cu "blog/utility"

	"github.com/gofiber/fiber/v2"
)

type ArticleControllersImpl struct{}

func NewArticlesControllers() ArticleControllers {
	return &ArticleControllersImpl{}
}

func (article *ArticleControllersImpl) GetDetailArticle(c *fiber.Ctx) error {
	articleSlug := c.Params("slug")
	value, exist := (*s.ArticlesContent)[articleSlug]
	if !exist {
		return &r.WebError{
			Code:    400,
			Message: "Gagal meneumakan artikel",
		}
	}

	data := wi.WebInterface{
		Markdown: value.Html,
		Styles:   []string{*cu.BaseCss},
	}

	return c.Render("detail.article", data)
}

func (article *ArticleControllersImpl) GetAllArticles(c *fiber.Ctx) error {
	resp := r.ApiResponse{
		Success: true,
		Message: "Berhasil Mendapatkan Data Article",
		Data:    *s.ArticlesStorage,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
