package main

import (
	"fmt"
	"github.com/erfidev/hotel-web-app/controllers"
	"net/http"
)

func main() {
	muxServer := http.NewServeMux()

	muxServer.HandleFunc("/" , controllers.Home)
	muxServer.HandleFunc("/about" , controllers.About)
	muxServer.Handle("/public/" , http.StripPrefix("/public/" , http.FileServer(http.Dir("./public"))))

	err := http.ListenAndServe(":3000" , muxServer)
	if err != nil {
		fmt.Println("we have the fucking error on starting server")
	}
}
