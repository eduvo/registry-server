package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/dancannon/gorethink"
	"log"
	"net/http"
	"os"
)

var (
	dbsession *gorethink.Session
	conf      Conf
	m 				*martini.Martini
)

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
	dbsession, err = gorethink.Connect(map[string]interface{}{
		"address":  conf.dbaddress,
		"database": conf.dbname,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func runTLS(m *martini.Martini) {
	log.Println("listening on port " + conf.serverport)
	log.Fatalln(http.ListenAndServeTLS(":"+conf.serverport, "cert.pem", "key.pem", m))
}

func main() {

	m := martini.New()
	m.Use(martini.Recovery())
	store := sessions.NewCookieStore([]byte(conf.cookiesecret))
	m.Use(sessions.Sessions("registry", store))
	m.Use(render.Renderer())
	logger := log.New(os.Stdout, "", log.LstdFlags)
	m.Map(logger)

	r := martini.NewRouter()
	m.Handlers(
		Xtralogger,
	)

	r.Get("/", func(session sessions.Session) string {
		if isLogin := IsLogin(session); isLogin {
			return "login"
		} else {
			return "no login"
		}
	})

	r.Get("/ping", func() string {
		return "OK"
	})

	r.Get("/who", Who)

	r.Get("/sign_in", SignIn)
	r.Post("/login", LogIn)

	runTLS(m)

}
