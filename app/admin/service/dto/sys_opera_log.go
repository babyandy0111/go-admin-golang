package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

const (
	OperaStatusEnabel  = "1" // 狀態-正常
	OperaStatusDisable = "2" // 狀態-關閉
)

type SysOperaLogGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title" search:"type:contains;column:title;table:sys_opera_log" comment:"操作模組"`
	Method         string `form:"method" search:"type:contains;column:method;table:sys_opera_log" comment:"函數"`
	RequestMethod  string `form:"requestMethod" search:"type:contains;column:request_method;table:sys_opera_log" comment:"請求方式: GET POST PUT DELETE"`
	OperUrl        string `form:"operUrl" search:"type:contains;column:oper_url;table:sys_opera_log" comment:"訪問位置"`
	OperIp         string `form:"operIp" search:"type:exact;column:oper_ip;table:sys_opera_log" comment:"client ip"`
	Status         int    `form:"status" search:"type:exact;column:status;table:sys_opera_log" comment:"狀態 1:正常 2:關閉"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:sys_opera_log" comment:"建立時間"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:sys_opera_log" comment:"更新時間"`
	SysOperaLogOrder
}

type SysOperaLogOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_opera_log" form:"createdAtOrder"`
}

func (m *SysOperaLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysOperaLogControl struct {
	ID            int       `uri:"Id" comment:"流水號"` // 流水號
	Title         string    `json:"title" comment:"操作模組"`
	BusinessType  string    `json:"businessType" comment:"操作類型"`
	BusinessTypes string    `json:"businessTypes" comment:""`
	Method        string    `json:"method" comment:"函數"`
	RequestMethod string    `json:"requestMethod" comment:"請求方式"`
	OperatorType  string    `json:"operatorType" comment:"操作類型"`
	OperName      string    `json:"operName" comment:"操作者"`
	DeptName      string    `json:"deptName" comment:"部門名稱"`
	OperUrl       string    `json:"operUrl" comment:"訪問位置"`
	OperIp        string    `json:"operIp" comment:"client ip"`
	OperLocation  string    `json:"operLocation" comment:"訪問位置"`
	OperParam     string    `json:"operParam" comment:"請求參數"`
	Status        string    `json:"status" comment:"操作狀態"`
	OperTime      time.Time `json:"operTime" comment:"操作時間"`
	JsonResult    string    `json:"jsonResult" comment:"response"`
	Remark        string    `json:"remark" comment:"備註"`
	LatencyTime   string    `json:"latencyTime" comment:"response time"`
	UserAgent     string    `json:"userAgent" comment:"ua"`
}

func (s *SysOperaLogControl) Generate() (*models.SysOperaLog, error) {
	return &models.SysOperaLog{
		Model:         common.Model{Id: s.ID},
		Title:         s.Title,
		BusinessType:  s.BusinessType,
		BusinessTypes: s.BusinessTypes,
		Method:        s.Method,
		RequestMethod: s.RequestMethod,
		OperatorType:  s.OperatorType,
		OperName:      s.OperName,
		DeptName:      s.DeptName,
		OperUrl:       s.OperUrl,
		OperIp:        s.OperIp,
		OperLocation:  s.OperLocation,
		OperParam:     s.OperParam,
		Status:        s.Status,
		OperTime:      s.OperTime,
		JsonResult:    s.JsonResult,
		Remark:        s.Remark,
		LatencyTime:   s.LatencyTime,
		UserAgent:     s.UserAgent,
	}, nil
}

func (s *SysOperaLogControl) GetId() interface{} {
	return s.ID
}

type SysOperaLogGetReq struct {
	Id int `uri:"id"`
}

func (s *SysOperaLogGetReq) GetId() interface{} {
	return s.Id
}

// SysOperaLogDeleteReq 功能刪除請求參數
type SysOperaLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysOperaLogDeleteReq) GetId() interface{} {
	return s.Ids
}
