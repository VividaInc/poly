package log

import (
  "log"
  "os"
  "path"
  "runtime"
)

type Logger struct {
}

const (
  Info   int = 0
  Debug  int = 1
  Severe int = 2
)

func NewLogger() *Logger {
  return &Logger{}
}

func (l *Logger) WriteToAccess(lvl int, msg string) error {
  const pth string = "zmem.Program/zmem.App/etc/access.log"
  switch (lvl) {
  case Info:
    err := l.PtrOut(pth, "[INFO] %s\n", msg)
    if err != nil {
      return err
    }
    break;
  case Debug:
    err := l.PtrOut(pth, "[DEBUG] %s\n", msg)
    if err != nil {
      return err
    }
    break;
  case Severe:
    err := l.PtrOut(pth, "[SEVERE] %s\n", msg)
    if err != nil {
      return err
    }
    break;
  default:
    break;
  }
  return nil
}

func (l *Logger) WriteToSystem(lvl int, msg string) error {
  const pth string = "zmem.Program/zmem.App/etc/system.log"
  buffer := make([]byte, 1<<16)
  switch (lvl) {
  case Info:
    err := l.PtrOut(pth, "[INFO] %s\n", msg)
    if err != nil {
      return err
    }
    break;
  case Debug:
    err := l.PtrOut(pth, "[DEBUG] %s\n", msg)
    if err != nil {
      return err
    }
    break;
  case Severe:
    err := l.PtrOut(pth, "[SEVERE] %s\n", msg)
    if err != nil {
      return err
    }
    runtime.Stack(buffer, true)
    err = l.PtrOut(pth, "%s\n", string(buffer))
    if err != nil {
      return err
    }
    break;
  default:
    break;
  }
  return nil
}

func (l *Logger) PtrOut(pth string, prm string, msg string) error {
  var ttl string = path.Join("C:/Users/Da'Jour Christophe/Documents/zmem/", pth)
  if _, err := os.Stat(ttl); os.IsNotExist(err) {
    return err
  }
  ptr, err := WriteToFile(ttl)
  if err != nil {
    return err
  }
  log.SetOutput(ptr)
  log.Printf(prm, msg)
  CloseFile(ptr)
  return nil
}

func (l *Logger) StdOut(prm string, msg string) {
  log.Printf(prm, msg)
}

