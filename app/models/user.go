package models

import (
	"github.com/revel/revel"
)

type User struct {
	Id			int64	`db:"id" json:"id"`
	UID			string	`db:"uid" json:"uid"`
	Username	string	`db:"username" json:"username"`
	FullName	string	`db:"fullName" json:"fullname"`
	Email		string	`db:"email" json:"email"`
	Mobile		string	`db:"phoneNumber" json:"phone-number"`
	InviteCode	string	`db:"inviteCode" json:"invite-code"`
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


