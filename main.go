package main

import (
	"./db"
	"./route"

	"github.com/go-martini/martini"
)

func main() {
	db.Init()
	m := martini.Classic()

	m.Get("/", route.GoalsIndex)
	m.Run()
}
