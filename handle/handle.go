package handle

import (
  "io/ioutil"
	"net/http"

	"dajour.christophe.org/config"
  "dajour.christophe.org/env"
  "dajour.christophe.org/err"
  "dajour.christophe.org/log"
  "dajour.christophe.org/render"
)

type Handle interface {
}

type Handler struct {
}

type DefaultHandler func(http.ResponseWriter, *http.Request) *err.RequestError

func NewHandler() *Handler {
  return &Handler{}
}

func (h DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  conf := config.HandConf(".ecce/etc/socket.conf")
  var (
    logger  *log.Logger
    opts    *ResponseError
    content []byte
  )
  logger = log.NewLogger()
  if err := h(w, r); err != nil {
    switch err.StatusCode {
    case http.StatusNotFound:
      stack := stack(3)
      _ = logger.WriteToSystem(log.Severe, string(stack))
      if env.Env == env.Development {
        opts = &ResponseError{
          string(stack),
          err.Message,
        }
      } else {
        content, _ = ioutil.ReadFile("C:/Users/Da'Jour Christophe/Documents/zmem/" + conf.PageNotFoundFile)
      }
      break;
    case http.StatusForbidden:
      stack := stack(3)
      _ = logger.WriteToSystem(log.Severe, string(stack))
      if env.Env == env.Development {
        opts = &ResponseError{
          string(stack),
          err.Message,
        }
      } else {
        content, _ = ioutil.ReadFile("C:/Users/Da'Jour Christophe/Documents/zmem/" + conf.ForbiddenFile)
      }
      break;
    case http.StatusInternalServerError:
      stack := stack(3)
      _ = logger.WriteToSystem(log.Severe, string(stack))
      if env.Env == env.Development {
        opts = &ResponseError{
          string(stack),
          err.Message,
        }
      } else {
        content, _ = ioutil.ReadFile("C:/Users/Da'Jour Christophe/Documents/zmem/" + conf.InternalServerErrorFile)
      }
      break;
    default:
      break;
    }
    if env.Env == env.Development {
      render.RenderErr(w, err.StatusCode, opts)
    } else {
      w.Header().Set("Content-Type", "text/html; charset=UTF-8")
      w.WriteHeader(err.StatusCode)
      w.Write(content)
    }
  }
}
