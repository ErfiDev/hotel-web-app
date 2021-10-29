package routes

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	// set up file server
	workDir , _ := os.Getwd()
	fileDirPath := http.Dir(filepath.Join(workDir , "public"))
	FileServer(router , "/public" , fileDirPath)

	return router
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}