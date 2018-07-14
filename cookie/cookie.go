package cookie

import (
	"net/http"
	"time"
)

func RemoveCookie(w http.ResponseWriter, n string) {
  cookie := &http.Cookie{
    Name: n,
    Value: "",
    Path: "/",
    Expires: time.Unix(0, 0),
    HttpOnly: true,
  }
  http.SetCookie(w, cookie)
}

func SaveCookie(w http.ResponseWriter, n string, v string) {
  expire := time.Now().AddDate(0, 0, 1)
  cookie := &http.Cookie{
    Name:    n,
    Value:   v,
    Expires: expire,
  }
  http.SetCookie(w, cookie)
}

