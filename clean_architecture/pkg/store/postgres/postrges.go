package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Settings struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Connection struct {
	Pool *sql.DB
}

func New(settings Settings) (*Connection, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		settings.Host, settings.Port, settings.User, settings.Password, settings.DBName)

	pool, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = pool.Ping()
	if err != nil {
		return nil, err
	}

	return &Connection{Pool: pool}, nil
}
