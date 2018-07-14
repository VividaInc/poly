package main

import "dajour.christophe.org/socket"

func main() {
  go func() {
    if err := socket.RunOnUnsecureAddr(); err != nil {
      panic(err)
    }
  }()
  if err := socket.RunOnSecureAddr(); err != nil {
    panic(err)
  }
}
