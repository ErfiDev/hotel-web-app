package main

import (
	"fmt"
	"github.com/erfidev/hotel-web-app/routes"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	newRoutes := routes.Routes()

	switch typ := newRoutes.(type) {
	case http.Handler:
		// Do nothing

	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but %T" , typ))
	}
}
