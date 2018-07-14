package controller

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "path"
  "path/filepath"
  "strconv"

  "github.com/VividaInc/poly/oauth"
  "github.com/VividaInc/poly/cookie"
  "github.com/VividaInc/poly/env"
  e "github.com/VividaInc/poly/err"
  "github.com/VividaInc/poly/log"
  "github.com/VividaInc/poly/protocol"
  "github.com/VividaInc/poly/protocol/db"
  "github.com/VividaInc/poly/render"
  "github.com/VividaInc/poly/store"
)

func About(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/about"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Title":   "What is the zmem secure messenger?",
    "SubTitle": "Chat securely with friends and family anywhere around the world in real-time. Don't miss out on your oppotunity. Your time is now.",
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["about"]),
  }
  if err := render.RenderHTML(w, "about", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func Authenticate(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/authenticate"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["authentication"]),
  }
  if err := render.RenderHTML(w, "authentication", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func Board(w http.ResponseWriter, r *http.Request) *e.RequestError {
  var (
    user *db.User
    err  *e.RequestError
  )
  err = validateController(r, "GET", "/board")
  if err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  user, err = store.GetUser(r)
  if err != nil {
    http.Redirect(w, r, "/sign-in", http.StatusTemporaryRedirect)
    return nil
  }
  opts := map[string]interface{}{
    "Fullname": user.Fullname,
    "Username": user.Username,
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["board"]),
  }
  if err := render.RenderHTML(w, "board", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }

  return nil
}

func Community(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/community"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Title":   "Community",
    "SubTitle": "",
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["community"]),
  }
  if err := render.RenderHTML(w, "community", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func Connect(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/connect"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Title":   "Connect",
    "SubTitle": "",
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["connect"]),
  }
  if err := render.RenderHTML(w, "connect", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func ContactUs(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/contact-us"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Title":   "Contact Us",
    "SubTitle": "",
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["contact-us"]),
  }
  if err := render.RenderHTML(w, "contact-us", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func CreateCategory(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "POST", "/nc"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  r.ParseForm()
  // action := r.FormValue("submit_post")
  category := r.FormValue("category")
  group    := r.FormValue("group")
  protocol.AddGroup(1, category, group)
  w.WriteHeader(http.StatusOK)
  return nil
}

func CreatePost(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "POST", "/np"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  r.ParseForm()
  // action := r.FormValue("submit_post")
  // category := r.FormValue("category")
  // group    := r.FormValue("group")
  sender    := r.FormValue("sender")
  recipient := r.FormValue("recipient")
  post      := r.FormValue("user_post")
  protocol.AddMessage(sender, recipient, post)
  w.WriteHeader(http.StatusOK)
  return nil
}

func DeletePost(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "POST", "/dp"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  r.ParseForm()
  // category := r.FormValue("category")
  // group    := r.FormValue("group")
  // user     := r.FormValue("user")
  id, err := strconv.Atoi(r.FormValue("_id"))
  if err != nil {
    return e.RETURNINTERNALERROR
  }
  protocol.DeleteMessage(1, id)
  w.WriteHeader(http.StatusOK)
  return nil
}

func Discover(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/discover"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Title":   "Discover",
    "SubTitle": "Chat securely with friends and family anywhere around the world in real-time. Don't miss out on your oppotunity. Your time is now.",
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["discover"]),
  }
  if err := render.RenderHTML(w, "discover", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }

  return nil
}

func ExistingUser(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "POST", "/eu"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  if ok := oauth.SignInUser(w, r); !ok {
    http.Redirect(w, r, "/sign-in", http.StatusTemporaryRedirect)
    return nil
  }
  http.Redirect(w, r, "/authenticate", http.StatusMovedPermanently)
  return nil
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/forgot-password"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["forgot-password"]),
  }
  if err := render.RenderHTML(w, "forgot-password", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func Index(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  opts := map[string]interface{}{
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["index"]),
  }
  if err := render.RenderHTML(w, "index", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func Logs(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/logs"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  pwd = filepath.ToSlash(pwd)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  content, _ := ioutil.ReadFile(path.Join(pwd, ".ecce/etc/user.log.json"))
  w.Write([]byte(content))
  return nil
}

func Messages(w http.ResponseWriter, r *http.Request) *e.RequestError {
  var (
    user *db.User
    err  *e.RequestError
  )
  err = validateController(r, "GET", "/messages")
  if err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  user, err = store.GetUser(r)
  if err != nil {
    return err
  }
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  w.Write(protocol.DecodeCategories(protocol.FindAllUserData(user.Id)))
  return nil
}

func NewUser(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "POST", "/nu"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  r.ParseForm()
  fullname := r.FormValue("fullname")
  username := r.FormValue("username")
  password := r.FormValue("password")
  protocol.StoreNewUserData(fullname, username, password)
  http.Redirect(w, r, "/authenticate", http.StatusMovedPermanently)
  return nil
}

func SignIn(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/sign-in"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  if ok := oauth.SignedInUser(r); ok {
    http.Redirect(w, r, "/authenticate", http.StatusTemporaryRedirect)
    return nil
  }
  opts := map[string]interface{}{
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["signin"]),
  }
  if err := render.RenderHTML(w, "sign-in", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func SignOff(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/sign-off"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  cookie.RemoveCookie(w, "ZmemAuthorization")
  http.Redirect(w, r, "/sign-in", http.StatusTemporaryRedirect)
  return nil
}

func SignUp(w http.ResponseWriter, r *http.Request) *e.RequestError {
  if err := validateController(r, "GET", "/sign-up"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  if ok := oauth.SignedInUser(r); ok {
    http.Redirect(w, r, "/authenticate", http.StatusTemporaryRedirect)
    return nil
  }
  opts := map[string]interface{}{
    "Resource": fmt.Sprintf("%v", env.CacheSession.Caches["signup"]),
  }
  if err := render.RenderHTML(w, "sign-up", opts); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Severe, err.Message)
    return err
  }
  return nil
}

func Upload(w http.ResponseWriter, r *http.Request) *e.RequestError {
  var filePath string
  if err := validateController(r, "POST", "/upload"); err != nil {
    logger := log.NewLogger()
    _ = logger.WriteToSystem(log.Info, err.Message)
    return err
  }
  r.ParseForm()
  r.ParseMultipartForm(32 << 20)
  fileType  := r.FormValue("type")
  sender    := r.FormValue("sender")
  recipient := r.FormValue("recipient")
  file, handler, err := r.FormFile("file")
  if err != nil {
    panic(err)
  }
  defer file.Close()
  if fileType == "image" {
    filePath = "./public/img/"
  } else if fileType == "video" {
    filePath = "./public/media/"
  }
  f, err := os.OpenFile(filePath + handler.Filename,
    os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  io.Copy(f, file)
  if fileType == "image" {
    protocol.AddMessage(sender, recipient, "<object class=\"msg__dpy-file\" style=\"background-image: url('" + filePath + handler.Filename + "');\"></object>")
  } else if fileType == "video" {
    protocol.AddMessage(sender, recipient, "<video width=\"320\" height=\"240\" autoplay><source type=\"video/mp4\" src=\"" + filePath + handler.Filename + "\" /></video>")
  }
  w.Header().Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  // w.Write([]byte("<img src=\"./public/img/" + handler.Filename + "\">"))
  return nil
}

/* ======================================================================== */

func validateController(r *http.Request, m string, p string) *e.RequestError {
  if r.Method != m {
    return e.RETURNINTERNALERROR
  }
  if r.URL.Path != p {
    return e.RETURNNOTFOUND
  }
  return nil
}

