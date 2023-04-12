package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal("Error")
	}

	sqlBytes, err := os.ReadFile("models/init.sql")
	if err != nil {
		log.Fatal(err)
	}
	sql := string(sqlBytes)

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()
}

func CloseDB() {
	db.Close()
}
