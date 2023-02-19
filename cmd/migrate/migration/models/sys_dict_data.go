package models

type DictData struct {
	DictCode  int    `gorm:"primaryKey;autoIncrement;comment:流水號" json:"dictCode" example:"1"`
	DictSort  int    `gorm:"comment:顯示順序" json:"dictSort"`
	DictLabel string `gorm:"size:128;comment:label" json:"dictLabel"`
	DictValue string `gorm:"size:255;comment:key" json:"dictValue"`
	DictType  string `gorm:"size:64;comment:type" json:"dictType"`
	CssClass  string `gorm:"size:128;" json:"cssClass"`
	ListClass string `gorm:"size:128;" json:"listClass"`
	IsDefault string `gorm:"size:8;" json:"isDefault"`
	Status    int    `gorm:"size:4;comment:狀態" json:"status"`
	Default   string `gorm:"size:8;comment:預設" json:"default"`
	Remark    string `gorm:"size:255;comment:備註" json:"remark"`
	ControlBy
	ModelTime
}

func (DictData) TableName() string {
	return "sys_dict_data"
}
