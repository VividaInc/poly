package main

import (
  "io/ioutil"
  "net/http"
  "testing"
)

func Test_Unit(t *testing.T) {

  const host string = "http://localhost:80"

  var validAddresses = []string{
    "/",
    "/about",
    "/authenticate",
    "/board",
    "/community",
    "/contact-us",
    "/discover",
    "/forgot-password",
    "/sign-in",
    "/sign-up",
  }

  for _, validAddress := range validAddresses {

    response, err := http.Get(host + validAddress)

    if err != nil {
      t.Error("Could not connect to host system")
    }

    defer response.Body.Close()

    content, _ := ioutil.ReadAll(response.Body)

    if len(content) < 0 {
      t.Error("Empty response from the host system")
    }
  }
}

func Test_Unit_Error(t *testing.T) {

  const host string = "http://localhost:80"

  var validAddresses = []string{
    "/random",
  }

  for _, validAddress := range validAddresses {

    response, err := http.Get(host + validAddress)

    if err != nil {
      t.Error("Could not connect to host system")
    }

    defer response.Body.Close()

    content, _ := ioutil.ReadAll(response.Body)

    if len(content) < 0 {
      t.Error("Empty response from the host system")
    }
  }
}
