package store

import (
  "net/http"
  e "poly/err"
  "poly/protocol/db"
)

func GetUser(r *http.Request) (*db.User, *e.RequestError) {
  cookie, err := r.Cookie("ZmemAuthorization")
  if err != nil {
    return nil, e.RETURNINTERNALERROR
  }
  user := db.GetUserFromDatabase(cookie.Value)
  if err := validateAuth(user, cookie.Value); err != nil {
    return nil, err
  }
  return user, nil
}

func validateAuth(user *db.User, current string) *e.RequestError {
  if user.Username != current {
    return e.RETURNINTERNALERROR
  }
  return nil
}

