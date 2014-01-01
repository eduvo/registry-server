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

func main() {

	//println(conf.servername)
	m := martini.Classic()
	if err := http.ListenAndServeTLS(conf.serverport, "cert.pem", "key.pem", m); err != nil {
		log.Println("Please generate SSL certificate: go run $GOROOT/src/pkg/crypto/tls/generate_cert.go --host='localhost'")
		os.Exit(0)
	}
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("registry", store))
	m.Use(render.Renderer())
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

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
