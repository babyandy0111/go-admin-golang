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

type CompanyType struct {
	api.Api
}

// GetPage 查詢CompanyType列表
// @Summary 查詢CompanyType列表
// @Description 查詢CompanyType列表
// @Tags CompanyType
// @Param pageSize query int false "每頁幾筆"
// @Param pageIndex query int false "頁數"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.CompanyType}} "{"code": 200, "data": [...]}"
// @Router /api/v1/company-type [get]
// @Security Bearer
func (e CompanyType) GetPage(c *gin.Context) {
	req := dto.CompanyTypeGetPageReq{}
	s := service.CompanyType{}
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
	list := make([]models.CompanyType, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢CompanyType失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 查詢CompanyType單筆
// @Summary 查詢CompanyType
// @Description 查詢CompanyType
// @Tags CompanyType
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.CompanyType} "{"code": 200, "data": [...]}"
// @Router /api/v1/company-type/{id} [get]
// @Security Bearer
func (e CompanyType) Get(c *gin.Context) {
	req := dto.CompanyTypeGetReq{}
	s := service.CompanyType{}
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
	var object models.CompanyType

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢CompanyType失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(object, "查詢成功")
}

// Insert 建立CompanyType
// @Summary 建立CompanyType
// @Description 建立CompanyType
// @Tags CompanyType
// @Accept application/json
// @Product application/json
// @Param data body dto.CompanyTypeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "建立成功"}"
// @Router /api/v1/company-type [post]
// @Security Bearer
func (e CompanyType) Insert(c *gin.Context) {
	req := dto.CompanyTypeInsertReq{}
	s := service.CompanyType{}
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
		e.Error(500, err, fmt.Sprintf("建立CompanyType失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "建立成功")
}

// Update 更新CompanyType
// @Summary 更新CompanyType
// @Description 更新CompanyType
// @Tags CompanyType
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CompanyTypeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/company-type/{id} [put]
// @Security Bearer
func (e CompanyType) Update(c *gin.Context) {
	req := dto.CompanyTypeUpdateReq{}
	s := service.CompanyType{}
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
		e.Error(500, err, fmt.Sprintf("更新CompanyType失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 刪除CompanyType
// @Summary 刪除CompanyType
// @Description 刪除CompanyType
// @Tags CompanyType
// @Param data body dto.CompanyTypeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/company-type [delete]
// @Security Bearer
func (e CompanyType) Delete(c *gin.Context) {
	s := service.CompanyType{}
	req := dto.CompanyTypeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("刪除CompanyType失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
