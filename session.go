package main

import (
	"github.com/codegangsta/martini-contrib/render"
	//"github.com/davecgh/go-spew/spew"
	"net/http"
)

func SignIn(r render.Render) {
	r.HTML(200, "login", nil)
}

func LogIn(res http.ResponseWriter, req *http.Request) {
	success, _ := Auth(req.FormValue("email"), req.FormValue("password"))
	if success {
		http.Redirect(res, req, "http://www.google.com", http.StatusMovedPermanently)
	} else {

	}

}
