package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func StartConn() (*DB, err) {
	var host, port, user, password, dbname = os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")

	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		fmt.Println("Can't open database")
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Can't ping database")
		return nil, err
	}

	return db, nil
}

func CloseCon(db *DB) {
	db.Close()
}