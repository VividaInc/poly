package socket

import (
  "fmt"
  "net/http"
  "os"
  "path"
  "path/filepath"
  "strings"
  "time"

  "dajour.christophe.org/asset"
  "dajour.christophe.org/config"
  "dajour.christophe.org/router"
)

func handleCache(w http.ResponseWriter, r *http.Request) {
  mimes := map[string]interface{}{
    ".css": "text/css; charset=UTF-8",
    ".min.css": "text/css; charset=UTF-8",
  }
  assetHandler := asset.NewHandler()
  if cnt, ext, err := assetHandler.Fetch(r); err != nil {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("Cound not find asset."))
  } else {
    contentType := fmt.Sprintf("%v", mimes[ext])
    w.Header().Set("Content-Type", contentType)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(cnt))
  }
}

func runOnSecureCache(cert string, key string) {
  if err := http.ListenAndServeTLS("127.0.0.1:8443", cert, key, http.HandlerFunc(handleCache)); err != nil {
    panic(err)
  }
}

func RunOnSecureAddr() error {
  conf := config.ConfSock(".ecce/etc/socket.conf")
  rt    := router.NewRouter()
  prot, mux := bindOnAddr(rt, 1)
  rt.Mux = mux
  rt.EstRoutes()
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  pwd = filepath.ToSlash(pwd)
  cert := path.Join(pwd, conf.TLSCert)
  key  := path.Join(pwd, conf.TLSKey)
  go runOnSecureCache(cert, key)
  if err := http.ListenAndServeTLS(prot.Addr, cert, key, mux); err != nil {
    return err
  }
  return nil
}

func RunOnUnsecureAddr() error {
  rt  := router.NewRouter()
  prot, _ := bindOnAddr(rt, 0)
  handler := http.HandlerFunc(rt.Redirect)
  if err := http.ListenAndServe(prot.Addr, handler); err != nil {
    return err
  }
  return nil
}

func bindOnAddr(rt *router.Router, t int) (*http.Server, *http.ServeMux) {
  var addr string
  conf := config.ConfSock(".ecce/etc/socket.conf")
  if t == 0 {
    addr = addrConcat(conf.Host, conf.Redirect)
  } else if  t == 1 {
    addr = addrConcat(conf.Host, conf.Port)
  }
  mux   := rt.NewMux()
  read  := conf.ReadTimeout * time.Second
  write := conf.WriteTimeout * time.Second
  max   := 1 << 20
  prot := &http.Server{
    Addr:           addr,
    Handler:        mux,
    ReadTimeout:    read,
    WriteTimeout:   write,
    MaxHeaderBytes: max,
  }
  return prot, mux
}

func addrConcat(h string, p string) string {
  var (
    addr   string
    buffer []string
  )
  host := os.Getenv("ZMEMHOST")
  port := os.Getenv("ZMEMPORT")
  if len(host) == 0 {
    host = h
  }
  if len(port) == 0 {
    port = p
  }
  buffer = []string{host, ":", port}
  addr   = strings.Join(buffer, "")
  return addr
}
