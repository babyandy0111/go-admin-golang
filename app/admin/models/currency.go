package models

import (
	"go-admin/common/models"
)

type Currency struct {
	models.Model

	Currency string `json:"currency" gorm:"type:varchar(10);comment:幣別代號"`
	Label    string `json:"label" gorm:"type:varchar(50);comment:幣別顯示名稱"`
	Name     string `json:"name" gorm:"type:varchar(50);comment:幣別名稱"`
	models.ModelTime
	models.ControlBy
}

func (Currency) TableName() string {
	return "currency"
}

func (e *Currency) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Currency) GetId() interface{} {
	return e.Id
}
