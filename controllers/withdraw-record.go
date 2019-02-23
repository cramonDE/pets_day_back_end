package controllers

import (
	"encoding/json"
	Models "pets-day/models"

	"github.com/astaxie/beego"
)

type WithdrawRecordController struct {
	beego.Controller
}

func (c *WithdrawRecordController) Get() {
	account := c.Ctx.Input.Params()[":account"]
	balance, err := Models.GetWithdrawRecords(account)
	if nil == err {
		c.Data["json"] = &balance
	}
	c.ServeJSON()
}

func (c *WithdrawRecordController) Post() {
	var withdrawRecord Models.WithdrawRecord
	json.Unmarshal(c.Ctx.Input.RequestBody, &withdrawRecord)
	insertId, err := Models.InsertWithdrawRecord(&withdrawRecord)
	if nil == err {
		c.Data["json"] = insertId
	}
	c.ServeJSON()
}
