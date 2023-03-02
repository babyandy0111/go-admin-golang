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

type Account struct {
	api.Api
}

// GetPage 查詢Account列表
// @Summary 查詢Account列表
// @Description 查詢Account列表
// @Tags Account
// @Param pageSize query int false "每頁幾筆"
// @Param pageIndex query int false "頁數"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Account}} "{"code": 200, "data": [...]}"
// @Router /api/v1/account [get]
// @Security Bearer
func (e Account) GetPage(c *gin.Context) {
	req := dto.AccountGetPageReq{}
	s := service.Account{}
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
	list := make([]models.Account, 0)
	var count int64
	fmt.Println(req.PageIndex, req.PageIndex)
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢Account失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 查詢Account單筆
// @Summary 查詢Account
// @Description 查詢Account
// @Tags Account
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Account} "{"code": 200, "data": [...]}"
// @Router /api/v1/account/{id} [get]
// @Security Bearer
func (e Account) Get(c *gin.Context) {
	req := dto.AccountGetReq{}
	s := service.Account{}
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
	var object models.Account

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢Account失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(object, "查詢成功")
}

// Insert 建立Account
// @Summary 建立Account
// @Description 建立Account
// @Tags Account
// @Accept application/json
// @Product application/json
// @Param data body dto.AccountInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "建立成功"}"
// @Router /api/v1/account [post]
// @Security Bearer
func (e Account) Insert(c *gin.Context) {
	req := dto.AccountInsertReq{}
	s := service.Account{}
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
		e.Error(500, err, fmt.Sprintf("建立Account失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "建立成功")
}

// Update 更新Account
// @Summary 更新Account
// @Description 更新Account
// @Tags Account
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AccountUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/account/{id} [put]
// @Security Bearer
func (e Account) Update(c *gin.Context) {
	req := dto.AccountUpdateReq{}
	s := service.Account{}
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
		e.Error(500, err, fmt.Sprintf("更新Account失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 刪除Account
// @Summary 刪除Account
// @Description 刪除Account
// @Tags Account
// @Param data body dto.AccountDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/account [delete]
// @Security Bearer
func (e Account) Delete(c *gin.Context) {
	s := service.Account{}
	req := dto.AccountDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("刪除Account失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
