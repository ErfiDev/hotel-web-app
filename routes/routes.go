package routes

import (
	"net/http"

	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	router.Get("/", controllers.Repo.Home)
	router.Get("/about", controllers.Repo.About)
	for _, route := range []string{"/rooms", "/rooms/generals", "/rooms/majors"} {
		router.Get(route, controllers.Repo.Rooms)
	}
	router.Get("/book-now", controllers.Repo.BookNow)
	router.Get("/contact", controllers.Repo.Contact)
	router.Get("/make-reservation", controllers.Repo.MakeReservation)
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	router.Get("/reservation-summary", controllers.Repo.ReservationSummary)
	router.Get("/choose-room/{id}", controllers.Repo.ChooseRoom)
	router.Get("/login", controllers.Repo.Login)
	router.Get("/logout", controllers.Repo.Logout)
	router.Route("/admin", func(mux chi.Router) {
		// mux.Use(Authenticate)
		mux.Get("/", controllers.Repo.Admin)
		mux.Get("/dashboard", controllers.Repo.AdminDashboard)
		mux.Get("/reservations", controllers.Repo.AdminReservations)
		mux.Get("/newReservations", controllers.Repo.NewReservations)
		mux.Get("/reservation/{id}", controllers.Repo.SingleReservation)
		mux.Post("/api/updateReservation", controllers.Repo.ApiUpdateReservation)
	})

	// POST routes
	router.Post("/book-now", controllers.Repo.BookNowPost)
	router.Post("/make-reservation", controllers.Repo.MakeReservationPost)
	router.Post("/search-availability", controllers.Repo.SearchAvailability)
	router.Post("/user/login", controllers.Repo.LoginPost)

	// Api GET routes
	router.Get("/api/allReservations", controllers.Repo.AllReservationsApi)

	// Custom 404 page
	router.NotFound(NotFound)

	return router
}

func NotFound(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res, req, "404.page.gohtml", &models.TmpData{
		Data: map[string]interface{}{
			"title": "page not found",
			"path":  "/404",
		},
	})
}
