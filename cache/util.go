package cache

import "io/ioutil"

func CopyFile(f1 string, f2 string) error {
  cnt, err := ioutil.ReadFile(f1)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(f2, cnt, 0644)
  if err != nil {
    return err
  }
  return nil
}
