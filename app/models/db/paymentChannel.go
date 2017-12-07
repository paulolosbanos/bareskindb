package db

type PaymentChannel struct {
	Id      int64  `db:"payment_channel_id" json:"payment-channel-id"`
	Name    string `db:"payment_channel_name" json:"payment-channel-name"`
	Account string `db:"payment_channel_account" json:"payment-channel-account"`
}
