package routers

import (
	"pets-day/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/upload-handler", &controllers.UploadHandlerController{})
}
