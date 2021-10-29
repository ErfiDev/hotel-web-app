package routes

import (
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
