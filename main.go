package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"log"
	"net/http"
	"os"
)

var conf Conf

func init() {
	if err := Flags(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err := Config() // setups conf var
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func runTLS(m *martini.ClassicMartini) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	log.Println("listening on port " + port)
	log.Fatalln(http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", m))
}

func main() {

	//println(conf.servername)
	m := martini.Classic()
	m.Use(martini.Logger())
	store := sessions.NewCookieStore([]byte(conf.cookiesecret))
	m.Use(sessions.Sessions("registry", store))
	m.Use(martini.Recovery())
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

	runTLS(m)

}
