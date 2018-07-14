package asset

import (
  "io/ioutil"
  "os"
  "path"
  "path/filepath"
)

func FindFile(name string, dir string) (bool, string, error) {
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  pwd = filepath.ToSlash(pwd)
  var filePath string
  found := false
  filePath = path.Join(dir, name)
  dir      = path.Join(pwd, dir)
  files, err := ioutil.ReadDir(dir)
  if err != nil {
    return found, "", err
  }
  for _, f := range files {
    if name == f.Name() {
      found = true
    }
  }
  return found, filePath, nil
}

func ReadFile(filePart string) (string, error) {
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  pwd = filepath.ToSlash(pwd)
  file := path.Join(pwd, filePart)
  if _, err := os.Stat(file); os.IsNotExist(err) {
    return "", err
  }
  content, err := ioutil.ReadFile(file)
  if err != nil {
    return "", err
  }
  return string(content), nil
}

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
