package routes

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/utils"
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
	for _ , route := range []string{"/rooms" , "/rooms/generals" , "/rooms/majors"} {
		router.Get(route , controllers.Repo.Rooms)
	}
	router.Get("/book-now" , controllers.Repo.BookNow)
	router.Get("/contact" , controllers.Repo.Contact)
	router.Get("/make-reservation" , controllers.Repo.MakeReservation)
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*" , http.StripPrefix("/static" , fileServer))
	router.Get("/reservation-summary" , controllers.Repo.ReservationSummary)
	router.Get("/choose-room/{id}" , controllers.Repo.ChooseRoom)

	// POST routes
	router.Post("/book-now" , controllers.Repo.BookNowPost)
	router.Post("/make-reservation" , controllers.Repo.MakeReservationPost)
	router.Post("/search-availability" , controllers.Repo.SearchAvailability)

	// Custom 404 page
	router.NotFound(NotFound)

	return router
}

func NotFound(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , req , "404.page.gohtml" , &models.TmpData{
		Data: map[string]interface{}{
			"title": "page not found",
			"path": "/404",
		},
	})
}