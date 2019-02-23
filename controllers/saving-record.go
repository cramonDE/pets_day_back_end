package controllers

import (
	"encoding/json"
	Models "pets-day/models"

	"github.com/astaxie/beego"
)

type SavingRecordController struct {
	beego.Controller
}

func (c *SavingRecordController) Get() {
	account := c.Ctx.Input.Params()[":account"]
	balance, err := Models.GetSavingRecords(account)
	if nil == err {
		c.Data["json"] = &balance
	}
	c.ServeJSON()
}

func (c *SavingRecordController) Post() {
	var savingRecord Models.SavingRecord
	json.Unmarshal(c.Ctx.Input.RequestBody, &savingRecord)
	insertId, err := Models.InsertSavingRecord(&savingRecord)
	if nil == err {
		c.Data["json"] = insertId
	}
	c.ServeJSON()
}

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
