package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db gorm.DB

//
// Initializes database connection and assigns globally available
// gorm.DB.
//
func initDatabase() {
	var err error
	db, err = gorm.Open("postgres", "dbname=tallyup sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	db.DB()
}
