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

type SysLoginLog struct {
	models.Model
	Username      string    `json:"username" gorm:"size:128;comment:使用者"`
	Status        string    `json:"status" gorm:"size:4;comment:狀態"`
	Ipaddr        string    `json:"ipaddr" gorm:"size:255;comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"size:255;comment:歸屬地"`
	Browser       string    `json:"browser" gorm:"size:255;comment:瀏覽器"`
	Os            string    `json:"os" gorm:"size:255;comment:系統"`
	Platform      string    `json:"platform" gorm:"size:255;comment:平台"`
	LoginTime     time.Time `json:"loginTime" gorm:"comment:登入時間"`
	Remark        string    `json:"remark" gorm:"size:255;comment:備註"`
	Msg           string    `json:"msg" gorm:"size:255;comment:訊息"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:建立時間"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最後更新時間"`
	models.ControlBy
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}

func (e *SysLoginLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysLoginLog) GetId() interface{} {
	return e.Id
}

// SaveLoginLog 从队列中取得登入Log
func SaveLoginLog(message storage.Messager) (err error) {
	//准备db
	db := sdk.Runtime.GetDbByKey(message.GetPrefix())
	if db == nil {
		err = errors.New("db not exist")
		log.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		return err
	}
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		log.Errorf("json Marshal error, %s", err.Error())
		return err
	}
	var l SysLoginLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		return err
	}
	err = db.Create(&l).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		return err
	}
	return nil
}
