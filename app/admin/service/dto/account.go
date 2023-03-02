package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AccountGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:account" comment:"客戶名稱"`
	AccountOrder
}

type AccountOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:account"`
	SfId          string `form:"sfIdOrder"  search:"type:order;column:sf_id;table:account"`
	AccountNumber string `form:"accountNumberOrder"  search:"type:order;column:account_number;table:account"`
	Name          string `form:"nameOrder"  search:"type:order;column:name;table:account"`
	EnName        string `form:"enNameOrder"  search:"type:order;column:en_name;table:account"`
	AccountSource string `form:"accountSourceOrder"  search:"type:order;column:account_source;table:account"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:account"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:account"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:account"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:account"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:account"`
}

func (m *AccountGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AccountInsertReq struct {
	Id            int    `json:"-" comment:"客戶流水號"` // 客戶流水號
	SfId          string `json:"sfId" comment:"sf key"`
	AccountNumber string `json:"accountNumber" comment:"客戶編號"`
	Name          string `json:"name" comment:"客戶名稱"`
	EnName        string `json:"enName" comment:"客戶英文名稱"`
	AccountSource string `json:"accountSource" comment:"客戶來源"`
	common.ControlBy
}

func (s *AccountInsertReq) Generate(model *models.Account) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SfId = s.SfId
	model.AccountNumber = s.AccountNumber
	model.Name = s.Name
	model.EnName = s.EnName
	model.AccountSource = s.AccountSource
	model.CreateBy = s.CreateBy // 需要紀錄資料是被誰建立的
}

func (s *AccountInsertReq) GetId() interface{} {
	return s.Id
}

type AccountUpdateReq struct {
	Id            int    `uri:"id" comment:"客戶流水號"` // 客戶流水號
	SfId          string `json:"sfId" comment:"sf key"`
	AccountNumber string `json:"accountNumber" comment:"客戶編號"`
	Name          string `json:"name" comment:"客戶名稱"`
	EnName        string `json:"enName" comment:"客戶英文名稱"`
	AccountSource string `json:"accountSource" comment:"客戶來源"`
	common.ControlBy
}

func (s *AccountUpdateReq) Generate(model *models.Account) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SfId = s.SfId
	model.AccountNumber = s.AccountNumber
	model.Name = s.Name
	model.EnName = s.EnName
	model.AccountSource = s.AccountSource
	model.UpdateBy = s.UpdateBy // 需要紀錄資料是被誰更新的
}

func (s *AccountUpdateReq) GetId() interface{} {
	return s.Id
}

// AccountGetReq 查詢參數
type AccountGetReq struct {
	Id int `uri:"id"`
}

func (s *AccountGetReq) GetId() interface{} {
	return s.Id
}

// AccountDeleteReq 刪除參數
type AccountDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AccountDeleteReq) GetId() interface{} {
	return s.Ids
}
