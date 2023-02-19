package models

import (
	"time"
)

type SysOperaLog struct {
	Model
	Title         string    `json:"title" gorm:"type:varchar(255);comment:操作模組"`
	BusinessType  string    `json:"businessType" gorm:"type:varchar(128);comment:操作類型"`
	BusinessTypes string    `json:"businessTypes" gorm:"type:varchar(128);comment:BusinessTypes"`
	Method        string    `json:"method" gorm:"type:varchar(128);comment:函數"`
	RequestMethod string    `json:"requestMethod" gorm:"type:varchar(128);comment:請求方式: GET POST PUT DELETE"`
	OperatorType  string    `json:"operatorType" gorm:"type:varchar(128);comment:操作類型"`
	OperName      string    `json:"operName" gorm:"type:varchar(128);comment:操作者"`
	DeptName      string    `json:"deptName" gorm:"type:varchar(128);comment:部門名稱"`
	OperUrl       string    `json:"operUrl" gorm:"type:varchar(255);comment:訪問位置"`
	OperIp        string    `json:"operIp" gorm:"type:varchar(128);comment:client ip"`
	OperLocation  string    `json:"operLocation" gorm:"type:varchar(128);comment:訪問位置"`
	OperParam     string    `json:"operParam" gorm:"type:text;comment:請求參數"`
	Status        string    `json:"status" gorm:"type:varchar(4);comment:操作狀態 1:正常 2:關閉"`
	OperTime      time.Time `json:"operTime" gorm:"type:timestamp;comment:操作時間"`
	JsonResult    string    `json:"jsonResult" gorm:"type:varchar(255);comment:response"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);comment:備註"`
	LatencyTime   string    `json:"latencyTime" gorm:"type:varchar(128);comment:response time"`
	UserAgent     string    `json:"userAgent" gorm:"type:varchar(255);comment:ua"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:建立時間"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最後更新時間"`
	ControlBy
}

func (SysOperaLog) TableName() string {
	return "sys_opera_log"
}
