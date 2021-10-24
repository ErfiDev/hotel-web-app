package main

import (
	"fmt"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/erfidev/hotel-web-app/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	muxServer := http.NewServeMux()

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
		switch secondArgs {
		case "development":
			appConfig.Development = true

		case "production":
			appConfig.Development = false

		default:
			appConfig.Development = true
		}
	} else {
		appConfig.Development = true
	}

	// Send to the getAppConfig package
	utils.GetAppConfig(&appConfig)


	muxServer.HandleFunc("/" , controllers.Home)
	muxServer.HandleFunc("/about" , controllers.About)
	muxServer.Handle("/public/" , http.StripPrefix("/public/" , http.FileServer(http.Dir("./public"))))

	err := http.ListenAndServe(":3000" , muxServer)
	if err != nil {
		fmt.Println("we have the fucking error on starting server")
		return
	}
	fmt.Println("run on :3000 port")
}
