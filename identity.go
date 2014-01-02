package main

import (
  "github.com/codegangsta/martini-contrib/sessions"
  "net/http"
)

type Application struct {
  id int
  name string
  version string
  environment string
  public_key string
  secret_key string
  backchannel string
}

type Domain struct {
  id int
  name string
  identifier string
  application Application
  stylesheet string
}

type Account struct {
  id int
  identity Identity
  domain Domain
}

type Identity struct {
  id int
  email string
  password string
  accounts []Account
}

type Session struct {
  account Account
}

func Who(res http.ResponseWriter, req *http.Request, session sessions.Session) string {
  application := Application{1, "testapp", "0.0.1", "development", "12345678", "12345678901234567890123456789012", ""}
  return Decrypt([]byte(application.secret_key), []byte(req.FormValue("p")))
}
