package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type QuoteGetPageReq struct {
	dto.Pagination `search:"-"`
	QuoteNumber    string `form:"quoteNumber"  search:"type:exact;column:quote_number;table:quote" comment:"報價單號"`
	CompanyTypeId  string `form:"companyTypeId"  search:"type:exact;column:company_type_id;table:quote" comment:"公司流水號"`
	AccountId      string `form:"accountId"  search:"type:exact;column:account_id;table:quote" comment:"客戶"`
	QuoteOrder
}

type QuoteOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:quote"`
	SfId          string `form:"sfIdOrder"  search:"type:order;column:sf_id;table:quote"`
	QuoteNumber   string `form:"quoteNumberOrder"  search:"type:order;column:quote_number;table:quote"`
	CompanyTypeId string `form:"companyTypeIdOrder"  search:"type:order;column:company_type_id;table:quote"`
	AccountId     string `form:"accountIdOrder"  search:"type:order;column:account_id;table:quote"`
	CurrencyId    string `form:"currencyIdOrder"  search:"type:order;column:currency_id;table:quote"`
	QuoteName     string `form:"quoteNameOrder"  search:"type:order;column:quote_name;table:quote"`
	ProjectName   string `form:"projectNameOrder"  search:"type:order;column:project_name;table:quote"`
	QuoteNote     string `form:"quoteNoteOrder"  search:"type:order;column:quote_note;table:quote"`
	Stage         string `form:"stageOrder"  search:"type:order;column:stage;table:quote"`
	Status        string `form:"statusOrder"  search:"type:order;column:status;table:quote"`
	Total         string `form:"totalOrder"  search:"type:order;column:total;table:quote"`
	Tax           string `form:"taxOrder"  search:"type:order;column:tax;table:quote"`
	GrandTotal    string `form:"grandTotalOrder"  search:"type:order;column:grand_total;table:quote"`
	Discount      string `form:"discountOrder"  search:"type:order;column:discount;table:quote"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:quote"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:quote"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:quote"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:quote"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:quote"`
}

func (m *QuoteGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type QuoteInsertReq struct {
	Id            int    `json:"-" comment:"代號"` // 代號
	SfId          string `json:"sfId" comment:"sf key"`
	QuoteNumber   string `json:"quoteNumber" comment:"報價單號"`
	CompanyTypeId string `json:"companyTypeId" comment:"公司流水號"`
	AccountId     string `json:"accountId" comment:"客戶"`
	CurrencyId    string `json:"currencyId" comment:"幣別"`
	QuoteName     string `json:"quoteName" comment:"報價名稱"`
	ProjectName   string `json:"projectName" comment:"專案名稱"`
	QuoteNote     string `json:"quoteNote" comment:"備註"`
	Stage         string `json:"stage" comment:"階段"`
	Status        string `json:"status" comment:"狀態"`
	Total         string `json:"total" comment:"報價單總額"`
	Tax           string `json:"tax" comment:"税"`
	GrandTotal    string `json:"grandTotal" comment:"含稅總額"`
	Discount      string `json:"discount" comment:"折扣"`
	common.ControlBy
}

func (s *QuoteInsertReq) Generate(model *models.Quote) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SfId = s.SfId
	model.QuoteNumber = s.QuoteNumber
	model.CompanyTypeId = s.CompanyTypeId
	model.AccountId = s.AccountId
	model.CurrencyId = s.CurrencyId
	model.QuoteName = s.QuoteName
	model.ProjectName = s.ProjectName
	model.QuoteNote = s.QuoteNote
	model.Stage = s.Stage
	model.Status = s.Status
	model.Total = s.Total
	model.Tax = s.Tax
	model.GrandTotal = s.GrandTotal
	model.Discount = s.Discount
	model.CreateBy = s.CreateBy // 需要紀錄資料是被誰建立的
}

func (s *QuoteInsertReq) GetId() interface{} {
	return s.Id
}

type QuoteUpdateReq struct {
	Id            int    `uri:"id" comment:"代號"` // 代號
	SfId          string `json:"sfId" comment:"sf key"`
	QuoteNumber   string `json:"quoteNumber" comment:"報價單號"`
	CompanyTypeId string `json:"companyTypeId" comment:"公司流水號"`
	AccountId     string `json:"accountId" comment:"客戶"`
	CurrencyId    string `json:"currencyId" comment:"幣別"`
	QuoteName     string `json:"quoteName" comment:"報價名稱"`
	ProjectName   string `json:"projectName" comment:"專案名稱"`
	QuoteNote     string `json:"quoteNote" comment:"備註"`
	Stage         string `json:"stage" comment:"階段"`
	Status        string `json:"status" comment:"狀態"`
	Total         string `json:"total" comment:"報價單總額"`
	Tax           string `json:"tax" comment:"税"`
	GrandTotal    string `json:"grandTotal" comment:"含稅總額"`
	Discount      string `json:"discount" comment:"折扣"`
	common.ControlBy
}

func (s *QuoteUpdateReq) Generate(model *models.Quote) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SfId = s.SfId
	model.QuoteNumber = s.QuoteNumber
	model.CompanyTypeId = s.CompanyTypeId
	model.AccountId = s.AccountId
	model.CurrencyId = s.CurrencyId
	model.QuoteName = s.QuoteName
	model.ProjectName = s.ProjectName
	model.QuoteNote = s.QuoteNote
	model.Stage = s.Stage
	model.Status = s.Status
	model.Total = s.Total
	model.Tax = s.Tax
	model.GrandTotal = s.GrandTotal
	model.Discount = s.Discount
	model.UpdateBy = s.UpdateBy // 需要紀錄資料是被誰更新的
}

func (s *QuoteUpdateReq) GetId() interface{} {
	return s.Id
}

// QuoteGetReq 查詢參數
type QuoteGetReq struct {
	Id int `uri:"id"`
}

func (s *QuoteGetReq) GetId() interface{} {
	return s.Id
}

// QuoteDeleteReq 刪除參數
type QuoteDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *QuoteDeleteReq) GetId() interface{} {
	return s.Ids
}
