package main

import (
	"net/http"
	"time"
  "./data"
)

type Account struct {
	Id       string `gorethink:"id,omitempty"`
	Identity data.Identity
	Domain   data.Domain
	Created  time.Time
}

type Session struct {
	Id      string `gorethink:"id,omitempty"`
	Account Account
	Created time.Time
}

func Who(res http.ResponseWriter, req *http.Request) string {

	// fake data
	application := data.Application{"1", "testapp", "0.0.1", "development", "12345678", "12345678901234567890123456789012", "", time.Now()}
	domain := data.Domain{"1", "test", "test", "test.css", time.Now(), application}
	identity := data.Identity{"1", "toto@toto.com", "xxx", time.Now()}
	account := Account{"1", identity, domain, time.Now()}
	accounts := map[string]Account{"toto@toto.com": account}

	email := Decrypt(application.Secret_key, req.FormValue("p"))
	return accounts[email].Id
}
