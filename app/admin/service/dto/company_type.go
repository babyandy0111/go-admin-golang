package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CompanyTypeGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:company_type" comment:"本司名稱"`
	CompanyTypeOrder
}

type CompanyTypeOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:company_type"`
	Code        string `form:"codeOrder"  search:"type:order;column:code;table:company_type"`
	Name        string `form:"nameOrder"  search:"type:order;column:name;table:company_type"`
	Description string `form:"descriptionOrder"  search:"type:order;column:description;table:company_type"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:company_type"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:company_type"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:company_type"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:company_type"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:company_type"`
}

func (m *CompanyTypeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CompanyTypeInsertReq struct {
	Id          int    `json:"-" comment:"本司流水號"` // 本司流水號
	Code        string `json:"code" comment:"本司代碼"`
	Name        string `json:"name" comment:"本司名稱"`
	Description string `json:"description" comment:"該代碼基本描述"`
	common.ControlBy
}

func (s *CompanyTypeInsertReq) Generate(model *models.CompanyType) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Code = s.Code
	model.Name = s.Name
	model.Description = s.Description
	model.CreateBy = s.CreateBy // 需要紀錄資料是被誰建立的
}

func (s *CompanyTypeInsertReq) GetId() interface{} {
	return s.Id
}

type CompanyTypeUpdateReq struct {
	Id          int    `uri:"id" comment:"本司流水號"` // 本司流水號
	Code        string `json:"code" comment:"本司代碼"`
	Name        string `json:"name" comment:"本司名稱"`
	Description string `json:"description" comment:"該代碼基本描述"`
	common.ControlBy
}

func (s *CompanyTypeUpdateReq) Generate(model *models.CompanyType) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Code = s.Code
	model.Name = s.Name
	model.Description = s.Description
	model.UpdateBy = s.UpdateBy // 需要紀錄資料是被誰更新的
}

func (s *CompanyTypeUpdateReq) GetId() interface{} {
	return s.Id
}

// CompanyTypeGetReq 查詢參數
type CompanyTypeGetReq struct {
	Id int `uri:"id"`
}

func (s *CompanyTypeGetReq) GetId() interface{} {
	return s.Id
}

// CompanyTypeDeleteReq 刪除參數
type CompanyTypeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CompanyTypeDeleteReq) GetId() interface{} {
	return s.Ids
}
