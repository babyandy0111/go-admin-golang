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

type Article struct {
	api.Api
}

// GetPage 查詢內容列表
// @Summary 查詢內容列表
// @Description 查詢內容列表
// @Tags 內容
// @Param status query string false "狀態"
// @Param pageSize query int false "每頁幾筆"
// @Param pageIndex query int false "頁數"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Article}} "{"code": 200, "data": [...]}"
// @Router /api/v1/article [get]
// @Security Bearer
func (e Article) GetPage(c *gin.Context) {
	req := dto.ArticleGetPageReq{}
	s := service.Article{}
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
	list := make([]models.Article, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢內容失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 查詢內容單筆
// @Summary 查詢內容
// @Description 查詢內容
// @Tags 內容
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Article} "{"code": 200, "data": [...]}"
// @Router /api/v1/article/{id} [get]
// @Security Bearer
func (e Article) Get(c *gin.Context) {
	req := dto.ArticleGetReq{}
	s := service.Article{}
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
	var object models.Article

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢內容失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(object, "查詢成功")
}

// Insert 建立內容
// @Summary 建立內容
// @Description 建立內容
// @Tags 內容
// @Accept application/json
// @Product application/json
// @Param data body dto.ArticleInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "建立成功"}"
// @Router /api/v1/article [post]
// @Security Bearer
func (e Article) Insert(c *gin.Context) {
	req := dto.ArticleInsertReq{}
	s := service.Article{}
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
		e.Error(500, err, fmt.Sprintf("建立內容失敗，\r\n err: %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "建立成功")
}

// Update 更新內容
// @Summary 更新內容
// @Description 更新內容
// @Tags 內容
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ArticleUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/article/{id} [put]
// @Security Bearer
func (e Article) Update(c *gin.Context) {
	req := dto.ArticleUpdateReq{}
	s := service.Article{}
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
		e.Error(500, err, fmt.Sprintf("更新內容失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 刪除內容
// @Summary 刪除內容
// @Description 刪除內容
// @Tags 內容
// @Param data body dto.ArticleDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/article [delete]
// @Security Bearer
func (e Article) Delete(c *gin.Context) {
	s := service.Article{}
	req := dto.ArticleDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("刪除內容失敗，\r\n err: %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
