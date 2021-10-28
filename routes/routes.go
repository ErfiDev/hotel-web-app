package routes

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	server := http.NewServeMux()

	server.HandleFunc("/" , controllers.Repo.Home)
	server.HandleFunc("/about" , controllers.Repo.About)
	server.Handle("/public/" , http.StripPrefix("/public/" , http.FileServer(http.Dir("./public"))))
	server.HandleFunc("/middleware" , YourMiddleware(controllers.Repo.Middleware))

	return server
}
