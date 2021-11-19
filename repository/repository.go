package repository

import "github.com/erfidev/hotel-web-app/models"

type DatabaseRepository interface {
	InsertReservation(reservation models.Reservation) error
}
