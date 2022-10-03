package models

import (
	_ "github.com/lib/pq"
	// "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	// _ "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	fmt.Println("Init")
	host := "127.0.0.1"
	port := "5433"
	user := "postgres"
	password := "admin"
	dbname := "my-posting-site"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	fmt.Println("Connecting to DB...")
	DB, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		fmt.Printf("Connect DB failed, err:%v\n", err)
		return
	}
	fmt.Println("Connected to DB")

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
}
