package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysPost struct {
	api.Api
}

// GetPage
// @Summary 職稱列表資料
// @Description 取得JSON
// @Tags 職稱
// @Param postName query string false "postName"
// @Param postCode query string false "postCode"
// @Param postId query string false "postId"
// @Param status query string false "status"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @Security Bearer
func (e SysPost) GetPage(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostPageReq{}
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

	list := make([]models.SysPost, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get
// @Summary 取得職稱訊息
// @Description 取得JSON
// @Tags 職稱
// @Param id path int true "流水號"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{postId} [get]
// @Security Bearer
func (e SysPost) Get(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostGetReq{}
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
	var object models.SysPost

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("職稱訊息取得失敗！錯误详情：%s", err.Error()))
		return
	}

	e.OK(object, "查詢成功")
}

// Insert
// @Summary 新增職稱
// @Description 取得JSON
// @Tags 職稱
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysPostInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [post]
// @Security Bearer
func (e SysPost) Insert(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostInsertReq{}
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
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("新建職稱失敗！錯误详情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "建立成功")
}

// Update
// @Summary 更新職稱
// @Description 取得JSON
// @Tags 職稱
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysPostUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{id} [put]
// @Security Bearer
func (e SysPost) Update(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostUpdateReq{}
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

	req.SetUpdateBy(user.GetUserId(c))

	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("職稱更新失敗！錯误详情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 刪除職稱
// @Description 刪除資料
// @Tags 職稱
// @Param id body dto.SysPostDeleteReq true "請求參數"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [delete]
// @Security Bearer
func (e SysPost) Delete(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostDeleteReq{}
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
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("職稱刪除失敗！錯误详情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "刪除成功")
}
