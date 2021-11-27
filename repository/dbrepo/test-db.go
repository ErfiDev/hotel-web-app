package dbrepo

import (
	"database/sql"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/repository"
	"time"
)

type testDbRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewTestDBRepo(a *config.AppConfig) repository.DatabaseRepository {
	return &testDbRepo{
		App: a,
	}
}

func (psdb testDbRepo) InsertReservation(reservation models.Reservation) (int , error) {
	return 1 , nil
}

func (psdb testDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) error {
	return nil
}

func (psdb testDbRepo) SearchAvailability(roomId int ,start , end time.Time) (bool , error) {
	return false , nil
}

func (psdb testDbRepo) SearchAvailabilityForAllRooms(start , end time.Time) ([]models.Room , error) {
	var rooms []models.Room
	return rooms , nil
}

func (psdb testDbRepo) FindRoomById(id int) (models.Room,error){
	var room models.Room
	return room , nil
}