package db

type Product struct {
	Id        int64  `db:"id" json:"-"`
	ProductId string `db:"product_id" json:"product-id"`
	Name      string `db:"product_name" json:"product-name"`
	Cost      int64  `db:"product_cost_unit" json:"product-cost-unit"`
	Srp       int64  `db:"product_srp_unit" json:"product-srp-unit"`
	ImageName string `db:"product_image_name" json:"product-image-name"`
	Status    string `db:"product_status" json:"product-status"`
}
