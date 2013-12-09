package main

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	//"github.com/davecgh/go-spew/spew"
	"net/http"
)

func SignIn(r render.Render) {
	r.HTML(200, "login", nil)
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
