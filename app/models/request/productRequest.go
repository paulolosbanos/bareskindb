package request

import "bareskin-api/app/models/db"

type ProductRequest struct {
	Products []db.Product `json:"products"`
}