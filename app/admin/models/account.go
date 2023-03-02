package models

import (
	"go-admin/common/models"
)

type Account struct {
	models.Model

	SfId          string `json:"sfId" gorm:"type:varchar(50);comment:sf key"`
	AccountNumber string `json:"accountNumber" gorm:"type:varchar(50);comment:客戶編號"`
	Name          string `json:"name" gorm:"type:varchar(100);comment:客戶名稱"`
	EnName        string `json:"enName" gorm:"type:varchar(100);comment:客戶英文名稱"`
	AccountSource string `json:"accountSource" gorm:"type:varchar(100);comment:客戶來源"`
	models.ModelTime
	models.ControlBy
}

func (Account) TableName() string {
	return "account"
}

func (e *Account) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Account) GetId() interface{} {
	return e.Id
}
