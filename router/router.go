package router

import (
  "net/http"
  "os"
  "path"
  "path/filepath"

  "dajour.christophe.org/controller"
  "dajour.christophe.org/handle"
)

type Router struct {
  Mux *http.ServeMux
}

func NewRouter() *Router {
  return &Router{}
}

func (r *Router) EstRoutes() {

  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  pwd = filepath.ToSlash(pwd)

  r.Mux.Handle("/public/", http.StripPrefix("/public/",
    http.FileServer(http.Dir(path.Join(pwd, "zmem.Resource/public/")))))

  // static routes
  r.Mux.Handle("/about", handle.DefaultHandler(controller.About))
  r.Mux.Handle("/board", handle.DefaultHandler(controller.Board))
  r.Mux.Handle("/community", handle.DefaultHandler(controller.Community))
  r.Mux.Handle("/connect", handle.DefaultHandler(controller.Connect))
  r.Mux.Handle("/contact-us", handle.DefaultHandler(controller.ContactUs))
  r.Mux.Handle("/discover", handle.DefaultHandler(controller.Discover))
  r.Mux.Handle("/forgot-password", handle.DefaultHandler(controller.ForgotPassword))
  r.Mux.Handle("/", handle.DefaultHandler(controller.Index))
  r.Mux.Handle("/sign-in", handle.DefaultHandler(controller.SignIn))
  r.Mux.Handle("/sign-up", handle.DefaultHandler(controller.SignUp))

  // oAuth routes
  r.Mux.Handle("/authenticate", handle.DefaultHandler(controller.Authenticate))

  // API routes
  r.Mux.Handle("/dp", handle.DefaultHandler(controller.DeletePost))
  r.Mux.Handle("/eu", handle.DefaultHandler(controller.ExistingUser))
  r.Mux.Handle("/logs", handle.DefaultHandler(controller.Logs))
  r.Mux.Handle("/messages", handle.DefaultHandler(controller.Messages))
  r.Mux.Handle("/nc", handle.DefaultHandler(controller.CreateCategory))
  r.Mux.Handle("/np", handle.DefaultHandler(controller.CreatePost))
  r.Mux.Handle("/nu", handle.DefaultHandler(controller.NewUser))
  r.Mux.Handle("/sign-off", handle.DefaultHandler(controller.SignOff))
  r.Mux.Handle("/upload", handle.DefaultHandler(controller.Upload))
}

func (r *Router) Redirect(res http.ResponseWriter, req *http.Request) {
  http.Redirect(res, req, "https://localhost", http.StatusTemporaryRedirect)
}

func (r *Router) NewMux() *http.ServeMux {
  return http.NewServeMux()
}
