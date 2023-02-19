package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:流水號"  json:"userId"`
	Username string `json:"username" gorm:"type:varchar(64);comment:使用者"`
	Password string `json:"-" gorm:"type:varchar(128);comment:密碼"`
	NickName string `json:"nickName" gorm:"type:varchar(128);comment:暱稱"`
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:手機號碼"`
	RoleId   int    `json:"roleId" gorm:"type:bigint;comment:角色ID"`
	Salt     string `json:"-" gorm:"type:varchar(255);comment:slat"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:頭貼"`
	Sex      string `json:"sex" gorm:"type:varchar(255);comment:性别"`
	Email    string `json:"email" gorm:"type:varchar(128);comment:信箱"`
	DeptId   int    `json:"deptId" gorm:"type:bigint;comment:部門"`
	PostId   int    `json:"postId" gorm:"type:bigint;comment:職稱"`
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:備註"`
	Status   string `json:"status" gorm:"type:varchar(4);comment:狀態"`
	ControlBy
	ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

// Encrypt 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}
