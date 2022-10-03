package models

import (
	"errors"

	_ "github.com/lib/pq"
	// "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	// _ "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	fmt.Println("Connecting to DB...")
	err := Connect()
	if err != nil {
		fmt.Print("Error connecting DB:", err)
		return
	}
	fmt.Println("Connected to DB")

	err = CreateTables()
	if err != nil {
		fmt.Println(err)
	}
}

func Connect() error {
	host := "127.0.0.1"
	port := "5433"
	user := "postgres"
	password := "admin"
	dbname := "my-posting-site"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	result, err := sqlx.Connect("postgres", dsn)
	DB = result

	if err != nil {
		return errors.New("connect DB failed, err:" + err.Error())
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	return nil
}
