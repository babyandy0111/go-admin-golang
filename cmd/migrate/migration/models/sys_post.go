package models

type SysPost struct {
	PostId   int    `gorm:"primaryKey;autoIncrement" json:"postId"` //職稱编号
	PostName string `gorm:"size:128;" json:"postName"`              //職稱名稱
	PostCode string `gorm:"size:128;" json:"postCode"`              //職稱代碼
	Sort     int    `gorm:"size:4;" json:"sort"`                    //職稱排序
	Status   int    `gorm:"size:4;" json:"status"`                  //狀態
	Remark   string `gorm:"size:255;" json:"remark"`                //描述
	ControlBy
	ModelTime
}

func (SysPost) TableName() string {
	return "sys_post"
}
