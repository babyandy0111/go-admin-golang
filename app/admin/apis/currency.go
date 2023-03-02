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

type Currency struct {
	api.Api
}

// GetPage 查詢Currency列表
// @Summary 查詢Currency列表
// @Description 查詢Currency列表
// @Tags Currency
// @Param pageSize query int false "每頁幾筆"
// @Param pageIndex query int false "頁數"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Currency}} "{"code": 200, "data": [...]}"
// @Router /api/v1/currency [get]
// @Security Bearer
func (e Currency) GetPage(c *gin.Context) {
	req := dto.CurrencyGetPageReq{}
	s := service.Currency{}
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
	list := make([]models.Currency, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢Currency失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 查詢Currency單筆
// @Summary 查詢Currency
// @Description 查詢Currency
// @Tags Currency
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Currency} "{"code": 200, "data": [...]}"
// @Router /api/v1/currency/{id} [get]
// @Security Bearer
func (e Currency) Get(c *gin.Context) {
	req := dto.CurrencyGetReq{}
	s := service.Currency{}
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
	var object models.Currency

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢Currency失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(object, "查詢成功")
}

// Insert 建立Currency
// @Summary 建立Currency
// @Description 建立Currency
// @Tags Currency
// @Accept application/json
// @Product application/json
// @Param data body dto.CurrencyInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "建立成功"}"
// @Router /api/v1/currency [post]
// @Security Bearer
func (e Currency) Insert(c *gin.Context) {
	req := dto.CurrencyInsertReq{}
	s := service.Currency{}
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
		e.Error(500, err, fmt.Sprintf("建立Currency失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "建立成功")
}

// Update 更新Currency
// @Summary 更新Currency
// @Description 更新Currency
// @Tags Currency
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CurrencyUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/currency/{id} [put]
// @Security Bearer
func (e Currency) Update(c *gin.Context) {
	req := dto.CurrencyUpdateReq{}
	s := service.Currency{}
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
		e.Error(500, err, fmt.Sprintf("更新Currency失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 刪除Currency
// @Summary 刪除Currency
// @Description 刪除Currency
// @Tags Currency
// @Param data body dto.CurrencyDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/currency [delete]
// @Security Bearer
func (e Currency) Delete(c *gin.Context) {
	s := service.Currency{}
	req := dto.CurrencyDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("刪除Currency失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
