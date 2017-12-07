package db

import (
	"github.com/revel/revel"
)

type User struct {
	Id          int64  `db:"user_id" json:"user-id"`
	UID         string `db:"uid" json:"uid"`
	Username    string `db:"username" json:"username"`
	Pass        string `db:"pass" json:"pass"`
	FullName    string `db:"fullName" json:"fullname"`
	Email       string `db:"email" json:"email"`
	Mobile      string `db:"phoneNumber" json:"phone-number"`
	Membership  string `db:"membership" json:"membership"`
	InviteCode  string `db:"inviteCode" json:"invite-code"`
	CreatedDate string `db:"created_date" json:"created_date"`
}

func (b *User) Validate(v *revel.Validation) {

	v.Check(b.Username,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))

	v.Check(b.FullName,
		revel.ValidRequired())

	v.Check(b.Mobile,
		revel.ValidRequired())
}
