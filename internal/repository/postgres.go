package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *Repository) InitDb() error {
	query1 := `CREATE TABLE IF NOT EXISTS cams
	(
		id SERIAL PRIMARY KEY,
		camname CHARACTER VARYING(30),
		lon double precision,
		lat double precision,
		addr CHARACTER VARYING(50) UNIQUE
	);`
	_, err := r.db.Exec(query1)
	if err != nil {
		return err
	}
	query2 := `CREATE TABLE IF NOT EXISTS tasks
	(
		id SERIAL PRIMARY KEY,
		query CHARACTER VARYING(30),
		status CHARACTER VARYING(60),
		count smallint
	);`
	_, err = r.db.Exec(query2)
	if err != nil {
		return err
	}
	return nil
}
