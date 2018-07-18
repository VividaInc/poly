package db

import (
	"database/sql"
  "os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
  "github.com/VividaInc/poly/env"
)

type DatabaseConnection struct {
}

func NewDatabaseConnection() *DatabaseConnection {
  return &DatabaseConnection{}
}

/* ======================================================================== */

func GetAllCategories(uid int) []Category {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  id := strconv.Itoa(uid)

  rows, err := conn.Query("SELECT * FROM categories WHERE uid=" + id)

  if err != nil {
    panic(err)
  }

  var category Category

  categories := make([]Category, 0)

  for rows.Next() {

    if err = rows.Scan(&category.Id, &category.Uid, &category.Did, &category.Rid, &category.Title); err != nil {
      panic(err)
    }

    if category.Uid == uid {
      categories = append(categories, category)
    }
  }

  conn.Close()

  return categories
}

func GetCategory(uid int, title string) Category {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  id := strconv.Itoa(uid)

  var sql string = "SELECT * FROM categories WHERE uid=" + id + " AND title=\"" + title + "\""

  rows, err := conn.Query(sql)

  if err != nil {
    panic(err)
  }

  var category Category

  for rows.Next() {

    if err = rows.Scan(&category.Id, &category.Uid, &category.Did, &category.Rid, &category.Title); err != nil {
      panic(err)
    }
  }

  conn.Close()

  return category
}

func GetAllGroups(uid int) []Group {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  id := strconv.Itoa(uid)

  rows, err := conn.Query("SELECT * FROM groups WHERE uid=" + id)

  if err != nil {
    panic(err)
  }

  var group Group

  groups := make([]Group, 0)

  for rows.Next() {

    if err = rows.Scan(&group.Id, &group.Uid, &group.Did, &group.Rid, &group.Title, &group.Notifications); err != nil {
      panic(err)
    }

    if group.Uid == uid {
      groups = append(groups, group)
    }
  }

  conn.Close()

  return groups
}

func GetAllViews(uid int) []View {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  id := strconv.Itoa(uid)

  rows, err := conn.Query("SELECT * FROM views WHERE uid=" + id)

  if err != nil {
    panic(err)
  }

  var view View

  views := make([]View, 0)

  for rows.Next() {

    if err = rows.Scan(&view.Id, &view.Uid, &view.Did, &view.Rid, &view.User, &view.Username, &view.Body, &view.Type, &view.Status); err != nil {
      panic(err)
    }

    if view.Uid == uid {
      views = append(views, view)
    }
  }

  conn.Close()

  return views
}

func GetAllMessagesFromUser(uid int) []Message {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  id := strconv.Itoa(uid)

  rows, err := conn.Query("SELECT * FROM messages WHERE uid=" + id)

  if err != nil {
    panic(err)
  }

  var msg Message

  msgs := make([]Message, 0)

  for rows.Next() {

    if err = rows.Scan(&msg.Id, &msg.Uid, &msg.Did, &msg.Rid, &msg.Sent, &msg.Received); err != nil {
      panic(err)
    }

    if msg.Uid == uid {
      msgs = append(msgs, msg)
    }
  }

  conn.Close()

  return msgs
}

func GetAllMessages() []Message {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  rows, err := conn.Query("SELECT * FROM messages")

  if err != nil {
    panic(err)
  }

  var msg Message

  msgs := make([]Message, 0)

  for rows.Next() {

    if err = rows.Scan(&msg.Id, &msg.Uid, &msg.Did, &msg.Rid, &msg.Sent, &msg.Received); err != nil {
      panic(err)
    }

    msgs = append(msgs, msg)
  }

  conn.Close()

  return msgs
}

func GetAllUsersFromDatabase() []User {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  rows, err := conn.Query("SELECT * FROM users")

  if err != nil {
    panic(err)
  }

  var user User

  users := make([]User, 0)

  for rows.Next() {

    if err = rows.Scan(&user.Id, &user.Uid, &user.Did, &user.Rid, &user.Fullname, &user.Username, &user.Password); err != nil {
      panic(err)
    }

    users = append(users, user)
  }

  conn.Close()

  return users
}

func GetUserFromDatabase(username string) *User {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  rows, err := conn.Query("SELECT * FROM users WHERE username=\"" + username + "\"")

  if err != nil {
    panic(err)
  }

  var user User

  for rows.Next() {

    if err = rows.Scan(&user.Id, &user.Uid, &user.Did, &user.Rid, &user.Fullname, &user.Username, &user.Password); err != nil {
      panic(err)
    }
  }

  conn.Close()

  return &user
}

/* ======================================================================== */

func SaveRecipientNewMessage(uid int, msg string) int64 {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  messages := GetAllMessages()

  stmt, err := conn.Prepare("INSERT messages SET id=?, uid=?, did=?, rid=?, sender=?, recipient=?")

  if err != nil {
    panic(err)
  }

  res, err := stmt.Exec(messages[len(messages) - 1].Id + 1, uid, 51991, 41987, "", msg)

  if err != nil {
    panic(err)
  }

  id, err := res.LastInsertId()

  if err != nil {
    panic(err)
  }

  conn.Close()

  return id
}

func SaveSenderNewMessage(uid int, msg string) int64 {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  messages := GetAllMessages()

  stmt, err := conn.Prepare("INSERT messages SET id=?, uid=?, did=?, rid=?, sender=?, recipient=?")

  if err != nil {
    panic(err)
  }

  res, err := stmt.Exec(messages[len(messages) - 1].Id + 1, uid, 51991, 41987, msg, "")

  if err != nil {
    panic(err)
  }

  id, err := res.LastInsertId()

  if err != nil {
    panic(err)
  }

  conn.Close()

  return id
}

func DeleteUserMessage(uid int, id int) int64 {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  stmt, err := conn.Prepare("DELETE FROM messages WHERE uid=? AND did=? AND id=?")

  if err != nil {
    panic(err)
  }

  res, err := stmt.Exec(uid, 51991, id)

  if err != nil {
    panic(err)
  }

  affects, err := res.RowsAffected()

  if err != nil {
    panic(err)
  }

  conn.Close()

  return affects
}

func SaveNewGroup(uid int, category string, title string) int64 {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  groups := GetAllGroups(uid)
  cgy    := GetCategory(uid, category)

  stmt, err := conn.Prepare("INSERT messages SET id=?, uid=?, did=?, rid=?, title=?, notifications=?")

  if err != nil {
    panic(err)
  }

  res, err := stmt.Exec(groups[len(groups) - 1].Id + 1, uid, cgy.Did, cgy.Rid, title, 0)

  if err != nil {
    panic(err)
  }

  id, err := res.LastInsertId()

  if err != nil {
    panic(err)
  }

  conn.Close()

  return id
}

func SaveNewUser(fullname string, username string, password string) int64 {

  db := NewDatabaseConnection()

  conn, err := db.Connect()

  if err != nil {
    panic(err)
  }

  var id  int
  var uid int
  var did int
  var rid int

  users := GetAllUsersFromDatabase()

  id  = users[len(users) -1].Id + 1
  uid = id
  did = users[len(users) -1].Did + 1
  rid = users[len(users) -1].Rid + 1

  stmt, err := conn.Prepare("INSERT users SET id=?, uid=?, did=?, rid=?, fullname=?, username=?, password=?")

  if err != nil {
    panic(err)
  }

  res, err := stmt.Exec(id, uid, did, rid, fullname, username, password)

  if err != nil {
    panic(err)
  }

  _id, err := res.LastInsertId()

  if err != nil {
    panic(err)
  }

  conn.Close()

  return _id
}

/* ======================================================================== */

func (db *DatabaseConnection) Connect() (*sql.DB, error) {
  const (
    driver  string = "mysql"
    options string = "root:Dajour98*@/messages?charset=utf8"
  )
  url := os.Getenv("CLEARDB_DATABASE_URL")
  if env.Env == "DEVELOPMENT" || len(url) == 0{
    url = options
  }
  conn, err := sql.Open(driver, url)
  if err != nil {
    return nil, err
  }
  if err := conn.Ping(); err != nil {
    return nil, err
  }
  return conn, nil
}

/* ======================================================================== */

func BuildCategory(a []Group, b []Category) []Category {
  d := make([]Category, 0)
  for _, c := range b {
    for _, e := range a {
      if e.Did == c.Rid {
        c.Content = append(c.Content, e)
      }
    }
    d = append(d, c)
  }
  return d
}

func BuildGroup(a []View, b []Group) []Group {
  d := make([]Group, 0)
  for _, c := range b {
    for _, e := range a {
      if e.Did == c.Rid {
        c.Views = append(c.Views, e)
      }
    }
    d = append(d, c)
  }
  return d
}

func BuildView(a []Message, b []View) []View {
  d := make([]View, 0)
  for _, c := range b {
    for _, e := range a {
      if e.Did == c.Rid {
        c.Content = append(c.Content, e)
      }
    }
    d = append(d, c)
  }
  return d
}
