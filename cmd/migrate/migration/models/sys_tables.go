package models

type SysTables struct {
	TableId             int    `gorm:"primaryKey;autoIncrement;comment:表流水號" json:"tableId"`
	TBName              string `gorm:"column:table_name;size:255;comment:表名稱" json:"tableName"`
	TableComment        string `gorm:"size:255;comment:表備註" json:"tableComment"`
	ClassName           string `gorm:"size:255;comment:類名" json:"className"`
	TplCategory         string `gorm:"size:255;comment:包名" json:"tplCategory"`
	PackageName         string `gorm:"size:255;comment:" json:"packageName"`
	ModuleName          string `gorm:"size:255;comment:go文件名" json:"moduleName"`
	ModuleFrontName     string `gorm:"size:255;comment:前端文件名;" json:"moduleFrontName"`
	BusinessName        string `gorm:"size:255;comment:業務邏輯名稱" json:"businessName"`
	FunctionName        string `gorm:"size:255;comment:功能名稱" json:"functionName"`
	FunctionAuthor      string `gorm:"size:255;comment:功能作者" json:"functionAuthor"`
	PkColumn            string `gorm:"size:255;" json:"pkColumn"`
	PkGoField           string `gorm:"size:255;" json:"pkGoField"`
	PkJsonField         string `gorm:"size:255;" json:"pkJsonField"`
	Options             string `gorm:"size:255;" json:"options"`
	TreeCode            string `gorm:"size:255;" json:"treeCode"`
	TreeParentCode      string `gorm:"size:255;" json:"treeParentCode"`
	TreeName            string `gorm:"size:255;" json:"treeName"`
	Tree                bool   `gorm:"size:1;default:0;" json:"tree"`
	Crud                bool   `gorm:"size:1;default:1;" json:"crud"`
	Remark              string `gorm:"size:255;" json:"remark"`
	IsDataScope         int    `gorm:"size:1;" json:"isDataScope"`
	IsActions           int    `gorm:"size:1;" json:"isActions"`
	IsAuth              int    `gorm:"size:1;" json:"isAuth"`
	IsLogicalDelete     string `gorm:"size:1;" json:"isLogicalDelete"`
	LogicalDelete       bool   `gorm:"size:1;" json:"logicalDelete"`
	LogicalDeleteColumn string `gorm:"size:128;" json:"logicalDeleteColumn"`
	ModelTime
	ControlBy
}

func (SysTables) TableName() string {
	return "sys_tables"
}
