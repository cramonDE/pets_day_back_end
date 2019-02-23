package models

import "github.com/astaxie/beego/orm"

type User struct {
	Account      string `json:"account"`
	Name         string `json:"name"`
	City         string `json:"city"`
	Address      string `json:"address"`
	PhoneCall    string `json:"phone_call"`
	Mobile       string `json:"mobile"`
	Email        string `json:"email"`
	Job          string `json:"job"`
	Contact      string `json:"contact"`
	ContactPhone string `json:"contact_phone"`
	Income       int    `json:"income"`
	CreditRating string `json:"credit_rating"`
}


func UpdateUser(user *User, account string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE user SET account = ?, name = ?, city = ?, address = ?, phone_call = ?, mobile = ?, email = ?, job = ?, contact = ?, contact_phone = ?, income = ?, credit_rating = ? WHERE account = ?;", user.Account, user.Name, user.City, user.Address, user.PhoneCall, user.Mobile, user.Email, user.Job, user.Contact, user.ContactPhone, user.Income, user.CreditRating, account).Exec()
	return err
}
 

func GetOneUser(account string) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.Raw("SELECT * FROM user WHERE account = ?;", account).QueryRow(&user)
	return user, err
}
