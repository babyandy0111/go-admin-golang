package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/models"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysLoginLog struct {
	api.Api
}

// GetPage 登入Log列表
// @Summary 登入Log列表
// @Description 取得JSON
// @Tags 登入Log
// @Param username query string false "使用者"
// @Param ipaddr query string false "ip地址"
// @Param loginLocation  query string false "歸屬地"
// @Param status query string false "狀態"
// @Param beginTime query string false "开始時間"
// @Param endTime query string false "结束時間"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-login-log [get]
// @Security Bearer
func (e SysLoginLog) GetPage(c *gin.Context) {
	s := service.SysLoginLog{}
	req := dto.SysLoginLogGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list := make([]models.SysLoginLog, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 登入Log通过id取得
// @Summary 登入Log通过id取得
// @Description 取得JSON
// @Tags 登入Log
// @Param id path string false "id"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-login-log/{id} [get]
// @Security Bearer
func (e SysLoginLog) Get(c *gin.Context) {
	s := service.SysLoginLog{}
	req := dto.SysLoginLogGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysLoginLog
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(object, "查詢成功")
}

// Delete 登入Log刪除
// @Summary 登入Log刪除
// @Description 登入Log刪除
// @Tags 登入Log
// @Param data body dto.SysLoginLogDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-login-log [delete]
// @Security Bearer
func (e SysLoginLog) Delete(c *gin.Context) {
	s := service.SysLoginLog{}
	req := dto.SysLoginLogDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, "刪除失敗")
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
