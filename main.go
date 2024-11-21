package main

import (
	"blog/provider/config"
	"blog/provider/http"
	"blog/router"
	cu "blog/utility"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	env := config.ReadConfigigurationFile()
	port := env.GetString("PORT")
	PROD_CSS := "input"
	DEV_CSS := "output"
	// var multiLogFile io.Writer
	if env.GetString("ENV") == "PROD" {
		// multiLogFile = io.MultiWriter(logFile)
		cu.BaseCss = &PROD_CSS
	} else {
		// multiLogFile = io.MultiWriter(os.Stdout, logFile)
		cu.BaseCss = &DEV_CSS
	}
	// log.SetOutput(multiLogFile)

	engine := handlebars.New("./views/default", ".hbs")
	preFork := env.GetBool("FORK")

	server := http.NewServer(http.Server{
		Port:         port,
		RenderEngine: engine,
		PreFork:      preFork,
	})
	app := server.Setup()

	// ROUTERS CONFIGURATION
	routers := router.NewRouters(app)
	routers.Setup()

	server.StartListening(app)
}
