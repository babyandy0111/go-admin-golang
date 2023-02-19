package models

type DictType struct {
	DictId   int    `gorm:"primaryKey;autoIncrement;comment:流水號" json:"dictId"`
	DictName string `gorm:"size:128;comment:字典名稱" json:"dictName"`
	DictType string `gorm:"size:128;comment:字典類型" json:"dictType"`
	Status   int    `gorm:"size:4;comment:狀態" json:"status"`
	Remark   string `gorm:"size:255;comment:備註" json:"remark"`
	ControlBy
	ModelTime
}

func (DictType) TableName() string {
	return "sys_dict_type"
}
