package routes

import (
	"fmt"
	"net/http"
)

func YourMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter , req *http.Request) {
		fmt.Println("middleware working")
		handler.ServeHTTP(res , req)
	})
}
