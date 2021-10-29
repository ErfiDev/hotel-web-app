package main

import (
	"fmt"
	"github.com/alexedwards/scs"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/erfidev/hotel-web-app/routes"
	"github.com/erfidev/hotel-web-app/utils"
	"log"
	"net/http"
	"os"
	"time"
)

// Global variables
var appConfig = config.AppConfig{}
var Store scs.Store

func main() {
	// create template caches
	tmpCache , errCache := utils.CreateTemplateCache()
	if errCache != nil {
		log.Fatal("can't create template cache")
	}

	// init AppConfig tmpCache
	appConfig.TemplatesCache = tmpCache

	if len(os.Args) > 1 {
		secondArgs := os.Args[1]
		if secondArgs == "production" {
			appConfig.Development = false
		}
	} else {
		appConfig.Development = true
	}

	// init session manager
	sessionManager := scs.NewManager(Store)
	sessionManager.Lifetime(24 * time.Hour)
	sessionManager.Persist(true)
	sessionManager.Secure(!appConfig.Development)

	utils.GetAppConfig(&appConfig)
	controllers.SetRepo(controllers.NewRepository(&appConfig))
	routes.SetAppConfig(&appConfig)
	// server initializing

	routeHandler := routes.Routes()
	webServer := &http.Server{
		Addr: ":3000",
		Handler: routeHandler,
	}

	err := webServer.ListenAndServe()
	if err != nil {
		log.Fatal("error on ListenAndServe")
	}
	fmt.Println("we on port :3000")
}
