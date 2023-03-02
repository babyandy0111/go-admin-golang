package models

import (
	"go-admin/common/models"
)

type Quote struct {
	models.Model

	SfId          string `json:"sfId" gorm:"type:varchar(50);comment:sf key"`
	QuoteNumber   string `json:"quoteNumber" gorm:"type:varchar(25);comment:報價單號"`
	CompanyTypeId string `json:"companyTypeId" gorm:"type:int(11);comment:公司流水號"`
	AccountId     string `json:"accountId" gorm:"type:int(11);comment:客戶"`
	CurrencyId    string `json:"currencyId" gorm:"type:int(11);comment:幣別"`
	QuoteName     string `json:"quoteName" gorm:"type:varchar(255);comment:報價名稱"`
	ProjectName   string `json:"projectName" gorm:"type:varchar(255);comment:專案名稱"`
	QuoteNote     string `json:"quoteNote" gorm:"type:text;comment:備註"`
	Stage         string `json:"stage" gorm:"type:varchar(25);comment:階段"`
	Status        string `json:"status" gorm:"type:varchar(25);comment:狀態"`
	Total         string `json:"total" gorm:"type:decimal(11,2);comment:報價單總額"`
	Tax           string `json:"tax" gorm:"type:decimal(11,2);comment:税"`
	GrandTotal    string `json:"grandTotal" gorm:"type:decimal(11,2);comment:含稅總額"`
	Discount      string `json:"discount" gorm:"type:decimal(11,2);comment:折扣"`
	models.ModelTime
	models.ControlBy
}

func (Quote) TableName() string {
	return "quote"
}

func (e *Quote) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Quote) GetId() interface{} {
	return e.Id
}
