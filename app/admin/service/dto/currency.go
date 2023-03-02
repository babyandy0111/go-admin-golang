package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CurrencyGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:currency" comment:"幣別名稱"`
	CurrencyOrder
}

type CurrencyOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:currency"`
	Currency  string `form:"currencyOrder"  search:"type:order;column:currency;table:currency"`
	Label     string `form:"labelOrder"  search:"type:order;column:label;table:currency"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:currency"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:currency"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:currency"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:currency"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:currency"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:currency"`
}

func (m *CurrencyGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CurrencyInsertReq struct {
	Id       int    `json:"-" comment:"幣別流水號"` // 幣別流水號
	Currency string `json:"currency" comment:"幣別代號"`
	Label    string `json:"label" comment:"幣別顯示名稱"`
	Name     string `json:"name" comment:"幣別名稱"`
	common.ControlBy
}

func (s *CurrencyInsertReq) Generate(model *models.Currency) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Currency = s.Currency
	model.Label = s.Label
	model.Name = s.Name
	model.CreateBy = s.CreateBy // 需要紀錄資料是被誰建立的
}

func (s *CurrencyInsertReq) GetId() interface{} {
	return s.Id
}

type CurrencyUpdateReq struct {
	Id       int    `uri:"id" comment:"幣別流水號"` // 幣別流水號
	Currency string `json:"currency" comment:"幣別代號"`
	Label    string `json:"label" comment:"幣別顯示名稱"`
	Name     string `json:"name" comment:"幣別名稱"`
	common.ControlBy
}

func (s *CurrencyUpdateReq) Generate(model *models.Currency) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Currency = s.Currency
	model.Label = s.Label
	model.Name = s.Name
	model.UpdateBy = s.UpdateBy // 需要紀錄資料是被誰更新的
}

func (s *CurrencyUpdateReq) GetId() interface{} {
	return s.Id
}

// CurrencyGetReq 查詢參數
type CurrencyGetReq struct {
	Id int `uri:"id"`
}

func (s *CurrencyGetReq) GetId() interface{} {
	return s.Id
}

// CurrencyDeleteReq 刪除參數
type CurrencyDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CurrencyDeleteReq) GetId() interface{} {
	return s.Ids
}
