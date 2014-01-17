package data

import (
  "time"
  "fmt"
)

type AccountDB interface {
  //Get(id int) *Album
  //GetAll() []*Album
  //Find(band, title string, year int) []*Album
  Save(a *Account)
  //Update(a *Album) error
  //Delete(id int)
}

type Accounts struct {
  m   map[string]*Account
}

func (list *Accounts) Save(a *Account) {
  list.m[a.Id] = a
}

type Account struct {
  Id       string `gorethink:"id,omitempty"`
  Identity Identity
  Domain   Domain
  Created  time.Time
}

func (a *Account) String() string {
  return fmt.Sprintf("%s - %s (%s)", a.Identity.Email, a.Domain.Identifier, a.Id)
}
