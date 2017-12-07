package db

import "time"

type Sale struct {
	Id              int64     `db:"sale_id" json:"sale-id"`
	UserId          int64     `db:"user_id" json:"user-id"`
	ReferenceNumber string    `db:"reference_number" json:"reference-number"`
	BuyerName       string    `db:"buyer_name" json:"buyer-name"`
	BuyerNumber     string    `db:"buyer_number" json:"buyer-number"`
	SaleDate        time.Time `db:"sale_date" json:"sale-date"`
}
