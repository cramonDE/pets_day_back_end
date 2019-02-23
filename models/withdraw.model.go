package models

import (
	"github.com/astaxie/beego/orm"
)

type WithdrawRecord struct {
	Account   string `json:"account"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	WithdrawTime  string `json:"withdraw_time"`
	MoneyType string `json:"money_type"`
}

func GetWithdrawRecords(account string) (*[]WithdrawRecord, error) {
	o := orm.NewOrm()
	var withdrawRecord []WithdrawRecord
	_, err := o.Raw("SELECT * FROM withdrawRecord WHERE account = ?", account).QueryRows(&withdrawRecord)
	return &withdrawRecord, err
}

func InsertWithdrawRecord(withdrawRecord *WithdrawRecord) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("INSERT INTO withdrawRecord (account, name, amount, withdraw_time, money_type) VALUES (?, ?, ?, ?, ?);", withdrawRecord.Account, withdrawRecord.Name, withdrawRecord.Amount, withdrawRecord.WithdrawTime, withdrawRecord.MoneyType).Exec()
	insertId, _ := res.LastInsertId()
	if (nil != err) {
		return insertId, err
	}
	_, err = o.Raw("UPDATE balance SET amount = amount - ? WHERE account = ?;", withdrawRecord.Amount, withdrawRecord.Account).Exec()
	return insertId, err
}
