package db

type SoldProducts struct {
	Id        int64 `db:"sold_product_id" json:"sold-product-id"`
	SaleId   int64 `db:"sale_id" json:"sale-id"`
	ProductId string `db:"product_id" json:"product-id"`
	Quantity  int64 `db:"product_quantity" json:"product-quantity"`
}