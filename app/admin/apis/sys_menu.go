package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/models"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysMenu struct {
	api.Api
}

// GetPage Menu列表數据
// @Summary Menu列表數据
// @Description 獲取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [get]
// @Security Bearer
func (e SysMenu) GetPage(c *gin.Context) {
	s := service.SysMenu{}
	req := dto.SysMenuGetPageReq{}
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
	var list = make([]models.SysMenu, 0)
	err = s.GetPage(&req, &list).Error
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(list, "查詢成功")
}

// Get 獲取菜单详情
// @Summary Menu详情數据
// @Description 獲取JSON
// @Tags 菜单
// @Param id path string false "id"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu/{id} [get]
// @Security Bearer
func (e SysMenu) Get(c *gin.Context) {
	req := dto.SysMenuGetReq{}
	s := new(service.SysMenu)
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
	var object = models.SysMenu{}

	err = s.Get(&req, &object).Error
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.OK(object, "查詢成功")
}

// Insert 創建菜单
// @Summary 創建菜单
// @Description 獲取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysMenuInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [post]
// @Security Bearer
func (e SysMenu) Insert(c *gin.Context) {
	req := dto.SysMenuInsertReq{}
	s := new(service.SysMenu)
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
	// 设置創建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req).Error
	if err != nil {
		e.Error(500, err, "創建失敗")
		return
	}
	e.OK(req.GetId(), "創建成功")
}

// Update 修改菜单
// @Summary 修改菜单
// @Description 獲取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysMenuUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu/{id} [put]
// @Security Bearer
func (e SysMenu) Update(c *gin.Context) {
	req := dto.SysMenuUpdateReq{}
	s := new(service.SysMenu)
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
	err = s.Update(&req).Error
	if err != nil {
		e.Error(500, err, "更新失敗")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除數据
// @Tags 菜单
// @Param data body dto.SysMenuDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [delete]
// @Security Bearer
func (e SysMenu) Delete(c *gin.Context) {
	control := new(dto.SysMenuDeleteReq)
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(control, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Remove(control).Error
	if err != nil {
		e.Logger.Errorf("RemoveSysMenu error, %s", err)
		e.Error(500, err, "删除失敗")
		return
	}
	e.OK(control.GetId(), "删除成功")
}

// GetMenuRole 根据登录角色名称獲取菜单列表數据（左菜单使用）
// @Summary 根据登录角色名称獲取菜单列表數据（左菜单使用）
// @Description 獲取JSON
// @Tags 菜单
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menurole [get]
// @Security Bearer
func (e SysMenu) GetMenuRole(c *gin.Context) {
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.SetMenuRole(user.GetRoleName(c))

	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}

	e.OK(result, "")
}

//// GetMenuIDS 獲取角色对应的菜单id數组
//// @Summary 獲取角色对应的菜单id數组，设置角色权限使用
//// @Description 獲取JSON
//// @Tags 菜单
//// @Param id path int true "id"
//// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
//// @Router /api/v1/menuids/{id} [get]
//// @Security Bearer
//func (e SysMenu) GetMenuIDS(c *gin.Context) {
//	s := new(service.SysMenu)
//	r := service.SysRole{}
//	m := dto.SysRoleByName{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&m, binding.JSON).
//		MakeService(&s.Service).
//		MakeService(&r.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	var data models.SysRole
//	err = r.GetWithName(&m, &data).Error
//
//	//data.RoleName = c.GetString("role")
//	//data.UpdateBy = user.GetUserId(c)
//	//result, err := data.GetIDS(s.Orm)
//
//	if err != nil {
//		e.Logger.Errorf("GetIDS error, %s", err.Error())
//		e.Error(500, err, "獲取失敗")
//		return
//	}
//	e.OK(result, "")
//}

// GetMenuTreeSelect 根据角色ID查詢菜单下拉树结构
// @Summary 角色修改使用的菜单列表
// @Description 獲取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param roleId path int true "roleId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menuTreeselect/{roleId} [get]
// @Security Bearer
func (e SysMenu) GetMenuTreeSelect(c *gin.Context) {
	m := service.SysMenu{}
	r := service.SysRole{}
	req := dto.SelectRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&m.Service).
		MakeService(&r.Service).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	result, err := m.SetLabel()
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}

	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = r.GetRoleMenuId(req.RoleId)
		if err != nil {
			e.Error(500, err, "")
			return
		}
	}
	e.OK(gin.H{
		"menus":       result,
		"checkedKeys": menuIds,
	}, "獲取成功")
}
