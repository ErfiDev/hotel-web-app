package routes

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/go-chi/chi"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	// Middlewares
	router.Use(YourMiddleware)

	router.Get("/" , controllers.Repo.Home)
	router.Get("/about" , controllers.Repo.About)
	router.Handle("/public" , http.StripPrefix("/public/" , http.FileServer(http.Dir("./public"))))

	return router
}
