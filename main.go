package main

import (
	"blog/provider/config"
	"blog/provider/http"
	s "blog/provider/storage"
	"blog/router"
	u "blog/utility"
	"fmt"
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
	PROD_CSS := "final-min"
	DEV_CSS := "output"

	if env.GetString("ENV") == "PROD" {
		// multiLogFile = io.MultiWriter(logFile)
		css, err := u.LoadCss(PROD_CSS)
		if err != nil {
			fmt.Println("Gagal Load CSS Di Production")
		}
		u.BaseCss = &css
	} else {
		// multiLogFile = io.MultiWriter(os.Stdout, logFile)
		css, err := u.LoadCss(DEV_CSS)
		if err != nil {
			fmt.Println("Gagal Load CSS Di Development")
		}
		u.BaseCss = &css
	}

	// Load All Article
	s.InitArticlesStorage()
	s.InitArticlesContent()
	u.InitGoldmark()
	err = u.LoadAllData("./markdown_files")
	if err != nil {
		fmt.Println("Error Saat Melakukan Load Markdown")
		fmt.Println(err)
	}

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
