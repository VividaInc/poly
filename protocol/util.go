package protocol

import (
	"encoding/json"

	"polypentec/protocol/db"
)

func DecodeCategories(categories []db.Category) []byte {
	content, err := json.Marshal(categories)
	if err != nil {
		panic(err)
	}
	return content
}
