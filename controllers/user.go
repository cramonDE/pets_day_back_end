package controllers

import (
	// "fmt"
	"encoding/json"
	Models "pets-day/models"
	// "strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) Patch() {
	var user Models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	account := c.Ctx.Input.Params()[":account"]
	err := Models.UpdateUser(&user, account)
	if nil == err {
		c.Data["json"] = true
	} else {
		c.Data["json"] = false
	}
	c.ServeJSON()
}

 

func (c *UserController) GetOne() {
	account := c.Ctx.Input.Params()[":account"]
	user, err := Models.GetOneUser(account)
	if nil == err {
		c.Data["json"] = user
	}
	c.ServeJSON()
}
