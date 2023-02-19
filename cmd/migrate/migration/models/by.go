package models

import (
	"time"

	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:建立者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:流水號"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:建立時間"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最後更新時間"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:刪除時間"`
}
