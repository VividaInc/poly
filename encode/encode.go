package encode

import "encoding/json"

func EncodeUsers(a []byte) []User {
  users := make([]User, 0)
  _ = json.Unmarshal(a, &users)
  return users
}

func EncodeCategories(a []byte) []Category {
  categories := make([]Category, 0)
  _ = json.Unmarshal(a, &categories)
  return categories
}

func DecodeCategories(a []Category) []byte {
  eCategories, _ := json.MarshalIndent(a, "", "  ")
  return eCategories
}

