package main

import "github.com/go-martini/martini"

//
// HandleIndex says hello to the world.
//
// GET /
//
func HandleIndex() string {
	return "Hello World!"
}

func main() {
	m := martini.Classic()

	m.Get("/", HandleIndex())
	m.Run()
}
