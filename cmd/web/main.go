package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/routes"
	"github.com/erfidev/hotel-web-app/utils"
	"log"
	"net/http"
	"os"
)

// Global variables
var appConfig = config.AppConfig{}
var sessionManager *scs.SessionManager

func main() {
	// Register value and type into encoding/Gob .Register()
	gob.Register(models.Reservation{})

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
	sessionManager = scs.New()
	sessionManager.Cookie.Secure = !appConfig.Development
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	appConfig.Session = sessionManager


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
