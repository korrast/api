package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConn() (*gorm.DB, error) {
	var host, port, user, password, dbname = os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("DSN :", dsn)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Can't open database")
		return nil, err
	}

	return conn, nil
}

func CloseConn(conn *gorm.DB) error {
	db, err := conn.DB()
	if err != nil {
		fmt.Println("Couldn't get DB from conn")
		return err
	}

	db.Close()
	return nil
}
