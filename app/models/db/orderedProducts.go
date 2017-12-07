package db

type OrderedProducts struct {
	Id        int64 `db:"order_product_id" json:"order-product-id"`
	OrderId   int64 `db:"order_id" json:"order-id"`
	ProductId string `db:"product_id" json:"product-id"`
	Quantity  int64 `db:"product_quantity" json:"product-quantity"`
}
