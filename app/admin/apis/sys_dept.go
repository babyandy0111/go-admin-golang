package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysDept struct {
	api.Api
}

// GetPage
// @Summary 分頁部門列表資料
// @Description 分頁列表
// @Tags 部門
// @Param deptName query string false "deptName"
// @Param deptId query string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept [get]
// @Security Bearer
func (e SysDept) GetPage(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
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
	list := make([]models.SysDept, 0)
	list, err = s.SetDeptPage(&req)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(list, "查詢成功")
}

// Get
// @Summary 取得部門資料
// @Description 取得JSON
// @Tags 部門
// @Param deptId path string false "deptId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security Bearer
func (e SysDept) Get(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetReq{}
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
	var object models.SysDept

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}

	e.OK(object, "查詢成功")
}

// Insert 新增部門
// @Summary 新增部門
// @Description 取得JSON
// @Tags 部門
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDeptInsertReq true "data"
// @Success 200 {string} string	"{"code": 200, "message": "新增成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "新增失敗"}"
// @Router /api/v1/dept [post]
// @Security Bearer
func (e SysDept) Insert(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptInsertReq{}
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

	// 設定建立人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "建立失敗")
		return
	}
	e.OK(req.GetId(), "建立成功")
}

// Update
// @Summary 更新部門
// @Description 取得JSON
// @Tags 部門
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysDeptUpdateReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "新增成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "新增失敗"}"
// @Router /api/v1/dept/{deptId} [put]
// @Security Bearer
func (e SysDept) Update(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptUpdateReq{}
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
	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 刪除部門
// @Description 刪除資料
// @Tags 部門
// @Param data body dto.SysDeptDeleteReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "刪除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "刪除失敗"}"
// @Router /api/v1/dept [delete]
// @Security Bearer
func (e SysDept) Delete(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptDeleteReq{}
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

// Get2Tree User管理 左侧部門树
func (e SysDept) Get2Tree(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
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
	list := make([]dto.DeptLabel, 0)
	list, err = s.SetDeptTree(&req)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(list, "")
}

// GetDeptTreeRoleSelect TODO: 此API需要調整不应该将list和选中放在一起
func (e SysDept) GetDeptTreeRoleSelect(c *gin.Context) {
	s := service.SysDept{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := pkg.StringToInt(c.Param("roleId"))
	result, err := s.SetDeptLabel()
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = s.GetWithRoleId(id)
		if err != nil {
			e.Error(500, err, err.Error())
			return
		}
	}
	e.OK(gin.H{
		"depts":       result,
		"checkedKeys": menuIds,
	}, "")
}
