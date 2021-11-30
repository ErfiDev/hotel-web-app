package repository

import (
	"github.com/erfidev/hotel-web-app/models"
	"time"
)

type DatabaseRepository interface {
	InsertReservation(reservation models.Reservation) (int , error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
	SearchAvailability(roomId int, start , end time.Time) (bool , error)
	SearchAvailabilityForAllRooms(start , end time.Time) ([]models.Room , error)
	FindRoomById(id int) (models.Room,error)
	InsertUser(user models.User) (bool , error)
	GetUserById(id int) (models.User , error)
	Authenticate(email , password string) (bool , string)
	UpdateUser(user models.User) (bool , error)
}
