package controllers

import (
	"fmt"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/utils"
	"net/http"
)

type data struct{
	Head string
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
	utils.RenderTemplate(res , "main.page.gohtml" , data{"home"})
	fmt.Println(Repo.App)

}

func (r Repository) About(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , "about.page.gohtml" , data{"fuck"})
}