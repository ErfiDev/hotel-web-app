package dbrepo

import (
	"database/sql"
	"time"

	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/repository"
)

type testDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewTestDBRepo(a *config.AppConfig) repository.DatabaseRepository {
	return &testDbRepo{
		App: a,
	}
}

func (psdb testDbRepo) InsertReservation(reservation models.Reservation) (int, error) {
	return 1, nil
}

func (psdb testDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) error {
	return nil
}

func (psdb testDbRepo) SearchAvailability(roomId int, start, end time.Time) (bool, error) {
	return false, nil
}

func (psdb testDbRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (psdb testDbRepo) FindRoomById(id int) (models.Room, error) {
	var room models.Room
	return room, nil
}

func (psdb testDbRepo) InsertUser(user models.User) (bool, error) {
	return true, nil
}

func (psdb testDbRepo) GetUserById(id int) (models.User, error) {
	return models.User{}, nil
}

func (psdb testDbRepo) Authenticate(email, password string) (int, bool, string) {
	return 0, true, ""
}

func (psdb testDbRepo) UpdateUser(user models.User) (bool, error) {
	return true, nil
}

func (psdb testDbRepo) AllReservations() ([]models.Reservation, error) {
	return []models.Reservation{}, nil
}

func (psdb testDbRepo) GetReservationById(id int) (models.Reservation, error) {
	return models.Reservation{}, nil
}

func (psdb testDbRepo) AllNewReservations() ([]models.Reservation, error) {
	return []models.Reservation{}, nil
}

func (psdb testDbRepo) UpdateReservation(res models.Reservation) (bool, error) {
	return true, nil
}

func (psdb testDbRepo) DeleteReservation(id int) (bool, error) {
	return true, nil
}

func (psdb testDbRepo) CompleteReservation(id int) (bool, error) {
	return true, nil
}
