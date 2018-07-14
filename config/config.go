package config

import (
  "strconv"
  "strings"
  "time"
)

type SocketConfig struct {
  Host         string
  Port         string
  Redirect     string
  ReadTimeout  time.Duration
  WriteTimeout time.Duration
  TLSCert      string
  TLSKey       string
}

type HandlerConfig struct {
  PageNotFoundFile        string
  ForbiddenFile           string
  InternalServerErrorFile string
}

type EnvConfig struct {
  Environment string
}

func EnvConf(filePart string) *EnvConfig {
  const (
    ENVIRONMENTSETTING string = "ENVIRONMENT"
  )
  conf := &EnvConfig{}
  content, err := ReadFile(filePart)
  if  err != nil {
    panic(err)
  }
  lines := strings.Split(content, "\n")
  for _, line := range lines {
    if len(line) > 0 {
      if ok := ValidateConfComment(line); ok {
        line = strings.Trim(line, "\n")
        if strings.Contains(line, "[") || strings.Contains(line, "]") {
          // do something later ...
        } else {
          comps := strings.Split(line, " ")
          if len(comps) > 1 {
            key := comps[0]
            val := strings.TrimSpace(comps[1])
            if len(key) != 0 && len(val) != 0 {
              switch key {
              case ENVIRONMENTSETTING:
                conf.Environment = val
                break;
              default:
                break;
              }
            }
          }
        }
      }
    }
  }
  return conf
}

func HandConf(filePart string) *HandlerConfig {
  const (
    PAGENOTFOUNDFILE        string = "PAGENOTFOUNDDOM"
    FORBBIDENFILE           string = "STATUSFORBIDDENDOM"
    INTERNALSERVERERRORFILE string = "INTERNALSERVERERRORDOM"
  )
  conf := &HandlerConfig{}
  content, err := ReadFile(filePart)
  if  err != nil {
    panic(err)
  }
  lines := strings.Split(content, "\n")
  for _, line := range lines {
    if len(line) > 0 {
      if ok := ValidateConfComment(line); ok {
        line = strings.Trim(line, "\n")
        if strings.Contains(line, "[") || strings.Contains(line, "]") {
          // do something later ...
        } else {
          comps := strings.Split(line, " ")
          if len(comps) > 1 {
            key := comps[0]
            val := strings.TrimSpace(comps[1])
            if len(key) != 0 && len(val) != 0 {
              switch key {
              case PAGENOTFOUNDFILE:
                conf.PageNotFoundFile = val
                break;
              case FORBBIDENFILE:
                conf.ForbiddenFile = val
                break;
              case INTERNALSERVERERRORFILE:
                conf.InternalServerErrorFile = val
                break;
              default:
                break;
              }
            }
          }
        }
      }
    }
  }
  return conf
}

func ConfSock(filePart string) *SocketConfig {
  const (
    HOST           string = "HOST"
    PORT           string = "PORT"
    REDIRECT       string = "REDIRECT"
    READTIMEOUT    string = "READTIMEOUT"
    WRITETIMEOUT   string = "WRITETIMEOUT"
    TLSCERTIFICATE string = "TLSCERTIFICATE"
    TLSKEY         string = "TLSKEY"
  )
  conf := &SocketConfig{}
  content, err := ReadFile(filePart)
  if  err != nil {
    panic(err)
  }
  lines := strings.Split(content, "\n")
  for _, line := range lines {
    if len(line) > 0 {
      if ok := ValidateConfComment(line); ok {
        line = strings.Trim(line, "\n")
        if strings.Contains(line, "[") || strings.Contains(line, "]") {
          // do something later ...
        } else {
          comps := strings.Split(line, " ")
          if len(comps) > 1 {
            key := comps[0]
            val := strings.TrimSpace(comps[1])
            if len(key) != 0 && len(val) != 0 {
              switch key {
              case HOST:
                conf.Host = val
                break;
              case PORT:
                conf.Port = val
                break;
              case REDIRECT:
                conf.Redirect = val
                break;
              case READTIMEOUT:
                tmp, _ := strconv.Atoi(val)
                conf.ReadTimeout = time.Duration(tmp)
                break;
              case WRITETIMEOUT:
                tmp, _ := strconv.Atoi(val)
                conf.WriteTimeout = time.Duration(tmp)
                break;
              case TLSCERTIFICATE:
                conf.TLSCert = val
                break;
              case TLSKEY:
                conf.TLSKey = val
                break;
              default:
                break;
              }
            }
          }
        }
      }
    }
  }
  return conf
}
