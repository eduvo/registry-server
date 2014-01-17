package data

import (
  "time"
  "fmt"
)

type DomainDB interface {
  //Get(id int) *Album
  //GetAll() []*Album
  //Find(band, title string, year int) []*Album
  Save(a *Domain)
  //Update(a *Album) error
  //Delete(id int)
}

type Domains struct {
  m   map[string]*Domain
}

func (list *Domains) Save(a *Domain) {
  list.m[a.Identifier] = a
}

type Domain struct {
  Id          string    `gorethink:"id,omitempty"`
  Name        string    `gorethink:"name"`
  Identifier  string    `gorethink:"identifier"`
  Stylesheet  string    `gorethink:"stylesheet"`
  Created     time.Time `gorethink:"created,omitempty"`
  Application Application
}

func (a *Domain) String() string {
  return fmt.Sprintf("%s - %s (%s)", a.Name, a.Identifier, a.Id)
}
