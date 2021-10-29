package routes

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

var appConfig *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	appConfig = a
}

func Routes() http.Handler {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(ServeSession)

	router.Get("/" , controllers.Repo.Home)
	router.Get("/about" , controllers.Repo.About)
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*" , http.StripPrefix("/static" , fileServer))


	return router
}