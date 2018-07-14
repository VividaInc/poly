package db

type Message struct {
  Id       int    `json:"_id,omitempty"`
  Uid      int    `json:"_uid,omitempty"`
  Did      int    `json:"_did,omitempty"`
  Rid      int    `json:"_rid,omitempty"`
  Received string `json:"received,omitempty"`
  Sent     string `json:"sent,omitempty"`
}

type View struct {
  Id       int       `json:"_id,omitempty"`
  Uid      int       `json:"_uid,omitempty"`
  Did      int       `json:"_did,omitempty"`
  Rid      int       `json:"_rid,omitempty"`
  User     string    `json:"user,omitempty"`
  Username string    `json:"username,omitempty"`
  Body     string    `json:"body,omitempty"`
  Content  []Message `json:"content"`
  Type     string    `json:"_type,omitempty"`
  Status   string    `json:"_status,omitempty"`
}

type Group struct {
  Id            int    `json:"_id,omitempty"`
  Uid           int    `json:"_uid,omitempty"`
  Did           int    `json:"_did,omitempty"`
  Rid           int    `json:"_rid,omitempty"`
  Title         string `json:"title"`
  Notifications int    `json:"notifications"`
  Views         []View `json:"views"`
}

type Category struct {
  Id      int     `json:"_id,omitempty"`
  Uid     int     `json:"_uid,omitempty"`
  Did     int     `json:"_did,omitempty"`
  Rid     int     `json:"_rid,omitempty"`
  Title   string  `json:"title,omitempty"`
  Content []Group `json:"content"`
}

type User struct {
  Id       int        `json:"_id,omitempty"`
  Uid      int        `json:"_uid,omitempty"`
  Did      int        `json:"_did,omitempty"`
  Rid      int        `json:"_rid,omitempty"`
  Fullname string     `json:"fullname,omitempty"`
  Username string     `json:"username,omitempty"`
  Password string     `json:"password,omitempty"`
  Content  []Category `json:"content"`
}
