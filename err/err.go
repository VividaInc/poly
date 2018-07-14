package err

type RequestError struct {
  Err        error
  Message    string
  StatusCode int
}
