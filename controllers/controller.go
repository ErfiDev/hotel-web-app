package controllers

import (
	"github.com/erfidev/hotel-web-app/utils"
	"net/http"
)

type data struct{
	Head string
}

func Home(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , "main.page.gohtml" , data{"home"})
}

func About(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , "about.page.gohtml" , data{"fuck"})
}