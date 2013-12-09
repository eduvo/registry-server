package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(render.Renderer())

	m.Get("/", func(session sessions.Session) string {
		isLogin := IsLogin(session)
		if isLogin {
			return "login"
		} else {
			return "no login"
		}
	})

	m.Get("/sign_in", SignIn)
	m.Post("/login", LogIn)

	m.Run()
}
