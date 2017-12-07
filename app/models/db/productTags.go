package db

type ProductTags struct {
	Id					int64	`db:"product_tag_id" json:"product-tag-id"`
	productId			int64	`db:"product_id" json:"product-id"`
	tagId				int64	`db:"tag_id" json:"tag-id"`
}
