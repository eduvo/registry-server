package main

import (
	"net/http"
	"time"
)

type Application struct {
	Id          string    `gorethink:"id,omitempty"`
	Name        string    `gorethink:"name"`
	Version     string    `gorethink:"version"`
	Environment string    `gorethink:"environment"`
	Public_key  string    `gorethink:"publickey"`
	Secret_key  string    `gorethink:"secretkey"`
	Backchannel string    `gorethink:"backchannel"`
	Created     time.Time `gorethink:"created,omitempty"`
}

type Domain struct {
	Id          string    `gorethink:"id,omitempty"`
	Name        string    `gorethink:"name"`
	Identifier  string    `gorethink:"identifier"`
	Stylesheet  string    `gorethink:"stylesheet"`
	Created     time.Time `gorethink:"created,omitempty"`
	Application Application
}

type Account struct {
	Id       string `gorethink:"id,omitempty"`
	Identity Identity
	Domain   Domain
	Created  time.Time
}

type Identity struct {
	Id       string `gorethink:"id,omitempty"`
	Email    string
	Password string
	Created  time.Time
}

type Session struct {
	Id      string `gorethink:"id,omitempty"`
	Account Account
	Created time.Time
}

func Who(res http.ResponseWriter, req *http.Request) string {

	// fake data
	application := Application{"1", "testapp", "0.0.1", "development", "12345678", "12345678901234567890123456789012", "", time.Now()}
	domain := Domain{"1", "test", "test", "test.css", time.Now(), application}
	identity := Identity{"1", "toto@toto.com", "xxx", time.Now()}
	account := Account{"1", identity, domain, time.Now()}
	accounts := map[string]Account{"toto@toto.com": account}

	email := Decrypt(application.Secret_key, req.FormValue("p"))
	return accounts[email].Id
}
