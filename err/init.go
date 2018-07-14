package err

import (
	"net/http"
)

var (
  RETURNNOTFOUND      *RequestError
  RETURNINTERNALERROR *RequestError
)

func init() {
	RETURNNOTFOUND      = &RequestError{nil,"Page Not Found",http.StatusNotFound}
	RETURNINTERNALERROR = &RequestError{nil,"Internal Server Error",http.StatusInternalServerError}
}
