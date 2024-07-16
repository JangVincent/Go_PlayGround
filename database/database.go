package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	UniqueViolationErr = pq.ErrorCode("23505")
)

func GetDB() (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
		"localhost", 
		"5432", 
		"root", 
		"test", 
		"root",
	)

	db, err := sqlx.Connect("postgres", dataSource)

	if err != nil {
		log.Printf("fail to get database connection pool :: %v\n", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("fail to ping database :: %v\n", err)
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("fail to close database connection pool :: %v\n", err)
	}
}
