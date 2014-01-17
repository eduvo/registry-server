package data

import (
  "time"
  "fmt"
)

type SessionDB interface {
  //Get(id int) *Album
  //GetAll() []*Album
  //Find(band, title string, year int) []*Album
  Save(a *Session)
  //Update(a *Album) error
  //Delete(id int)
}

type Sessions struct {
  m   map[string]*Session
}

func (list *Sessions) Save(a *Session) {
  list.m[a.Id] = a
}

type Session struct {
  Id      string `gorethink:"id,omitempty"`
  Account Account
  Created time.Time
}

func (a *Session) String() string {
  return fmt.Sprintf("%s", a.Id)
}
