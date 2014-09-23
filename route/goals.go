package route

import (
	"encoding/json"
	"log"

	"../model"
	"../db"
)

//
// GoalsIndex returns an array of goals as JSON.
//
// GET /
//
func GoalsIndex() []byte {
	goals := []model.Goal{}
	db.Db.Find(&goals)
	goalsJSON, err := json.Marshal(goals)

	if err != nil {
		log.Fatal(err)
	}

	return goalsJSON
}
