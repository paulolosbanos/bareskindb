package db

import "time"

type Order struct {
	Id              int64     `db:"order_id" json:"-"`
	UserId          int64     `db:"user_id" json:"user-id"`
	Name            string    `db:"name" json:"name"`
	Address         string    `db:"address" json:"address"`
	Mobile          string    `db:"mobile" json:"mobile"`
	PaymentChannel  string    `db:"payment_channel" json:"payment-channel"`
	PaymentDate     time.Time `db:"payment_date" json:"payment-date"`
	ReferenceNumber string    `db:"reference_number" json:"reference-number"`
	OrderDate       time.Time `db:"order_date" json:"order-date"`
	OrderStatus     string    `db:"order_status" json:"order-status"`
}
