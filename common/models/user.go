package models

import (
	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
)

// BaseUser 密碼登入基礎User
type BaseUser struct {
	Username     string `json:"username" gorm:"type:varchar(100);comment:使用者"`
	Salt         string `json:"-" gorm:"type:varchar(255);comment:slat;<-"`
	PasswordHash string `json:"-" gorm:"type:varchar(128);comment:密碼hash;<-"`
	Password     string `json:"password" gorm:"-"`
}

// SetPassword 設定密碼
func (u *BaseUser) SetPassword(value string) {
	u.Password = value
	u.generateSalt()
	u.PasswordHash = u.GetPasswordHash()
}

// GetPasswordHash 取得密碼hash
func (u *BaseUser) GetPasswordHash() string {
	passwordHash, err := pkg.SetPassword(u.Password, u.Salt)
	if err != nil {
		return ""
	}
	return passwordHash
}

// generateSalt 生成slat值
func (u *BaseUser) generateSalt() {
	u.Salt = pkg.GenerateRandomKey16()
}

// Verify 驗證密碼
func (u *BaseUser) Verify(db *gorm.DB, tableName string) bool {
	db.Table(tableName).Where("username = ?", u.Username).First(u)
	return u.GetPasswordHash() == u.PasswordHash
}
