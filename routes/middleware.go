package routes

import (
	"fmt"
	"net/http"
)

func YourMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func (res http.ResponseWriter , req *http.Request) {
		fmt.Println(req.URL)
		fmt.Println(req.Cookies())
		fmt.Println(req.Method)

		handler.ServeHTTP(res , req)
	}
}
