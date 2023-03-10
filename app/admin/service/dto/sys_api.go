package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// SysApiGetPageReq 功能列表請求參數
type SysApiGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"  search:"type:contains;column:title;table:sys_api" comment:"標题"`
	Path           string `form:"path"  search:"type:contains;column:path;table:sys_api" comment:"地址"`
	Action         string `form:"action"  search:"type:exact;column:action;table:sys_api" comment:"類型"`
	ParentId       string `form:"parentId"  search:"type:exact;column:parent_id;table:sys_api" comment:"按鈕id"`
	SysApiOrder
}

type SysApiOrder struct {
	TitleOrder     string `search:"type:order;column:title;table:sys_api" form:"titleOrder"`
	PathOrder      string `search:"type:order;column:path;table:sys_api" form:"pathOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_api" form:"createdAtOrder"`
}

func (m *SysApiGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysApiInsertReq 功能建立請求參數
type SysApiInsertReq struct {
	Id     int    `json:"-" comment:"流水號"` // 流水號
	Handle string `json:"handle" comment:"handle"`
	Title  string `json:"title" comment:"標题"`
	Path   string `json:"path" comment:"地址"`
	Type   string `json:"type" comment:""`
	Action string `json:"action" comment:"類型"`
	common.ControlBy
}

func (s *SysApiInsertReq) Generate(model *models.SysApi) {
	model.Handle = s.Handle
	model.Title = s.Title
	model.Path = s.Path
	model.Type = s.Type
	model.Action = s.Action
}

func (s *SysApiInsertReq) GetId() interface{} {
	return s.Id
}

// SysApiUpdateReq 功能更新請求參數
type SysApiUpdateReq struct {
	Id     int    `uri:"id" comment:"流水號"` // 流水號
	Handle string `json:"handle" comment:"handle"`
	Title  string `json:"title" comment:"標题"`
	Path   string `json:"path" comment:"地址"`
	Type   string `json:"type" comment:""`
	Action string `json:"action" comment:"類型"`
	common.ControlBy
}

func (s *SysApiUpdateReq) Generate(model *models.SysApi) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Handle = s.Handle
	model.Title = s.Title
	model.Path = s.Path
	model.Type = s.Type
	model.Action = s.Action
}

func (s *SysApiUpdateReq) GetId() interface{} {
	return s.Id
}

// SysApiGetReq 功能取得請求參數
type SysApiGetReq struct {
	Id int `uri:"id"`
}

func (s *SysApiGetReq) GetId() interface{} {
	return s.Id
}

// SysApiDeleteReq 功能刪除請求參數
type SysApiDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysApiDeleteReq) GetId() interface{} {
	return s.Ids
}
