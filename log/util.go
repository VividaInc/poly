package log

import "os"

func WriteToFile(ttl string) (*os.File, error) {
  var prm os.FileMode = 0666
  f, err := os.OpenFile(ttl, os.O_RDWR | os.O_CREATE | os.O_APPEND, prm)
  if err != nil {
    return nil, err
  }
  return f, nil
}

func CloseFile(ptr *os.File) {
  defer ptr.Close()
}
