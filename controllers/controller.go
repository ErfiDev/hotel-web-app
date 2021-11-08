package controllers

import (
	"fmt"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/forms"
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
	data := make(map[string]interface{})
	data["path"] = "/"
	data["title"] = "Home Page"
	remoteIp := req.RemoteAddr
	r.App.Session.Put(req.Context() , "remote_ip" , remoteIp)


	utils.RenderTemplate(res , req , "landing.page.gohtml" , &models.TmpData{
		Data: data,
	})
}

func (r Repository) About(res http.ResponseWriter, req *http.Request) {
	utils.RenderTemplate(res , req , "about.page.gohtml" , &models.TmpData{
		Data: map[string]interface{}{
			"path": "/about",
			"title": "about page",
		},
	})
}


func (r Repository) Rooms(res http.ResponseWriter , req *http.Request) {
	pageData := models.TmpData{
		Data: map[string]interface{}{
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
		Data : map[string]interface{}{
			"title": "Book now",
			"path": "/book-now",
		},
		Form: forms.New(nil),
	}

	utils.RenderTemplate(res , req , "book.page.gohtml" , &pageData)
}

func (r Repository) Contact(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , req , "contact.page.gohtml" , &models.TmpData{
		Data: map[string]interface{}{
			"title": "contact page",
			"path": "/contact",
		},
	})
}

func (r Repository) MakeReservation(res http.ResponseWriter , req *http.Request) {
	utils.RenderTemplate(res , req , "make-reservation.page.gohtml" , &models.TmpData{
		Data: map[string]interface{}{
			"title": "make your reservation page",
			"path": "/book-now",
		},
		Form: forms.New(nil),
	})
}

func (r Repository) BookNowPost(res http.ResponseWriter , req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	newBookNow := models.BookNow{
		Start: req.Form.Get("start-date"),
		End: req.Form.Get("ending-date"),
	}

	form := forms.New(req.PostForm)

	form.Has("start-date" , req)
	form.Has("ending-date" , req)

	if !form.Valid() {
		data := models.TmpData{
			Data: map[string]interface{} {
				"reservation": newBookNow,
				"path": "/book-now",
				"title": "Book now",
			},
			Form: form,
		}

		utils.RenderTemplate(res , req , "book.page.gohtml" , &data)
		return
	} else {
		res.Write([]byte("book now is complete"))
	}
}

func (r Repository) MakeReservationPost(res http.ResponseWriter , req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	reservationData := models.Reservation{
		Firstname: req.Form.Get("first_name"),
		Lastname: req.Form.Get("last_name"),
		Email: req.Form.Get("email"),
		Phone: req.Form.Get("phone"),
	}

	form := forms.New(req.PostForm)

	form.Has("first_name" , req)
	form.Has("last_name" , req)
	form.Has("email" , req)
	form.Has("phone" , req)

	if !form.Valid(){
		data := make(map[string]interface{})
		data["reservation"] = reservationData
		data["title"] = "Make reservation"
		data["path"] = "/book-now"

		utils.RenderTemplate(res , req , "make-reservation.page.gohtml" , &models.TmpData{
			Data: data,
			Form: form,
		})

		return
	}

	r.App.Session.Put(req.Context() , "reservation" , reservationData)

	http.Redirect(res , req , "/reservation-summary" , http.StatusSeeOther)
}

func (r Repository) ReservationSummary(res http.ResponseWriter , req *http.Request) {
	reservation , isOk := r.App.Session.Get(req.Context() , "reservation").(models.Reservation)
	if !isOk {
		r.App.Session.Put(req.Context() , "error" , "can't get reservation data from session")
		http.Redirect(res , req , "/" , http.StatusTemporaryRedirect)
		return
	}

	r.App.Session.Remove(req.Context() , "reservation")

	utils.RenderTemplate(res , req , "reservation.page.gohtml" , &models.TmpData{
		Data: map[string]interface{}{
			"reservation": reservation,
			"title": "Reservation summary",
			"path": "/book-now",
		},
	})
}