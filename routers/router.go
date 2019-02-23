package routers

import (
	"pets-day/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/?:account", &controllers.UserController{}, "get:GetOne")
	beego.Router("/user/?:account", &controllers.UserController{})
	beego.Router("/balance/?:account", &controllers.BalanceController{})
	beego.Router("/saving-record/?:account", &controllers.SavingRecordController{})
	beego.Router("/withdraw-record/?:account", &controllers.WithdrawRecordController{})
}
