package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxDBIdleConn = 5
const maxDBLifeTime = 5 * time.Minute

func ConnectDB(distro string) (*DB , error) {
	db , err := NewDataBase(distro)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxDBIdleConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	dbConn.SQL = db

	return dbConn , nil
}

func NewDataBase(constr string) (*sql.DB , error) {
	db , err := sql.Open("pgx" , constr)
	if err != nil {
		return nil , err
	}

	if err = db.Ping(); err != nil {
		return nil , err
	}

	return db , nil
}