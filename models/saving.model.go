package models

import (
	"github.com/astaxie/beego/orm"
)

type SavingRecord struct {
	Account   string `json:"account"`
	Name      string `json:"name"`
	Bank      string `json:"bank"`
	Amount    int    `json:"amount"`
	SaveTime  string `json:"save_time"`
	MoneyType string `json:"money_type"`
}

func GetSavingRecords(account string) (*[]SavingRecord, error) {
	o := orm.NewOrm()
	var savingRecords []SavingRecord
	_, err := o.Raw("SELECT * FROM savingRecord WHERE account = ?", account).QueryRows(&savingRecords)
	return &savingRecords, err
}

func InsertSavingRecord(savingRecord *SavingRecord) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("INSERT INTO savingRecord (account, name, bank, amount, save_time, money_type) VALUES (?, ?, ?, ?, ?, ?);", savingRecord.Account, savingRecord.Name, savingRecord.Bank, savingRecord.Amount, savingRecord.SaveTime, savingRecord.MoneyType).Exec()
	insertId, _ := res.LastInsertId()
	if (nil != err) {
		return insertId, err
	}
	_, err = o.Raw("UPDATE balance SET amount = amount + ? WHERE account = ?;", savingRecord.Amount, savingRecord.Account).Exec()
	return insertId, err
}
