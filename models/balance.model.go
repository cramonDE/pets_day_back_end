package models

import "github.com/astaxie/beego/orm"

type Balance struct {
	Account  string `json:"account"`
	Amount   int `json:"amount"`
	Interest int `json:"interest"`
}

func GetBalance(account string) (Balance, error) {
	o := orm.NewOrm()
	var balance Balance
	err := o.Raw("SELECT * FROM `balance` WHERE account = ? ", account).QueryRow(&balance)
	return balance, err
}

// func InsertLike(like *Like) (int64, error) {
// 	o := orm.NewOrm()
// 	res, err := o.Raw("INSERT INTO `like` (like_hotspot, like_user) VALUES (?, ?);", like.LikeHotspot, like.LikeUser).Exec()
// 	insertId, _ := res.LastInsertId()
// 	return insertId, err
// }

// func DeleteLike(id int) error {
// 	o := orm.NewOrm()
// 	_, err := o.Raw("DELETE FROM `like` WHERE like_id = ?", id).Exec()
// 	return err
// }

// func GetLikesByUserId(id int) (*[]Like, error) {
// 	o := orm.NewOrm()
// 	var likes []Like
// 	_, err := o.Raw("SELECT * FROM `like` WHERE like_user = ?", id).QueryRows(&likes)
// 	return &likes, err
// }
