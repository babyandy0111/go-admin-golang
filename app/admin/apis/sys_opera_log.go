package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysOperaLog struct {
	api.Api
}

// GetPage 操作Log列表
// @Summary 操作Log列表
// @Description 取得JSON
// @Tags 操作Log
// @Param title query string false "title"
// @Param method query string false "method"
// @Param requestMethod  query string false "requestMethod"
// @Param operUrl query string false "operUrl"
// @Param operIp query string false "operIp"
// @Param status query string false "status"
// @Param beginTime query string false "beginTime"
// @Param endTime query string false "endTime"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-opera-log [get]
// @Security Bearer
func (e SysOperaLog) GetPage(c *gin.Context) {
	s := service.SysOperaLog{}
	req := new(dto.SysOperaLogGetPageReq)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.SysOperaLog, 0)
	var count int64

	err = s.GetPage(req, &list, &count)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 操作Log透過id取得
// @Summary 操作Log透過id取得
// @Description 取得JSON
// @Tags 操作Log
// @Param id path string false "id"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-opera-log/{id} [get]
// @Security Bearer
func (e SysOperaLog) Get(c *gin.Context) {
	s := new(service.SysOperaLog)
	req := dto.SysOperaLogGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysOperaLog
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(object, "查詢成功")
}

// Delete 操作Log刪除
// DeleteSysMenu 操作Log刪除
// @Summary 刪除操作Log
// @Description 刪除資料
// @Tags 操作Log
// @Param data body dto.SysOperaLogDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-opera-log [delete]
// @Security Bearer
func (e SysOperaLog) Delete(c *gin.Context) {
	s := new(service.SysOperaLog)
	req := dto.SysOperaLogDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf("刪除失敗！錯误详情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
