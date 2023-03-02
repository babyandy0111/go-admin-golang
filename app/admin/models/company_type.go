package models

import (
	"go-admin/common/models"
)

type CompanyType struct {
	models.Model

	Code        string `json:"code" gorm:"type:varchar(10);comment:本司代碼"`
	Name        string `json:"name" gorm:"type:varchar(50);comment:本司名稱"`
	Description string `json:"description" gorm:"type:text;comment:該代碼基本描述"`
	models.ModelTime
	models.ControlBy
}

func (CompanyType) TableName() string {
	return "company_type"
}

func (e *CompanyType) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *CompanyType) GetId() interface{} {
	return e.Id
}
