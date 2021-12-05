package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/erfidev/hotel-web-app/repository"
	"golang.org/x/crypto/bcrypt"
)

type postgresDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(connection *sql.DB, app *config.AppConfig) repository.DatabaseRepository {
	return &postgresDbRepo{
		App: app,
		DB:  connection,
	}
}

func (psdb postgresDbRepo) InsertReservation(reservation models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `insert into reservations (first_name, last_name, email, phone, start_date, end_date,
	room_id, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id
`
	var newId int

	err := psdb.DB.QueryRowContext(
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
		return 0, err
	}

	return newId, nil
}

func (psdb postgresDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `insert into room_restrictions (start_date,end_date,
				room_id, reservation_id , restriction_id , created_at , updated_at)
				values ($1 , $2 , $3, $4, $5, $6, $7)`

	_, err := psdb.DB.ExecContext(
		ctx,
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

func (psdb postgresDbRepo) SearchAvailability(roomId int, start, end time.Time) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	statement := `select count(id) from room_restrictions
	where room_id = $1 and $2 < end_date and $3 > start_date`

	row := psdb.DB.QueryRowContext(
		ctx,
		statement,
		roomId,
		start,
		end,
	)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

func (psdb postgresDbRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select id , room_name from rooms
	where id not in (select room_id from room_restrictions where $1 < end_date and $2 > start_date)`

	var rooms []models.Room

	rows, err := psdb.DB.QueryContext(ctx, statement, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		room := models.Room{}
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)

		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (psdb postgresDbRepo) FindRoomById(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select id, room_name, updated_at, created_at from rooms where id = $1`

	var room models.Room

	row := psdb.DB.QueryRowContext(ctx, statement, id)
	if err := row.Err(); err != nil {
		return room, err
	}

	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.UpdatedAt,
		&room.CreatedAt,
	)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (psdb postgresDbRepo) InsertUser(user models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	insert into users (first_name , last_name , email , password , access_level) 
	values ($1 , $2 , $3 , $4 , $5);
	`

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return false, err
	}

	_, err = psdb.DB.ExecContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		hashPass,
		user.AccessLevel,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (psdb postgresDbRepo) GetUserById(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select id ,first_name,last_name,email,password,access_level from users
		where id = $1
	`

	var user models.User

	row := psdb.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.AccessLevel,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (psdb postgresDbRepo) Authenticate(email, password string) (int, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select id, password from users
		where email = $1
	`

	var pass string
	var id int

	row := psdb.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(&id, &pass)
	if err != nil {
		return 0, false, "can't find user with this email!"
	}

	// check passwords
	hashError := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	if hashError == bcrypt.ErrMismatchedHashAndPassword {
		return 0, false, "mismatched hash and password"
	} else if hashError != nil {
		return 0, false, "unexpected error"
	}

	return id, true, ""
}

func (psdb postgresDbRepo) UpdateUser(user models.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		update users set first_name = $1, last_name = $2, email = $3, password = $4 , access_level = $5
	`

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return false, err
	}

	_, errOnExec := psdb.DB.ExecContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		hashPass,
		user.AccessLevel,
	)

	if errOnExec != nil {
		return false, errOnExec
	}

	return true, nil
}

func (psdb postgresDbRepo) AllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
  	select r.id, r.start_date, r.end_date , r.email ,
	r.first_name , r.last_name , r.phone , r.created_at,
	r.updated_at , rm.room_name , rm.id
	from reservations r
	left join rooms rm on (r.room_id = rm.id) 
	order by r.start_date asc;
	`

	var reservations []models.Reservation

	rows, err := psdb.DB.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return reservations, err
	}

	for rows.Next() {
		res := models.Reservation{}
		err := rows.Scan(
			&res.ID,
			&res.StartDate,
			&res.EndDate,
			&res.Email,
			&res.FirstName,
			&res.LastName,
			&res.Phone,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.Room.RoomName,
			&res.RoomId,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, res)
	}

	return reservations, nil
}
