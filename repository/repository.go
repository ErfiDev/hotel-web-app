package repository

import "github.com/erfidev/hotel-web-app/models"

type DatabaseRepository interface {
	InsertReservation(reservation models.Reservation) (int , error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
}
