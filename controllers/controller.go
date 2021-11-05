package controllers

import (
	"encoding/json"
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


	utils.RenderTemplate(res , req , "landing.page.gohtml" , &models.TmpData{
		Data: data,
	})
}

func (r Repository) About(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , req , "about.page.gohtml" , &models.TmpData{
		Data: map[string]string{
			"path": "/about",
			"title": "about page",
		},
	})
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
		utils.RenderTemplate(res , req , "rooms.page.gohtml" , &pageData)

	case "/rooms/generals":
		pageData.Data["title"] = "Generals suite"
		utils.RenderTemplate(res , req , "generals.page.gohtml" , &pageData)

	case "/rooms/majors":
		pageData.Data["title"] = "Majors suite"
		utils.RenderTemplate(res , req , "majors.page.gohtml" , &pageData)

	default:
		res.Write([]byte("page not found"))
	}
}

func (r Repository) BookNow(res http.ResponseWriter , req *http.Request) {
	pageData := models.TmpData{
		Data : map[string]string{
			"title": "Book now",
			"path": "/book-now",
		},
	}

	utils.RenderTemplate(res , req , "book.page.gohtml" , &pageData)
}

func (r Repository) Contact(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , req , "contact.page.gohtml" , &models.TmpData{
		Data: map[string]string{
			"title": "contact page",
			"path": "/contact",
		},
	})
}

func (r Repository) MakeReservation(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , req , "make-reservation.page.gohtml" , &models.TmpData{
		Data: map[string]string{
			"title": "make your reservation page",
			"path": "/book-now",
		},
	})
}

func (r Repository) BookNowPost(res http.ResponseWriter , req *http.Request) {
	toJson , _ := json.Marshal(map[string]string{
		"msg": "Search for availability",
		"status": "200",
	})

	res.Write([]byte(toJson))
}