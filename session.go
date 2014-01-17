package main

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	//"github.com/davecgh/go-spew/spew"
	"net/http"
	"time"
	"./data"
)

func SignIn(r render.Render) {
	application := data.Application{"1", "testapp", "0.0.1", "development", "12345678", "12345678901234567890123456789012", "", time.Now()}
	domain := data.Domain{"1", "test", "test", "test.css", time.Now(), application}
	r.HTML(200, "login", domain)
}

func LogIn(res http.ResponseWriter, req *http.Request, session sessions.Session) {
	/*
	   [SPEC]
	   redirect to?
	   (a) referral
	   (b) pass by argument

	   (b) > (a) ?
	*/
	login, password := req.FormValue("email"), req.FormValue("password")
	success, _ := Auth(login, password)

	if success {
		// poke Application back-channel URL here
		session.Set("is_login", true)
		http.Redirect(res, req, "http://www.google.com", http.StatusFound)
	} else {
		// redirect to sign in page.
	}

}
