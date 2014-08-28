package main

import (
	"encoding/json"
	"log"

	"github.com/go-martini/martini"
	_ "github.com/lib/pq"
)

type Goal struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Tallies int64  `json:"tallies"`
}

//
// HandleIndex returns an array of goals.
//
// GET /
//
func HandleIndex() []byte {
	goals := []Goal{}
	db.Find(&goals)
	goalsJson, err := json.Marshal(goals)

	if err != nil {
		log.Fatal(err)
	}

	return goalsJson
}

func main() {
	initDatabase()
	m := martini.Classic()

	m.Get("/", HandleIndex)
	m.Run()
}
