package protocol

import (
	"encoding/json"

	"github.com/VividaInc/poly/protocol/db"
)

func DecodeCategories(categories []db.Category) []byte {
	content, err := json.Marshal(categories)
	if err != nil {
		panic(err)
	}
	return content
}
