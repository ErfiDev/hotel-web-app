package dbrepo

import (
	"context"
	"database/sql"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/repository"
	"time"
)

type postgresDbRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(connection *sql.DB ,app *config.AppConfig) repository.DatabaseRepository {
	return &postgresDbRepo{
		App : app,
		DB: connection,
	}
}

func (psdb postgresDbRepo) InsertReservation(reservation models.Reservation) error {
	ctx , cancel := context.WithTimeout(context.Background() , 3 * time.Second)
	defer cancel()

	statement := `insert into reservations (first_name, last_name, email, phone, start_date, end_date,
	room_id, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)
`

	_ , err :=psdb.DB.ExecContext(
		ctx,
		statement,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomId,
		time.Now(),
		time.Now(), 
	)
	
	if err != nil {
		return err
	}
	
	return nil
}