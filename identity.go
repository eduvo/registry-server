package main

import (
	"net/http"
	"strconv"
)

type Application struct {
	id          int
	name        string
	version     string
	environment string
	public_key  string
	secret_key  string
	backchannel string
}

type Domain struct {
	id          int
	name        string
	identifier  string
	application Application
	stylesheet  string
}

type Account struct {
	id       int
	identity Identity
	domain   Domain
}

type Identity struct {
	id       int
	email    string
	password string
}

type Session struct {
	account Account
}

func Who(res http.ResponseWriter, req *http.Request) string {

	// fake data
	application := Application{1, "testapp", "0.0.1", "development", "12345678", "12345678901234567890123456789012", ""}
	domain := Domain{1, "test", "test", application, "test.css"}
	identity := Identity{1, "toto@toto.com", "xxx"}
	account := Account{1, identity, domain}
	accounts := map[string]Account{"toto@toto.com": account}

	email := Decrypt(application.secret_key, req.FormValue("p"))
	return strconv.Itoa(accounts[email].id)
}
