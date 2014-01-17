package data

import (
  "time"
  "fmt"
)

type IdentityDB interface {
  //Get(id int) *Album
  //GetAll() []*Album
  //Find(band, title string, year int) []*Album
  Save(a *Identity)
  //Update(a *Album) error
  //Delete(id int)
}

type Identities struct {
  m   map[string]*Identity
}

func (list *Identities) Save(a *Identity) {
  list.m[a.Email] = a
}

type Identity struct {
  Id       string `gorethink:"id,omitempty"`
  Email    string
  Password string
  Created  time.Time
}

func (a *Identity) String() string {
  return fmt.Sprintf("%s (%s)", a.Email, a.Id)
}
