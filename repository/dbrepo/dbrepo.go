package dbrepo

import (
	"database/sql"
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/repository"
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
