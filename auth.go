package main

import (
	"github.com/codegangsta/martini-contrib/sessions"
)

func Auth(login string, password string) (bool, error) {
	if login == "11" && password == "22" {
		return true, nil
	}

	return false, nil
}

// TODO: I want a way to access session from anywhere, so I don't have to pass session in here.
func IsLogin(session sessions.Session) bool {
	isLogin := session.Get("is_login")
	if isLogin != nil {
		return true
	}
	return false
}
