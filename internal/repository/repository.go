package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "test_db"
)

type Repository struct {
	db *sql.DB
}

func New() (*Repository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("cannot connect to db", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("error ping db", err)
		return nil, err
	}
	log.Println("connect to db success")
	return &Repository{
		db: db,
	}, nil
}
