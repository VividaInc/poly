package oauth

import (
  "net/http"
  "polypentec/cookie"
  "polypentec/protocol/db"
  "polypentec/store"
)

func SignedInUser(r *http.Request) bool {
  ok := false
  _, err := store.GetUser(r)
  if err != nil {
    return ok
  }
  ok = true
  return ok
}

func SignInUser(w http.ResponseWriter, r *http.Request) bool {
  ok := false
  r.ParseForm()
  username := r.FormValue("username")
  password := r.FormValue("password")
  user := db.GetUserFromDatabase(username)
  if username != user.Username || password != user.Password {
    return ok
  }
  cookie.SaveCookie(w, "ZmemAuthorization", username)
  ok = true
  return ok
}
