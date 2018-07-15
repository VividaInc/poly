package render

import (
  "html/template"
  "net/http"
  "path"
  "strings"

  e "github.com/VividaInc/poly/err"
)

func RenderHTML(w http.ResponseWriter, n string, o interface{}) *e.RequestError {
  file := strings.Join([]string{"/zmem.Resource/private/template/", n, ".min.tmpl"}, "")
  tmpl, err := template.ParseFiles(
    path.Join("/zmem.Resource/private/template/layout.min.tmpl"),
    path.Join("/zmem.Resource/private/template/navigation.min.tmpl"),
    path.Join("/zmem.Resource/private/template/banner.min.tmpl"),
    path.Join("/zmem.Resource/private/template/footer.min.tmpl"), file,
  )
  if err != nil {
    panic(err)
    return e.RETURNINTERNALERROR
  }
  w.Header().Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  tmpl.ExecuteTemplate(w, "layout", o)
  return nil
}

func RenderErr(w http.ResponseWriter, s int, o interface{}) {
  file := strings.Join([]string{"/zmem.Resource/private/template/error.tmpl"}, "")
  tmpl, err := template.ParseFiles(file)
  if err != nil {
    panic(err)
  }
  w.Header().Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(s)
  tmpl.Execute(w, o)
}
