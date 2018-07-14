package file

import (
	"io/ioutil"
	"os"
  "path"
  "path/filepath"
)

func ReadLoad() []byte {
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  pwd = filepath.ToSlash(pwd)
  SOURCEFILE := path.Join(pwd, "zmem.Program/zmem.App/db/user.load.json")
  content, _ := ioutil.ReadFile(SOURCEFILE)
  return content
}

func WriteLoad(LOAD []byte) {
  const (
  	SOURCEFILE string      = "zmem.Program/zmem.App/db/user.load.json"
		SOURCEPERM os.FileMode = 0777
	)
  _ = ioutil.WriteFile(SOURCEFILE, LOAD, SOURCEPERM)
}
