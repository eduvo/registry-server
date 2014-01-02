package main

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

func Who(user string) bool {
  return true
}
