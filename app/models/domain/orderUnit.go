package domain

import "bareskin-api/app/models/db"

type OrderUnit struct {
	Product  db.Product `json:"product"`
	Quantity int64      `json:"quantity"`
}
