package file

import (
  "io/ioutil"
  "os"
  "path"
  "path/filepath"
  "strings"
)

func RemoveAllFiles(dir string, ext string) {
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  pwd         = filepath.ToSlash(pwd)
  dir         = path.Join(pwd, dir)
  files, err := ioutil.ReadDir(dir)
  if err != nil {
    panic(err)
  }
  for _, f := range files {
    name        := f.Name()
    absfilepath := path.Join(dir, name)
    absfileext  := Parse(name)
    if absfileext == ext {
      if err := os.Remove(absfilepath); err != nil {
        panic(err)
      }
    }
  }
}

func Parse(name string) string {
  const seperator string = "."
  var ext string
  exts := strings.Split(name, seperator)
  if len(exts) > 2 {
    ext = strings.Join([]string{
      seperator,
      exts[len(exts) - 2],
      exts[len(exts) - 1],
    }, seperator)
  } else {
    ext = strings.Join([]string{
      seperator,
      exts[len(exts) - 1],
    }, seperator)
  }
  ext = Validate(ext)
  return ext
}

func Validate(extension string) string {
  const fault string = ".."
  var buf []rune
  if strings.Contains(extension, fault) {
    for i, extPart := range []rune(extension) {
      if i > 0 {
        buf = append(buf, extPart)
      }
    }
  }
  extension = string(buf)
  return extension
}
