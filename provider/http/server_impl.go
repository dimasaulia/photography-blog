package http

import (
	r "blog/provider/response"
	u "blog/utility"
	"fmt"
	"strings"

	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type HttpServerImpl struct {
	Port         string
	RenderEngine fiber.Views
	Fork         bool
}

type Server struct {
	Port         string
	RenderEngine fiber.Views
	PreFork      bool
}

func NewServer(c Server) HttpServer {
	return HttpServerImpl{
		Port:         c.Port,
		RenderEngine: c.RenderEngine,
		Fork:         c.PreFork,
	}
}

func (config HttpServerImpl) Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		Views:   config.RenderEngine,
		Prefork: config.Fork,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			fmt.Println(err.Error())

			if strings.Contains(err.Error(), "Cannot GET") && err.Error() != "Cannot GET /favicon.ico" {
				return c.Render("errors", fiber.Map{
					"styles": []string{*u.BaseCss},
					"errors": "Halaman yang anda cari sepertinya ikut tenggelam bersama atlantis",
					"code":   404,
				})
			}

			var webError *r.WebError
			if errors.As(err, &webError) {
				return c.Render("errors", fiber.Map{
					"styles": []string{*u.BaseCss},
					"errors": err.Error(),
					"code":   webError.Code,
				})
			}

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// logFile, err := os.OpenFile("./app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening request log file: %v", err)
	// }
	// multiLogFile := io.MultiWriter(os.Stdout, logFile)

	app.Use(requestid.New())
	// app.Use(logger.New(logger.Config{
	// 	// For more options, see the Config section
	// 	Format:     "${time} ${ip}:${port} ${status} - ${method} ${path}\n",
	// 	TimeZone:   "Asia/Jakarta",
	// 	TimeFormat: "2006/01/02 15:04:05.000000",
	// 	Output:     multiLogFile,
	// }))

	app.Static("/public", "./public")

	return app
}

func (config HttpServerImpl) StartListening(app *fiber.App) {
	port := fmt.Sprintf(":%s", config.Port)
	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
