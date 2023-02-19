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

type SysConfig struct {
	api.Api
}

// GetPage 取得配置管理列表
// @Summary 取得配置管理列表
// @Description 取得配置管理列表
// @Tags 配置管理
// @Param configName query string false "名稱"
// @Param configKey query string false "key"
// @Param configType query string false "類型"
// @Param isFrontend query int false "是否前端"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页碼"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysApi}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-config [get]
// @Security Bearer
func (e SysConfig) GetPage(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	list := make([]models.SysConfig, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

// Get 取得配置管理
// @Summary 取得配置管理
// @Description 取得配置管理
// @Tags 配置管理
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SysConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-config/{id} [get]
// @Security Bearer
func (e SysConfig) Get(c *gin.Context) {
	req := dto.SysConfigGetReq{}
	s := service.SysConfig{}
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
	var object models.SysConfig

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(object, "查詢成功")
}

// Insert 建立配置管理
// @Summary 建立配置管理
// @Description 建立配置管理
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysConfigControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "建立成功"}"
// @Router /api/v1/sys-config [post]
// @Security Bearer
func (e SysConfig) Insert(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigControl{}
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

// Update 更新配置管理
// @Summary 更新配置管理
// @Description 更新配置管理
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysConfigControl true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/sys-config/{id} [put]
// @Security Bearer
func (e SysConfig) Update(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigControl{}
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

// Delete 刪除配置管理
// @Summary 刪除配置管理
// @Description 刪除配置管理
// @Tags 配置管理
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "刪除成功"}"
// @Router /api/v1/sys-config [delete]
// @Security Bearer
func (e SysConfig) Delete(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigDeleteReq{}
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

// Get2SysApp 取得系統配置訊息
// @Summary 取得系統前台配置訊息，主要注意这里不在验证权限
// @Description 取得系統配置訊息，主要注意这里不在验证权限
// @Tags 配置管理
// @Success 200 {object} response.Response{data=map[string]string} "{"code": 200, "data": [...]}"
// @Router /api/v1/app-config [get]
func (e SysConfig) Get2SysApp(c *gin.Context) {
	req := dto.SysConfigGetToSysAppReq{}
	s := service.SysConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	// 控制只读前台的資料
	req.IsFrontend = "1"
	list := make([]models.SysConfig, 0)
	err = s.GetWithKeyList(&req, &list)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	mp := make(map[string]string)
	for i := 0; i < len(list); i++ {
		key := list[i].ConfigKey
		if key != "" {
			mp[key] = list[i].ConfigValue
		}
	}
	e.OK(mp, "查詢成功")
}

// Get2Set 取得配置
// @Summary 取得配置
// @Description 界面操作設定配置值的取得
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Success 200 {object} response.Response{data=map[string]interface{}}	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/set-config [get]
// @Security Bearer
func (e SysConfig) Get2Set(c *gin.Context) {
	s := service.SysConfig{}
	req := make([]dto.GetSetSysConfigReq, 0)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.GetForSet(&req)
	if err != nil {
		e.Error(500, err, "查詢失敗")
		return
	}
	m := make(map[string]interface{}, 0)
	for _, v := range req {
		m[v.ConfigKey] = v.ConfigValue
	}
	e.OK(m, "查詢成功")
}

// Update2Set 設定配置
// @Summary 設定配置
// @Description 界面操作設定配置值
// @Tags 配置管理
// @Accept application/json
// @Product application/json
// @Param data body []dto.GetSetSysConfigReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "更新成功"}"
// @Router /api/v1/set-config [put]
// @Security Bearer
func (e SysConfig) Update2Set(c *gin.Context) {
	s := service.SysConfig{}
	req := make([]dto.GetSetSysConfigReq, 0)
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

	err = s.UpdateForSet(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK("", "更新成功")
}

// GetSysConfigByKEYForService 根据Key取得SysConfig的Service
// @Summary 根据Key取得SysConfig的Service
// @Description 根据Key取得SysConfig的Service
// @Tags 配置管理
// @Param configKey path string false "configKey"
// @Success 200 {object} response.Response{data=dto.SysConfigByKeyReq} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-config/{id} [get]
// @Security Bearer
func (e SysConfig) GetSysConfigByKEYForService(c *gin.Context) {
	var s = new(service.SysConfig)
	var req = new(dto.SysConfigByKeyReq)
	var resp = new(dto.GetSysConfigByKEYForServiceResp)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.GetWithKey(req, resp)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(resp, s.Msg)
}
