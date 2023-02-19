package handler

import "go-admin/common/models"

type SysRole struct {
	RoleId    int    `json:"roleId" gorm:"primaryKey;autoIncrement"` // 角色流水號
	RoleName  string `json:"roleName" gorm:"size:128;"`              // 角色名稱
	Status    string `json:"status" gorm:"size:4;"`                  //
	RoleKey   string `json:"roleKey" gorm:"size:128;"`               //角色代碼
	RoleSort  int    `json:"roleSort" gorm:""`                       //角色排序
	Flag      string `json:"flag" gorm:"size:128;"`                  //
	Remark    string `json:"remark" gorm:"size:255;"`                //備註
	Admin     bool   `json:"admin" gorm:"size:4;"`
	DataScope string `json:"dataScope" gorm:"size:128;"`
	Params    string `json:"params" gorm:"-"`
	MenuIds   []int  `json:"menuIds" gorm:"-"`
	DeptIds   []int  `json:"deptIds" gorm:"-"`
	models.ControlBy
	models.ModelTime
}

func (SysRole) TableName() string {
	return "sys_role"
}
