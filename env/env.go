package env

import (
  "io/ioutil"
  "os"
  "path"
  "path/filepath"
  "poly/cache"
  "poly/config"
)

const (
  Development string = "DEVELOPMENT"
  Production  string = "PRODUCTION"
  Staging     string = "STAGING"
)

var (
	CacheSession *cache.CacheSession
	Env          string
)

func init() {
  pwd, err := os.Getwd()
  if err != nil {
  	panic(err)
  }
  pwd   = filepath.ToSlash(pwd)
  conf := config.EnvConf(".ecce/etc/socket.conf")
  Env   = conf.Environment
  dir  := path.Join(pwd, "zmem.Resource/public/css/")
  files, err := ioutil.ReadDir(dir)
  if err != nil {
    panic(err)
  }
  CacheSession = cache.NewSession()
  for _, f := range files {
    if err := CacheSession.Save(f.Name(), dir); err != nil {
      panic(err)
    }
  }
}
