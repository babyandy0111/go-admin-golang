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
)

type SysDictData struct {
	api.Api
}

// GetPage
// @Summary 字典資料列表
// @Description 取得JSON
// @Tags 字典資料
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页碼"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data [get]
// @Security Bearer
func (e SysDictData) GetPage(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataGetPageReq{}
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

	list := make([]models.SysDictData, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get
// @Summary 通过流水號取得字典資料
// @Description 取得JSON
// @Tags 字典資料
// @Param dictCode path int true "字典流水號"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/{dictCode} [get]
// @Security Bearer
func (e SysDictData) Get(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataGetReq{}
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

	var object models.SysDictData

	err = s.Get(&req, &object)
	if err != nil {
		e.Logger.Warnf("Get error: %s", err.Error())
		e.Error(500, err, "查詢失敗")
		return
	}

	e.OK(object, "查詢成功")
}

// Insert
// @Summary 新增字典資料
// @Description 取得JSON
// @Tags 字典資料
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictDataInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "新增成功"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func (e SysDictData) Insert(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataInsertReq{}
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
		e.Error(500, err, "建立失敗")
		return
	}

	e.OK(req.GetId(), "建立成功")
}

// Update
// @Summary 更新字典資料
// @Description 取得JSON
// @Tags 字典資料
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictDataUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/dict/data/{dictCode} [put]
// @Security Bearer
func (e SysDictData) Update(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataUpdateReq{}
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
		e.Error(500, err, "更新失敗")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 刪除字典資料
// @Description 刪除資料
// @Tags 字典資料
// @Param dictCode body dto.SysDictDataDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/dict/data [delete]
// @Security Bearer
func (e SysDictData) Delete(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataDeleteReq{}
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
		e.Error(500, err, "刪除失敗")
		return
	}
	e.OK(req.GetId(), "刪除成功")
}

// GetAll 資料字典根据key取得 业務页面使用
// @Summary 資料字典根据key取得
// @Description 資料字典根据key取得
// @Tags 字典資料
// @Param dictType query int true "dictType"
// @Success 200 {object} response.Response{data=[]dto.SysDictDataGetAllResp}  "{"code": 200, "data": [...]}"
// @Router /api/v1/dict-data/option-select [get]
// @Security Bearer
func (e SysDictData) GetAll(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataGetPageReq{}
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
	list := make([]models.SysDictData, 0)
	err = s.GetAll(&req, &list)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	l := make([]dto.SysDictDataGetAllResp, 0)
	for _, i := range list {
		d := dto.SysDictDataGetAllResp{}
		e.Translate(i, &d)
		l = append(l, d)
	}

	e.OK(l, "查詢成功")
}
