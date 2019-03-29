package controllers

import (
	"github.com/astaxie/beego"
)

type UploadHandlerController struct {
	beego.Controller
}

func (c *UploadHandlerController) Post() {
	f, fh, err := c.GetFile("emoji")
	defer f.Close()
	if err != nil {
		c.Data["json"] = &map[string]interface{}{"path": "", "succ": false}
		c.ServeJSON()
	} else {
		c.SaveToFile("emoji", "static/img/"+fh.Filename)
		c.Data["json"] = &map[string]interface{}{"path": "/static/img/" + fh.Filename, "succ": true}
		c.ServeJSON()
	}
}
