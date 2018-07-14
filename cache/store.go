package cache

import (
  "os"
  "path"
  "path/filepath"
  "strings"
)

func (s *CacheSession) transfer_file_to_store(name string, dir string, ext string) error {
  const (
    cacheDir string = ".ecce/cache/"
  )
  pwd, err := os.Getwd()
  if err != nil {
    return err
  }
  pwd = filepath.ToSlash(pwd)
  f1 := path.Join(dir, s.concat_pre_w_ext(name, ext))
  f2 := path.Join(pwd, cacheDir, s.Caches[name].(string))
  if err := CopyFile(f1, f2); err != nil {
    return err
  }
  return nil
}

func (s *CacheSession) parse_file_prefix(name string) string {
  return strings.Split(name, ".")[0]
}

func (s *CacheSession) concat_pre_w_ext(pre string, ext string) string {
  return strings.Join([]string{pre, ext}, "")
}
