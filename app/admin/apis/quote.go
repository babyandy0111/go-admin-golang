package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Quote struct {
	api.Api
}

// GetPage 查詢報價單管理列表
// @Summary 查詢報價單管理列表
// @Description 查詢報價單管理列表
// @Tags 報價單管理
// @Param quoteNumber query string false "報價單號"
// @Param companyTypeId query string false "公司流水號"
// @Param accountId query string false "客戶"
// @Param pageSize query int false "每頁幾筆"
// @Param pageIndex query int false "頁數"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Quote}} "{"code": 200, "data": [...]}"
// @Router /api/v1/quote [get]
// @Security Bearer
func (e Quote) GetPage(c *gin.Context) {
	req := dto.QuoteGetPageReq{}
	s := service.Quote{}
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Quote, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢報價單管理失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 查詢報價單管理單筆
// @Summary 查詢報價單管理
// @Description 查詢報價單管理
// @Tags 報價單管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Quote} "{"code": 200, "data": [...]}"
// @Router /api/v1/quote/{id} [get]
// @Security Bearer
func (e Quote) Get(c *gin.Context) {
	req := dto.QuoteGetReq{}
	s := service.Quote{}
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
	var object models.Quote

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢報價單管理失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(object, "查詢成功")
}

// Insert 建立報價單管理
// @Summary 建立報價單管理
// @Description 建立報價單管理
// @Tags 報價單管理
// @Accept application/json
// @Product application/json
// @Param data body dto.QuoteInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "建立成功"}"
// @Router /api/v1/quote [post]
// @Security Bearer
func (e Quote) Insert(c *gin.Context) {
	req := dto.QuoteInsertReq{}
	s := service.Quote{}
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
	// 設定建立者
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("建立報價單管理失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "建立成功")
}

// Update 更新報價單管理
// @Summary 更新報價單管理
// @Description 更新報價單管理
// @Tags 報價單管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.QuoteUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/quote/{id} [put]
// @Security Bearer
func (e Quote) Update(c *gin.Context) {
	req := dto.QuoteUpdateReq{}
	s := service.Quote{}
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
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("更新報價單管理失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 刪除報價單管理
// @Summary 刪除報價單管理
// @Description 刪除報價單管理
// @Tags 報價單管理
// @Param data body dto.QuoteDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/quote [delete]
// @Security Bearer
func (e Quote) Delete(c *gin.Context) {
	s := service.Quote{}
	req := dto.QuoteDeleteReq{}
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

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("刪除報價單管理失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
