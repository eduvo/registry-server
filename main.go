package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/dancannon/gorethink"
	"log"
	"net/http"
	"os"
	"./actions"
	"./handlers"
	"./tools"
)

var (
	dbsession *gorethink.Session
	m 				*martini.Martini
	conf      tool.Conf
)

func init() {
	if err := tool.Flags(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err := conf.Config(tool.ConfigFile) // setups conf var
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	dbsession, err = gorethink.Connect(map[string]interface{}{
		"address":  conf.Dbaddress,
		"database": conf.Dbname,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func runTLS(m *martini.Martini) {
	log.Println("listening on port " + conf.Serverport)
	log.Fatalln(http.ListenAndServeTLS(":"+conf.Serverport, "config/cert.pem", "config/key.pem", m))
}

func main() {

	m := martini.New()
	m.Use(martini.Recovery())
	store := sessions.NewCookieStore([]byte(conf.Cookiesecret))
	m.Use(sessions.Sessions("registry", store))
	m.Use(render.Renderer())
	logger := log.New(os.Stdout, "", log.LstdFlags)
	m.Map(logger)

	r := martini.NewRouter()
	m.Handlers(
		handler.Xtralogger,
	)

	r.Get("/", func(session sessions.Session) string {
		if isLogin := action.IsLogin(session); isLogin {
			return "login"
		} else {
			return "no login"
		}
	})

	r.Get("/ping", func() string {
		return "OK"
	})

	r.Get("/who", action.Who)

	r.Get("/sign_in", action.SignIn)
	r.Post("/login", action.LogIn)

	m.Action(r.Handle)

	runTLS(m)

}
