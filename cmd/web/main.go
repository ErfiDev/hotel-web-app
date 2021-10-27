package main

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/erfidev/hotel-web-app/routes"
	"github.com/erfidev/hotel-web-app/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	// create template caches
	tmpCache , errCache := utils.CreateTemplateCache()
	if errCache != nil {
		log.Fatal("can't create template cache")
	}

	// init AppConfig
	appConfig := config.AppConfig{}
	appConfig.TemplatesCache = tmpCache

	if len(os.Args) > 1 {
		secondArgs := os.Args[1]
		if secondArgs == "production" {
			appConfig.Development = false
		}
	} else {
		appConfig.Development = true
	}

	// Send to the getAppConfig package
	utils.GetAppConfig(&appConfig)
	controllers.SetRepo(controllers.NewRepository(&appConfig))

	// server initializing
	routeHandler := routes.Routes(&appConfig)
	webServer := &http.Server{
		Addr: ":3000",
		Handler: routeHandler,
	}

	err := webServer.ListenAndServe()
	if err != nil {
		log.Fatal("error on ListenAndServe")
	}
}
