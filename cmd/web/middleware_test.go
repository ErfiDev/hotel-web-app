package main

import (
	"fmt"
	"github.com/erfidev/hotel-web-app/routes"
	"net/http"
	"testing"
)

type MyHandler struct{}
func (m *MyHandler) ServeHTTP(w http.ResponseWriter , r *http.Request) {

}

func TestMiddleware(t *testing.T) {
	var myNewHandler MyHandler
	returnedFromNoSurf := routes.NoSurf(&myNewHandler)

	switch typ := returnedFromNoSurf.(type) {
	case http.Handler:
		// Do nothing

	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but %T" , typ))
	}
}