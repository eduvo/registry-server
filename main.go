package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/sign_in", SignIn)
	m.Post("/login", LogIn)

	m.Run()
}
