package route

import (
	"encoding/json"
	"log"

	"../model"
)

//
// GoalsIndex returns an array of goals as JSON.
//
// GET /
//
func GoalsIndex() []byte {
	goals := []model.Goal{}
	model.Db.Find(&goals)
	goalsJSON, err := json.Marshal(goals)

	if err != nil {
		log.Fatal(err)
	}

	return goalsJSON
}
