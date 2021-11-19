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

func (psdb postgresDbRepo) InsertReservation(reservation models.Reservation) (int , error) {
	ctx , cancel := context.WithTimeout(context.Background() , 3 * time.Second)
	defer cancel()

	statement := `insert into reservations (first_name, last_name, email, phone, start_date, end_date,
	room_id, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id
`
	var newId int

	err :=psdb.DB.QueryRowContext(
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
	).Scan(&newId)
	
	if err != nil {
		return 0 , err
	}
	
	return newId , nil
}

func (psdb postgresDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) error {
	ctx , cancel := context.WithTimeout(context.Background() , 3 * time.Second)
	defer cancel()

	statement := `insert into room_restrictions (start_date,end_date,
				room_id, reservation_id , restriction_id , created_at , updated_at)
				values ($1 , $2 , $3, $4, $5, $6, $7)`

	_ , err := psdb.DB.ExecContext(
		ctx ,
		statement,
		roomRestriction.StartDate,
		roomRestriction.EndDate,
		roomRestriction.RoomId,
		roomRestriction.ReservationId,
		roomRestriction.RestrictionId,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (psdb postgresDbRepo) SearchAvailability(start , end time.Time) (bool , error) {
	ctx , cancel := context.WithTimeout(context.Background() , 3 * time.Second)
	defer cancel()

	var numRows int

	statement := `select count(id) from room_restrictions
	where $1 < end_date and $2 > start_date`

	row := psdb.DB.QueryRowContext(
		ctx,
		statement,
		start,
		end,
	)
	err := row.Scan(&numRows)
	if err != nil {
		return false , err
	}

	if numRows == 0 {
		return true , nil
	}

	return false , nil
}