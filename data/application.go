package data

import (
  "time"
  "fmt"
)

type ApplicationDB interface {
  //Get(id int) *Album
  //GetAll() []*Album
  //Find(band, title string, year int) []*Album
  Save(a *Application)
  //Update(a *Album) error
  //Delete(id int)
}

type Applications struct {
  m   map[string]*Application
}

func (list *Applications) Save(a *Application) {
  list.m[a.Public_key] = a
}

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

func (a *Application) String() string {
  return fmt.Sprintf("%s %s (%s)", a.Name, a.Version, a.Id)
}
