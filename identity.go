package main

import (
	"net/http"
	"strconv"
)

type Application struct {
	Id          int
	Name        string
	Version     string
	Environment string
	Public_key  string
	Secret_key  string
	Backchannel string
}

type Domain struct {
	Id          int
	Name        string
	Identifier  string
	Application Application
	Stylesheet  string
}

type Account struct {
	Id       int
	Identity Identity
	Domain   Domain
}

type Identity struct {
	Id       int
	Email    string
	Password string
}

type Session struct {
	Account Account
}

func Who(res http.ResponseWriter, req *http.Request) string {

	// fake data
	application := Application{1, "testapp", "0.0.1", "development", "12345678", "12345678901234567890123456789012", ""}
	domain := Domain{1, "test", "test", application, "test.css"}
	identity := Identity{1, "toto@toto.com", "xxx"}
	account := Account{1, identity, domain}
	accounts := map[string]Account{"toto@toto.com": account}

	email := Decrypt(application.Secret_key, req.FormValue("p"))
	return strconv.Itoa(accounts[email].Id)
}
