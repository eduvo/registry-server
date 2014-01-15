package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	r "github.com/dancannon/gorethink"
	"log"
	"net/http"
	"os"
)

var (
	dbsession *r.Session
	conf      Conf
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
	dbsession, err = r.Connect(map[string]interface{}{
		"address":  conf.dbaddress,
		"database": conf.dbname,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func runTLS(m *martini.ClassicMartini) {
	log.Println("listening on port " + conf.serverport)
	log.Fatalln(http.ListenAndServeTLS(":"+conf.serverport, "cert.pem", "key.pem", m))
}

func main() {

	m := martini.Classic()
	store := sessions.NewCookieStore([]byte(conf.cookiesecret))
	m.Use(sessions.Sessions("registry", store))
	m.Use(render.Renderer())
	logger := log.New(os.Stdout, "", log.LstdFlags)
	m.Map(logger)

	m.Handlers(
		Xtralogger,
	)

	m.Get("/", func(session sessions.Session) string {
		if isLogin := IsLogin(session); isLogin {
			return "login"
		} else {
			return "no login"
		}
	})

	m.Get("/ping", func() string {
		return "OK"
	})

	m.Get("/who", Who)

	m.Get("/sign_in", SignIn)
	m.Post("/login", LogIn)

	runTLS(m)

}
