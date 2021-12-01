package routes

import (
	"github.com/erfidev/hotel-web-app/utils"
	"github.com/justinas/nosurf"
	"net/http"
)

func NoSurf(next http.Handler) http.Handler {
	CSRFHandler := nosurf.New(next)

	CSRFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
		Secure: !appConfig.Development,
	})

	return CSRFHandler
}

func ServeSession(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter , req *http.Request) {
		if !utils.IsAuthenticated(req) {
			appConfig.Session.Put(req.Context() , "error" , "Log in first!")

			http.Redirect(res , req , "/login" , http.StatusSeeOther)
			return
		}

		next.ServeHTTP(res , req)
	})
}