package models

type SysRole struct {
	RoleId    int       `json:"roleId" gorm:"primaryKey;autoIncrement;comment:角色流水號"`
	RoleName  string    `json:"roleName" gorm:"size:128;comment:角色名稱"`
	Status    string    `json:"status" gorm:"size:4;"`
	RoleKey   string    `json:"roleKey" gorm:"size:128;comment:角色代碼"`
	RoleSort  int       `json:"roleSort" gorm:"comment:角色排序"`
	Flag      string    `json:"flag" gorm:"size:128;"`
	Remark    string    `json:"remark" gorm:"size:255;comment:備註"`
	Admin     bool      `json:"admin" gorm:"size:4;"`
	DataScope string    `json:"dataScope" gorm:"size:128;"`
	SysMenu   []SysMenu `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:RoleId;joinForeignKey:role_id;references:MenuId;joinReferences:menu_id;"`
	ControlBy
	ModelTime
}

func (SysRole) TableName() string {
	return "sys_role"
}
