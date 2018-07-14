package asset

import (
  "net/http"

  "github.com/VividaInc/poly/file"
)

type Asset interface {
  NewHandler()
  Fetch()
}

type AssetHandler struct {
}

func NewHandler() *AssetHandler {
  return &AssetHandler{}
}

func (h *AssetHandler) Fetch(r *http.Request) (string, string, error) {
  const (
    dir string = ".ecce/cache/"
    sep string = "."
  )
  var (
    name    string = r.URL.Path[1:]
    content string
    ext     string
  )
  if ok, assetPath, serr := FindFile(name, dir); ok {
    var rerr error
    content, rerr = ReadFile(assetPath)
    if rerr != nil {
      return "", "", rerr
    }
    ext = file.Parse(assetPath)
  } else if serr != nil {
    return "", "", serr
  }
  return content, ext, nil
}
