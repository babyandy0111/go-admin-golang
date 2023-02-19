package models

import (
	"time"
)

type SysLoginLog struct {
	Model
	Username      string    `json:"username" gorm:"type:varchar(128);comment:使用者"`
	Status        string    `json:"status" gorm:"type:varchar(4);comment:狀態"`
	Ipaddr        string    `json:"ipaddr" gorm:"type:varchar(255);comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"type:varchar(255);comment:歸屬地"`
	Browser       string    `json:"browser" gorm:"type:varchar(255);comment:瀏覽器"`
	Os            string    `json:"os" gorm:"type:varchar(255);comment:系統"`
	Platform      string    `json:"platform" gorm:"type:varchar(255);comment:平台"`
	LoginTime     time.Time `json:"loginTime" gorm:"type:timestamp;comment:登入時間"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);comment:備註"`
	Msg           string    `json:"msg" gorm:"type:varchar(255);comment:訊息"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:建立時間"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最後更新時間"`
	ControlBy
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}
