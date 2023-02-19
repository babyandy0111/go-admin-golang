package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt" gorm:"comment:建立時間"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"comment:更新時間"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"comment:刪除時間"`
}
