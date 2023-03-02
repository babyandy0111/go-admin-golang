package models

import (
	"time"

	"go-admin/common/models"
)

type Article struct {
	models.Model

	Title     string    `json:"title" gorm:"type:varchar(128);comment:標題"`
	Author    string    `json:"author" gorm:"type:varchar(128);comment:作者"`
	Content   string    `json:"content" gorm:"type:varchar(255);comment:内容"`
	Status    string    `json:"status" gorm:"type:int(1);comment:狀態"`
	PublishAt time.Time `json:"publishAt" gorm:"type:timestamp;comment:發佈時間"`
	models.ModelTime
	models.ControlBy
}

func (Article) TableName() string {
	return "article"
}

func (e *Article) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Article) GetId() interface{} {
	return e.Id
}
