package models

type SysApi struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement;comment:流水號"`
	Handle string `json:"handle" gorm:"size:128;comment:handle"`
	Title  string `json:"title" gorm:"size:128;comment:標题"`
	Path   string `json:"path" gorm:"size:128;comment:地址"`
	Type   string `json:"type" gorm:"size:16;comment:接口類型"`
	Action string `json:"action" gorm:"size:16;comment:請求類型"`
	ModelTime
	ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}
