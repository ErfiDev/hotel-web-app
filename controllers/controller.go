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
		"path": "/",
		"title": "home page",
	}
	remoteIp := req.RemoteAddr
	r.App.Session.Put(req.Context() , "remote_ip" , remoteIp)


	utils.RenderTemplate(res , "landing.page.gohtml" , models.TmpData{
		Data: data,
	})
}

func (r Repository) About(res http.ResponseWriter, req *http.Request) {
	getRemoteIp := r.App.Session.GetString(req.Context() , "remote_ip")

	stringMap := map[string]string{
		"remote_ip": getRemoteIp,
	}

	utils.RenderTemplate(res , "about.page.gohtml" , models.TmpData{
		Data: map[string]string{
			"path": "/about",
			"title": "about page",
		},
		StringMap: stringMap,
	})
}

func (r Repository) Middleware(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , "about.page.gohtml" , nil)
}

func (r Repository) Rooms(res http.ResponseWriter , req *http.Request) {
	pageData := models.TmpData{
		Data: map[string]string{
			"title": "Rooms page",
			"path": "/rooms",
		},
	}

	switch req.RequestURI {
	case "/rooms":
		utils.RenderTemplate(res , "rooms.page.gohtml" , pageData)

	case "/rooms/generals":
		pageData.Data["title"] = "Generals suite"
		utils.RenderTemplate(res , "generals.page.gohtml" , pageData)

	case "/rooms/majors":
		pageData.Data["title"] = "Majors suite"
		utils.RenderTemplate(res , "majors.page.gohtml" , pageData)

	default:
		res.Write([]byte("page not found"))
	}
}