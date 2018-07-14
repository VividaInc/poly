package protocol

import (
  "dajour.christophe.org/protocol/db"
)

func AddMessage(s string, r string, m string) {
  db.SaveRecipientNewMessage(db.GetUserFromDatabase(r).Uid, m)
  db.SaveSenderNewMessage(db.GetUserFromDatabase(s).Uid, m)
}

func DeleteMessage(uid int, mid int) {
  db.DeleteUserMessage(uid, mid)
}

func AddGroup(uid int, category string, title string) {
  db.SaveNewGroup(uid, category, title)
}

func FindAllUserData(uid int) []db.Category {
  e := db.GetAllCategories(uid)
  g := db.GetAllGroups(uid)
  v := db.GetAllViews(uid)
  m := db.GetAllMessagesFromUser(uid)
  d := db.BuildView(m, v)
  b := db.BuildGroup(d, g)
  c := db.BuildCategory(b, e)
  return c
}

func StoreNewUserData(fullname string, username string, password string) {
  db.SaveNewUser(fullname, username, password)
}
