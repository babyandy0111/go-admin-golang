package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ArticleGetPageReq struct {
	dto.Pagination `search:"-"`
	Status         string `form:"status"  search:"type:exact;column:status;table:article" comment:"狀態"`
	ArticleOrder
}

type ArticleOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:article"`
	Title     string `form:"titleOrder"  search:"type:order;column:title;table:article"`
	Author    string `form:"authorOrder"  search:"type:order;column:author;table:article"`
	Content   string `form:"contentOrder"  search:"type:order;column:content;table:article"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:article"`
	PublishAt string `form:"publishAtOrder"  search:"type:order;column:publish_at;table:article"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:article"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:article"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:article"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:article"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:article"`
}

func (m *ArticleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ArticleInsertReq struct {
	Id        int       `json:"-" comment:"流水號"` // 流水號
	Title     string    `json:"title" comment:"標題"`
	Author    string    `json:"author" comment:"作者"`
	Content   string    `json:"content" comment:"内容"`
	Status    string    `json:"status" comment:"狀態"`
	PublishAt time.Time `json:"publishAt" comment:"發佈時間"`
	common.ControlBy
}

func (s *ArticleInsertReq) Generate(model *models.Article) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Author = s.Author
	model.Content = s.Content
	model.Status = s.Status
	model.PublishAt = s.PublishAt
	model.CreateBy = s.CreateBy // 需要紀錄資料是被誰建立的
}

func (s *ArticleInsertReq) GetId() interface{} {
	return s.Id
}

type ArticleUpdateReq struct {
	Id        int       `uri:"id" comment:"流水號"` // 流水號
	Title     string    `json:"title" comment:"標題"`
	Author    string    `json:"author" comment:"作者"`
	Content   string    `json:"content" comment:"内容"`
	Status    string    `json:"status" comment:"狀態"`
	PublishAt time.Time `json:"publishAt" comment:"發佈時間"`
	common.ControlBy
}

func (s *ArticleUpdateReq) Generate(model *models.Article) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Author = s.Author
	model.Content = s.Content
	model.Status = s.Status
	model.PublishAt = s.PublishAt
	model.UpdateBy = s.UpdateBy // 需要紀錄資料是被誰更新的
}

func (s *ArticleUpdateReq) GetId() interface{} {
	return s.Id
}

// ArticleGetReq 查詢參數
type ArticleGetReq struct {
	Id int `uri:"id"`
}

func (s *ArticleGetReq) GetId() interface{} {
	return s.Id
}

// ArticleDeleteReq 刪除參數
type ArticleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ArticleDeleteReq) GetId() interface{} {
	return s.Ids
}
