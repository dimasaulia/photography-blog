package home_controllers

import (
	wi "blog/provider/interface"
	u "blog/utility"

	"github.com/gofiber/fiber/v2"
)

type HomeControllersImpl struct{}

func NewHomeControllers() HomeControllers {
	return &HomeControllersImpl{}
}

func (hci *HomeControllersImpl) Home(c *fiber.Ctx) error {
	data := wi.WebInterface{
		Styles: []string{*u.BaseCss},
	}
	return c.Render("home", data)
}
