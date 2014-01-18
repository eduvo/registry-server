package action

import (
  "net/http"
  "time"
  "../data"
  "../tools"
)

func Who(res http.ResponseWriter, req *http.Request) string {

  // fake data
  application := data.Application{"1", "testapp", "0.0.1", "development", "12345678", "12345678901234567890123456789012", "", time.Now()}
  domain := data.Domain{"1", "test", "test", "test.css", time.Now(), application}
  identity := data.Identity{"1", "toto@toto.com", "xxx", time.Now()}
  account := data.Account{"1", identity, domain, time.Now()}
  accounts := map[string]data.Account{"toto@toto.com": account}

  email := tool.Decrypt(application.Secret_key, req.FormValue("p"))
  return accounts[email].Id
}
