package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/erfidev/hotel-web-app/driver"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/routes"
	"github.com/erfidev/hotel-web-app/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// Global variables
var appConfig = config.AppConfig{}
var sessionManager *scs.SessionManager
var InfoLog *log.Logger
var ErrorLog *log.Logger

func main() {
	_ = godotenv.Load()

	db , err := InitProject()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	defer close(appConfig.MailChan)

	Listener()

	routeHandler := routes.Routes()
	webServer := &http.Server{
		Addr: ":3000",
		Handler: routeHandler,
	}

	fmt.Println("we on port :3000")
	err = webServer.ListenAndServe()
	if err != nil {
		log.Fatal("error on ListenAndServe")
	}
}

func InitProject() (*driver.DB , error) {
	// Register value and type into encoding/Gob .Register()
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.BookNow{})
	gob.Register(models.RoomRestriction{})
	gob.Register(models.Restriction{})
	gob.Register(models.User{})

	mailChannel := make(chan models.MailData)
	appConfig.MailChan = mailChannel

	// create template caches
	tmpCache , errCache := utils.CreateTemplateCache()
	if errCache != nil {
		log.Fatal("can't create template cache")
	}

	// init AppConfig tmpCache
	appConfig.TemplatesCache = tmpCache

	InfoLog = log.New(os.Stdout , "INFO\t" , log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stdout , "ERROR\t" , log.Ldate|log.Ltime|log.Lshortfile)

	appConfig.ErrorLog = ErrorLog
	appConfig.InfoLog = InfoLog

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

	// Connecting to database
	log.Println("Connecting to database...")
	db , err := driver.ConnectDB(`host=localhost port=5432 dbname=hotel user=postgres password=28963323`)
	if err != nil {
		return nil , err
	}

	utils.GetAppConfig(&appConfig)
	controllers.SetRepo(controllers.NewRepository(&appConfig , db))
	routes.SetAppConfig(&appConfig)
	// server initializing
	return db , nil
}