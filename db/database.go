package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//
// Db is used to interact with the database.
//
var Db gorm.DB

//
// Init initializes database connection and assigns globally available
// gorm.DB.
//
func Init() {
	var err error
	Db, err = gorm.Open("postgres", "dbname=tallyup sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	Db.DB()
}
