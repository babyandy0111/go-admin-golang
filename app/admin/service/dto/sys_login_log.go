package dto

import (
	"time"

	"go-admin/common/dto"
)

type SysLoginLogGetPageReq struct {
	dto.Pagination `search:"-"`
	Username       string `form:"username" search:"type:exact;column:username;table:sys_login_log" comment:"使用者"`
	Status         string `form:"status" search:"type:exact;column:status;table:sys_login_log" comment:"狀態"`
	Ipaddr         string `form:"ipaddr" search:"type:exact;column:ipaddr;table:sys_login_log" comment:"ip地址"`
	LoginLocation  string `form:"loginLocation" search:"type:exact;column:login_location;table:sys_login_log" comment:"歸屬地"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:ctime;table:sys_login_log" comment:"建立時間"`
	EndTime        string `form:"endTime" search:"type:lte;column:ctime;table:sys_login_log" comment:"建立時間"`
	SysLoginLogOrder
}

type SysLoginLogOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_login_log" form:"createdAtOrder"`
}

func (m *SysLoginLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysLoginLogControl struct {
	ID            int       `uri:"Id" comment:"主鍵"` // 主鍵
	Username      string    `json:"username" comment:"使用者"`
	Status        string    `json:"status" comment:"狀態"`
	Ipaddr        string    `json:"ipaddr" comment:"ip地址"`
	LoginLocation string    `json:"loginLocation" comment:"歸屬地"`
	Browser       string    `json:"browser" comment:"瀏覽器"`
	Os            string    `json:"os" comment:"系統"`
	Platform      string    `json:"platform" comment:"平台"`
	LoginTime     time.Time `json:"loginTime" comment:"登入時間"`
	Remark        string    `json:"remark" comment:"備註"`
	Msg           string    `json:"msg" comment:"訊息"`
}

type SysLoginLogGetReq struct {
	Id int `uri:"id"`
}

func (s *SysLoginLogGetReq) GetId() interface{} {
	return s.Id
}

// SysLoginLogDeleteReq 功能刪除請求參數
type SysLoginLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysLoginLogDeleteReq) GetId() interface{} {
	return s.Ids
}
