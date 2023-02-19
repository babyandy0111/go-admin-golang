package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysApi struct {
	api.Api
}

// GetPage 取得API管理列表
// @Summary 取得API管理列表
// @Description 取得API管理列表
// @Tags API管理
// @Param name query string false "名稱"
// @Param title query string false "標题"
// @Param path query string false "地址"
// @Param action query string false "類型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页碼"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysApi}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-api [get]
// @Security Bearer
func (e SysApi) GetPage(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiGetPageReq{}
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
	//資料权限检查
	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysApi, 0)
	var count int64
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 取得API管理
// @Summary 取得API管理
// @Description 取得API管理
// @Tags API管理
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SysApi} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-api/{id} [get]
// @Security Bearer
func (e SysApi) Get(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.SysApi{}
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
	p := actions.GetPermissionFromContext(c)
	var object models.SysApi
	err = s.Get(&req, p, &object).Error
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(object, "查詢成功")
}

// Update 更新API管理
// @Summary 更新API管理
// @Description 更新API管理
// @Tags API管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/sys-api/{id} [put]
// @Security Bearer
func (e SysApi) Update(c *gin.Context) {
	req := dto.SysApiUpdateReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)
	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, "更新失敗")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// DeleteSysApi 刪除API管理
// @Summary 刪除API管理
// @Description 刪除API管理
// @Tags API管理
// @Param data body dto.SysApiDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/sys-api [delete]
// @Security Bearer
func (e SysApi) DeleteSysApi(c *gin.Context) {
	req := dto.SysApiDeleteReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	p := actions.GetPermissionFromContext(c)
	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, "刪除失敗")
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
