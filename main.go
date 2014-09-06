package main

import (
	"./model"
	"./route"

	"github.com/go-martini/martini"
)

func main() {
	model.InitDatabase()
	m := martini.Classic()

	m.Get("/", route.GoalsIndex)
	m.Run()
}
