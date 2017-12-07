package controllers

import (
	"github.com/revel/revel"
	"bareskin-api/app/models/db"
	"bareskin-api/app/models/domain"
)

type UserCtrl struct {
	GorpController
}

//func (c UserCtrl) List() revel.Result {
//	return c.RenderJSON(users)
//}

func (c UserCtrl) parseUser() (db.User, error) {
	user := db.User{}
	err := c.Params.BindJSON(&user)
	//json.NewDecoder(c.Request.Body).Decode(&user)
	return user, err
}

func (c UserCtrl) Add() revel.Result {
	if user, err := c.parseUser(); err != nil {
		return c.RenderText("Unable to parse the User from JSON.")
	} else {
		// Validate the model
		user.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			return c.RenderText("You have error in your User.")
		} else {
			if err := c.Txn.Insert(&user); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.RenderJSON(user)
			}
		}
	}
}

func (c UserCtrl) Get(id int64) revel.Result {
	user := new(db.User)
	err := c.Txn.SelectOne(user,
		`SELECT * FROM User WHERE user_id = ?`, id)
	if err != nil {
		return c.RenderText("Error.  Item probably doesn't exist.")
	}
	return c.RenderJSON(user)
}

func (c UserCtrl) List() revel.Result {
	users, err := c.Txn.Select(db.User{}, `SELECT * FROM User`)
	if err != nil {
		return c.RenderJSON(domain.JSONError(err.Error()))
	}
	return c.RenderJSON(users)
}

func (c UserCtrl) Update(id int64) revel.Result {
	user, err := c.parseUser()
	if err != nil {
		return c.RenderText("Unable to parse the User from JSON.")
	}
	// Ensure the Id is set.
	user.Id = id
	success, err := c.Txn.Update(&user)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update bid item.")
	}
	return c.RenderText("Updated %v", id)
}

func (c UserCtrl) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&db.User{Id: id})
	if err != nil || success == 0 {
		return c.RenderText("Failed to remove User")
	}
	return c.RenderText("Deleted %v", id)
}
