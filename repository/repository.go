package repository

import (
	"time"

	"github.com/erfidev/hotel-web-app/models"
)

type DatabaseRepository interface {
	InsertReservation(reservation models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
	SearchAvailability(roomId int, start, end time.Time) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	FindRoomById(id int) (models.Room, error)
	InsertUser(user models.User) (bool, error)
	GetUserById(id int) (models.User, error)
	Authenticate(email, password string) (int, bool, string)
	UpdateUser(user models.User) (bool, error)
	AllReservations() ([]models.Reservation, error)
	GetReservationById(id int) (models.Reservation, error)
	AllNewReservations() ([]models.Reservation, error)
	UpdateReservation(models.Reservation) (bool, error)
}
