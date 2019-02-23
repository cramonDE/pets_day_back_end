package models

import (
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/orm"
)

func CreateDatabase() {
	o := orm.NewOrm()
	file, err := ioutil.ReadFile("static/create.sql")
	if nil != err {
		return
	}
	tables := strings.Split(string(file), ";")
	for _, table := range tables {
		o.Raw(table).Exec()
	}
}

func CreateInitialData() {
	o := orm.NewOrm()
	
	o.Raw("INSERT INTO `user` (account, name, city, address, phoneCall, mobile, email, job, contact, contactPhone, income, creditRating) VALUES ('440583199606254832' ,'小李', '广州', '广州市番禺区广东工业大学西8-227' ,'020-39332232', '13299492932', '545032132@qq.com', '工程师' , '小王', '18922449523','20000', '较好');").Exec()
}
