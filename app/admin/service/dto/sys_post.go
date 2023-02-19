package dto

import (
	"go-admin/app/admin/models"
	common "go-admin/common/models"

	"go-admin/common/dto"
)

// SysPostPageReq 列表或者搜索使用結構体
type SysPostPageReq struct {
	dto.Pagination `search:"-"`
	PostId         int    `form:"postId" search:"type:exact;column:post_id;table:sys_post" comment:"id"`         // id
	PostName       string `form:"postName" search:"type:contains;column:post_name;table:sys_post" comment:"名稱"`  // 名稱
	PostCode       string `form:"postCode" search:"type:contains;column:post_code;table:sys_post" comment:"流水號"` // 流水號
	Sort           int    `form:"sort" search:"type:exact;column:sort;table:sys_post" comment:"排序"`              // 排序
	Status         int    `form:"status" search:"type:exact;column:status;table:sys_post" comment:"狀態"`          // 狀態
	Remark         string `form:"remark" search:"type:exact;column:remark;table:sys_post" comment:"備註"`          // 備註
}

func (m *SysPostPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysPostInsertReq 增使用的結構体
type SysPostInsertReq struct {
	PostId   int    `uri:"id"  comment:"id"`
	PostName string `form:"postName"  comment:"名稱"`
	PostCode string `form:"postCode" comment:"流水號"`
	Sort     int    `form:"sort" comment:"排序"`
	Status   int    `form:"status"   comment:"狀態"`
	Remark   string `form:"remark"   comment:"備註"`
	common.ControlBy
}

func (s *SysPostInsertReq) Generate(model *models.SysPost) {
	model.PostName = s.PostName
	model.PostCode = s.PostCode
	model.Sort = s.Sort
	model.Status = s.Status
	model.Remark = s.Remark
	if s.ControlBy.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
	if s.ControlBy.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
}

// GetId 取得資料對應的ID
func (s *SysPostInsertReq) GetId() interface{} {
	return s.PostId
}

// SysPostUpdateReq 改使用的結構体
type SysPostUpdateReq struct {
	PostId   int    `uri:"id"  comment:"id"`
	PostName string `form:"postName"  comment:"名稱"`
	PostCode string `form:"postCode" comment:"流水號"`
	Sort     int    `form:"sort" comment:"排序"`
	Status   int    `form:"status"   comment:"狀態"`
	Remark   string `form:"remark"   comment:"備註"`
	common.ControlBy
}

func (s *SysPostUpdateReq) Generate(model *models.SysPost) {
	model.PostId = s.PostId
	model.PostName = s.PostName
	model.PostCode = s.PostCode
	model.Sort = s.Sort
	model.Status = s.Status
	model.Remark = s.Remark
	if s.ControlBy.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
	if s.ControlBy.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
}

func (s *SysPostUpdateReq) GetId() interface{} {
	return s.PostId
}

// SysPostGetReq 取得单个的結構体
type SysPostGetReq struct {
	Id int `uri:"id"`
}

func (s *SysPostGetReq) GetId() interface{} {
	return s.Id
}

// SysPostDeleteReq 刪除的結構体
type SysPostDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysPostDeleteReq) Generate(model *models.SysPost) {
	if s.ControlBy.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
	if s.ControlBy.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
}

func (s *SysPostDeleteReq) GetId() interface{} {
	return s.Ids
}
