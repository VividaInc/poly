package config

import (
  "io/ioutil"
  "os"
  "path"
  "path/filepath"
)

func ReadFile(filePart string) (string, error) {
  pwd, err := os.Getwd()
  if err != nil {
    return "", err
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

func ValidateConfComment(line string) bool {
  valid := false
  if len(line) != 0 {
    if []rune(line)[0] != 35 {
      valid = true
    }
  }
  return valid
}

func CloseFile(ptr *os.File) {
  defer ptr.Close()
}
