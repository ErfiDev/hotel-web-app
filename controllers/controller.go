package controllers

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/driver"
	"github.com/erfidev/hotel-web-app/forms"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/repository"
	"github.com/erfidev/hotel-web-app/repository/dbrepo"
	"github.com/erfidev/hotel-web-app/utils"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

var Repo *Repository

type Repository struct{
	App *config.AppConfig
	DB repository.DatabaseRepository
}

func NewRepository(app *config.AppConfig , db *driver.DB) *Repository {
	return &Repository{
		App : app,
		DB: dbrepo.NewPostgresRepo(db.SQL , app),
	}
}

func SetRepo(rep *Repository) {
	Repo = rep
}

func (r Repository) Home(res http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["path"] = "/"
	data["title"] = "Home Page"

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
	reservation , ok := r.App.Session.Get(req.Context() , "reservation").(models.Reservation)
	if !ok {
		http.Redirect(res , req , "/book-now" , http.StatusTemporaryRedirect)
		return
	}

	st := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")

	room , err := r.DB.FindRoomById(reservation.RoomId)
	if err != nil {
		utils.ServerError(res , err)
		return
	}

	reservation.Room = room

	StringMap := map[string]string{
		"startDate": st,
		"endDate": ed,
	}
	Data := map[string]interface{}{
		"title": "make reservation",
		"path": "/book-now",
		"reservation": reservation,
	}

	utils.RenderTemplate(res , req , "make-reservation.page.gohtml" , &models.TmpData{
		Data: Data,
		Form: forms.New(nil),
		StringMap: StringMap,
	})
}

func (r Repository) BookNowPost(res http.ResponseWriter , req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		utils.ServerError(res , err)
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
		layout := "2006-01-02"
		startDate , err := time.Parse(layout,newBookNow.Start)
		if err != nil {
			utils.ServerError(res , err)
			return
		}
		endDate , err := time.Parse(layout,newBookNow.End)
		if err != nil {
			utils.ServerError(res , err)
			return
		}

		rooms , err := r.DB.SearchAvailabilityForAllRooms(startDate , endDate)
		if err != nil {
			utils.ServerError(res , err)
			return
		}

		if len(rooms) <= 0 {
			r.App.Session.Put(req.Context() , "error" , "not availability")
			http.Redirect(res , req , "/book-now" , http.StatusSeeOther)
			return
		}

		data := make(map[string]interface{})
		data["title"] = "Choose a room"
		data["path"] = "/book-now"
		data["rooms"] = rooms

		reservation := models.Reservation{
			StartDate: startDate,
			EndDate:   endDate,
		}

		r.App.Session.Put(req.Context() , "reservation" , reservation)

		utils.RenderTemplate(res , req , "choose-room.page.gohtml" , &models.TmpData{
			Data: data,
		})
	}
}

func (r Repository) ChooseRoom(res http.ResponseWriter , req *http.Request) {
	roomId , err := strconv.Atoi(chi.URLParam(req , "id"))
	if err != nil {
		utils.ServerError(res , err)
		return
	}

	reservation := r.App.Session.Get(req.Context() , "reservation").(models.Reservation)

	reservation.RoomId = roomId

	r.App.Session.Put(req.Context() , "reservation" , reservation)
	http.Redirect(res , req , "/make-reservation" , http.StatusTemporaryRedirect)
}

func (r Repository) MakeReservationPost(res http.ResponseWriter , req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		utils.ServerError(res , err)
		return
	}

	sd := req.Form.Get("start_date")
	ed := req.Form.Get("end_date")

	layout := "2006-01-02"
	stParse, err := time.Parse(layout , sd)
	if err != nil {
		utils.ServerError(res , err)
		return
	}
	edParse , err := time.Parse(layout , ed)
	if err != nil {
		utils.ServerError(res , err)
		return
	}

	roomId := req.Form.Get("room_id")
	roomIdToInt, err := strconv.Atoi(roomId)
	if err != nil {
		utils.ServerError(res , err)
		return
	}

	reservationData := models.Reservation{
		FirstName: req.Form.Get("first_name"),
		LastName: req.Form.Get("last_name"),
		Email: req.Form.Get("email"),
		Phone: req.Form.Get("phone"),
		StartDate: stParse,
		EndDate: edParse,
		RoomId: roomIdToInt,
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

	reservationId,errInsert := r.DB.InsertReservation(reservationData)
	if errInsert != nil {
		utils.ServerError(res , errInsert)
		return
	}

	roomRestrictions := models.RoomRestriction{
		StartDate:     stParse,
		EndDate:       edParse,
		ReservationId: reservationId,
		RestrictionId: 1,
		RoomId:        roomIdToInt,
	}

	err = r.DB.InsertRoomRestriction(roomRestrictions)
	if err != nil {
		utils.ServerError(res , err)
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