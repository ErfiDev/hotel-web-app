package controllers

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/utils"
	"net/http"
)

type TemplateData struct{
	Title string
	Auth bool
	Username string
}

var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App : app,
	}
}

func SetRepo(rep *Repository){
	Repo = rep
}

func (r Repository) Home(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , "main.page.gohtml" , TemplateData{
		"Home page /",
		false,
		"",
	})
}

func (r Repository) About(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , "about.page.gohtml" , TemplateData{
		"About page /about",
		false,
		"",
	})
}

func (r Repository) Middleware(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , "about.page.gohtml" , nil)
}