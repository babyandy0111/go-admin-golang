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

type SysDictType struct {
	api.Api
}

// GetPage 字典類型列表資料
// @Summary 字典類型列表資料
// @Description 取得JSON
// @Tags 字典類型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页碼"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [get]
// @Security Bearer
func (e SysDictType) GetPage(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeGetPageReq{}
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
	list := make([]models.SysDictType, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 字典類型透過字典id取得
// @Summary 字典類型透過字典id取得
// @Description 取得JSON
// @Tags 字典類型
// @Param dictId path int true "字典類型流水號"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/{dictId} [get]
// @Security Bearer
func (e SysDictType) Get(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeGetReq{}
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
	var object models.SysDictType
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(object, "查詢成功")
}

// Insert 字典類型建立
// @Summary 新增字典類型
// @Description 取得JSON
// @Tags 字典類型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [post]
// @Security Bearer
func (e SysDictType) Insert(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeInsertReq{}
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
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf(" 建立字典類型失敗，詳情：%s", err.Error()))
		return
	}
	e.OK(req.GetId(), "建立成功")
}

// Update
// @Summary 更新字典類型
// @Description 取得JSON
// @Tags 字典類型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/{dictId} [put]
// @Security Bearer
func (e SysDictType) Update(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(500, err, err.Error())
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 刪除字典類型
// @Description 刪除資料
// @Tags 字典類型
// @Param dictCode body dto.SysDictTypeDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [delete]
// @Security Bearer
func (e SysDictType) Delete(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeDeleteReq{}
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
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "刪除成功")
}

// GetAll
// @Summary 字典類型全部資料 代碼生成使用API
// @Description 取得JSON
// @Tags 字典類型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type-option-select [get]
// @Security Bearer
func (e SysDictType) GetAll(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeGetPageReq{}
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
	list := make([]models.SysDictType, 0)
	err = s.GetAll(&req, &list)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(list, "查詢成功")
}
