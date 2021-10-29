package controllers

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/utils"
	"net/http"
)

var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App : app,
	}
}

func SetRepo(rep *Repository) {
	Repo = rep
}

func (r Repository) Home(res http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"erfan": "hanifezade",
		"title": "home page",
	}
	utils.RenderTemplate(res , "main.page.gohtml" , models.TmpData{
		Data: data,
	})
}

func (r Repository) About(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , "about.page.gohtml" , models.TmpData{
		Data: map[string]string{
			"about erfan": "erfanhanifezade",
			"title": "about page",
		},
	})
}

func (r Repository) Middleware(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , "about.page.gohtml" , nil)
}