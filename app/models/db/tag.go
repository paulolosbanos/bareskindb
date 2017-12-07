package db

type Tag struct {
	Id          int64     `db:"tag_id" json:"tag-id"`
	TagName     string    `db:"tag_name" json:"tag-name"`
}
