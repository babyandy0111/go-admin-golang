package models

import (
	"encoding/json"
	"errors"
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/storage"

	"go-admin/common/models"
)

type SysOperaLog struct {
	models.Model
	Title         string    `json:"title" gorm:"size:255;comment:操作模組"`
	BusinessType  string    `json:"businessType" gorm:"size:128;comment:操作類型"`
	BusinessTypes string    `json:"businessTypes" gorm:"size:128;comment:BusinessTypes"`
	Method        string    `json:"method" gorm:"size:128;comment:函數"`
	RequestMethod string    `json:"requestMethod" gorm:"size:128;comment:請求方式 GET POST PUT DELETE"`
	OperatorType  string    `json:"operatorType" gorm:"size:128;comment:操作類型"`
	OperName      string    `json:"operName" gorm:"size:128;comment:操作者"`
	DeptName      string    `json:"deptName" gorm:"size:128;comment:部門名稱"`
	OperUrl       string    `json:"operUrl" gorm:"size:255;comment:訪問位置"`
	OperIp        string    `json:"operIp" gorm:"size:128;comment:client ip"`
	OperLocation  string    `json:"operLocation" gorm:"size:128;comment:訪問位置"`
	OperParam     string    `json:"operParam" gorm:"text;comment:請求參數"`
	Status        string    `json:"status" gorm:"size:4;comment:操作狀態 1:正常 2:關閉"`
	OperTime      time.Time `json:"operTime" gorm:"comment:操作時間"`
	JsonResult    string    `json:"jsonResult" gorm:"size:255;comment:response"`
	Remark        string    `json:"remark" gorm:"size:255;comment:備註"`
	LatencyTime   string    `json:"latencyTime" gorm:"size:128;comment:response time"`
	UserAgent     string    `json:"userAgent" gorm:"size:255;comment:ua"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:建立時間"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最後更新時間"`
	models.ControlBy
}

func (SysOperaLog) TableName() string {
	return "sys_opera_log"
}

func (e *SysOperaLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysOperaLog) GetId() interface{} {
	return e.Id
}

// SaveOperaLog 从队列中取得操作Log
func SaveOperaLog(message storage.Messager) (err error) {
	//准备db
	db := sdk.Runtime.GetDbByKey(message.GetPrefix())
	if db == nil {
		err = errors.New("db not exist")
		log.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		log.Errorf("json Marshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var l SysOperaLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	// 超出100个字符返回值截断
	if len(l.JsonResult) > 100 {
		l.JsonResult = l.JsonResult[:100]
	}
	err = db.Create(&l).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	return nil
}
