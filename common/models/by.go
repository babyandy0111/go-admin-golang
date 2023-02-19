package models

import (
	"time"

	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主健流水號"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:建立時間"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新時間"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:刪除時間"`
}
