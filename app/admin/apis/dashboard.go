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
)

type Dashboard struct {
	api.Api
}

// GetSalesByM 取得業績年度每月報價總數
// @Summary 查詢內容列表
// @Description 查詢內容列表
// @Tags 內容
// @Param year query string false "年份"
// @Param currency query int false "幣別"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.GetSalesByMRes}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dashboard/GetSalesByM [get]
// @Security Bearer
func (e Dashboard) GetSalesByM(c *gin.Context) {
	req := dto.DashboardReq{}
	s := service.Dashboard{}
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

	list := make([]models.GetSalesByMRes, 0)
	var count int64

	userId := user.GetUserId(c)
	err = s.GetSalesByM(&req, &list, userId)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢內容失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

func (e Dashboard) GetSalesTop15(c *gin.Context) {
	req := dto.DashboardReq{}
	s := service.Dashboard{}
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

	list := make([]models.GetSalesTop20, 0)
	var count int64

	err = s.GetSalesTop15(&req, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢內容失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

func (e Dashboard) GetSalesByMAccount(c *gin.Context) {
	req := dto.DashboardReq{}
	s := service.Dashboard{}
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

	list := make([]models.GetSalesByMAccount, 0)
	var count int64
	userId := user.GetUserId(c)
	err = s.GetSalesByMAccount(&req, &list, userId)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢內容失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}

func (e Dashboard) GetSalesByProduct(c *gin.Context) {
	req := dto.DashboardReq{}
	s := service.Dashboard{}
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

	list := make([]models.GetSalesByProduct, 0)
	var count int64
	userId := user.GetUserId(c)
	err = s.GetSalesByProduct(&req, &list, userId)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("查詢內容失敗，\r\n err: %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查詢成功")
}
