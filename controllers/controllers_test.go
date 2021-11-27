package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type PostData struct {
	key , value string
}

var routeTests = []struct {
	name , method , url string
	params []PostData
	expectedStatusCode int
}{
	{"home" , "GET" , "/" , []PostData{} , http.StatusOK},
	{"about" , "GET" , "/about" , []PostData{} , http.StatusOK},
	{"contact" , "GET" , "/contact" , []PostData{} , http.StatusOK},
	{"book-now" , "GET" , "/book-now" , []PostData{} , http.StatusOK},
	{"rooms" , "GET" , "/rooms" , []PostData{} , http.StatusOK},
	{"rooms-generals" , "GET" , "/rooms/generals" , []PostData{} , http.StatusOK},
	{"rooms-majors" , "GET" , "/rooms/majors" , []PostData{} , http.StatusOK},
}

func TestControllers(t *testing.T) {
	routes := GetRoutes()
	ts := httptest.NewTLSServer(routes)

	for _ , route := range routeTests {
		if route.method == "GET" {
			res , err := ts.Client().Get(ts.URL + route.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if res.StatusCode != route.expectedStatusCode {
				t.Errorf("for %s , we need to %d but got %d" , route.name , route.expectedStatusCode , res.StatusCode)
			}
		} else {
			values := url.Values{}
			for _ , value := range route.params {
				values.Add(value.key , value.value)
			}
			res , err := ts.Client().PostForm(ts.URL + route.url , values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if res.StatusCode != route.expectedStatusCode {
				t.Errorf("for %s , we need to %d but got %d" , route.name , route.expectedStatusCode , res.StatusCode)
			}
		}
	}
}
