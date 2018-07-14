package protocol

import (
	"encoding/json"

	"dajour.christophe.org/protocol/db"
)

func DecodeCategories(categories []db.Category) []byte {
	content, err := json.Marshal(categories)
	if err != nil {
		panic(err)
	}
	return content
}
