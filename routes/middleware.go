package routes

import (
	"github.com/justinas/nosurf"
	"net/http"
)

func UserCheckMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter , req *http.Request) {
		mtd := req.Method
		if mtd == "GET" {
			req.AddCookie(&http.Cookie{
				Name: "erfan",
				Value: "fuckedup",
			})
			handler.ServeHTTP(res , req)
		} else {
			handler.ServeHTTP(res , req)
		}
	})
}

func NoSurf(next http.Handler) http.Handler {
	CSRFHandler := nosurf.New(next)

	CSRFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
	})

	return CSRFHandler
}
