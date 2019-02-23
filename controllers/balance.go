package controllers

import (
	Models "pets-day/models"

	"github.com/astaxie/beego"
)

type BalanceController struct {
	beego.Controller
}

func (c *BalanceController) Get() {
	account := c.Ctx.Input.Params()[":account"]
	balance, err := Models.GetBalance(account)
	if nil == err {
		c.Data["json"] = balance
	}
	c.ServeJSON()
}

// func (c *LikeController) Post() {
// 	var like Models.Like
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &like)
// 	insertId, err := Models.InsertLike(&like)
// 	if nil == err {
// 		c.Data["json"] = insertId
// 	}
// 	c.ServeJSON()
// }

// func (c *LikeController) Delete() {
// 	id, _ := strconv.Atoi(c.Ctx.Input.Params()[":id"])
// 	err := Models.DeleteLike(id)
// 	if nil == err {
// 		c.Data["json"] = true
// 	} else {
// 		c.Data["json"] = false
// 	}
// 	c.ServeJSON()
// }
