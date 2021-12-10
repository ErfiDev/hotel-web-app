package controllers

import (
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexedwards/scs/v2"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)

var funcMap template.FuncMap = template.FuncMap{
	"ToHumanDate":    utils.ToHumanDate,
	"GetDayFromDate": utils.GetDayFromDate,
}
var appConfig config.AppConfig
var sessionManager *scs.SessionManager
var InfoLog, ErrorLog *log.Logger

func GetRoutes() http.Handler {
	// Register value and type into encoding/Gob .Register()
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.RoomRestriction{})
	gob.Register(models.Restriction{})
	gob.Register(models.BookNow{})

	// create template caches
	tmpCache, errCache := CreateTestTemplateCache()
	if errCache != nil {
		log.Fatal("can't create template cache")
	}

	// init AppConfig tmpCache
	appConfig.TemplatesCache = tmpCache

	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	appConfig.ErrorLog = ErrorLog
	appConfig.InfoLog = InfoLog

	// init session manager
	sessionManager = scs.New()
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	appConfig.Session = sessionManager

	utils.GetAppConfig(&appConfig)
	SetRepo(NewTestRepository(&appConfig))

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Recoverer)
	// because we in testing and not have the CSRF_TOKEN in the browser session
	// router.Use(NoSurf)
	router.Use(ServeSession)

	router.Get("/", Repo.Home)
	router.Get("/about", Repo.About)
	for _, route := range []string{"/rooms", "/rooms/generals", "/rooms/majors"} {
		router.Get(route, Repo.Rooms)
	}
	router.Get("/book-now", Repo.BookNow)
	router.Get("/contact", Repo.Contact)
	router.Get("/make-reservation", Repo.MakeReservation)
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	router.Get("/reservation-summary", Repo.ReservationSummary)

	// POSTS
	router.Post("/book-now", Repo.BookNowPost)
	router.Post("/make-reservation", Repo.MakeReservationPost)
	router.Post("/search-availability", Repo.SearchAvailability)

	return router
}

func NoSurf(next http.Handler) http.Handler {
	CSRFHandler := nosurf.New(next)

	CSRFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   !appConfig.Development,
	})

	return CSRFHandler
}

func ServeSession(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	caches := map[string]*template.Template{}

	pages, err := filepath.Glob("./../views/*.page.gohtml")
	// [$../views/about.page.gohtml  &...]
	if err != nil {
		return caches, err
	}

	// [$../views/about.page.gohtml]
	for _, page := range pages {
		name := filepath.Base(page)
		// [about.page.gohtml]

		tmp, errNewTmp := template.New(name).Funcs(funcMap).ParseFiles(page)
		if errNewTmp != nil {
			return caches, errNewTmp
		}

		findLayout, _ := filepath.Glob("./../views/*.layout.gohtml")
		if len(findLayout) > 0 {
			if strings.Contains(name, "admin") {
				tmp.ParseGlob("./../views/admin.layout.gohtml")
			} else {
				tmp.ParseGlob("./../views/base.layout.gohtml")
			}
		}

		caches[name] = tmp
	}
	return caches, nil
}
