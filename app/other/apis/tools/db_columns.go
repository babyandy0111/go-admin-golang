package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/other/models/tools"
)

// GetDBColumnList 分頁列表資料
// @Summary 分頁列表資料 / page list data
// @Description 資料庫表列分頁列表 / database table column page list
// @Tags 工具 / 生成工具
// @Param tableName query string false "tableName / 資料表名稱"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页碼"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/columns/page [get]
func (e *Gen) GetDBColumnList(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	var data tools.DBColumns
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = pkg.StringToInt(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = pkg.StringToInt(index)
	}

	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫连接取得失敗")
		return
	}

	data.TableName = c.Request.FormValue("tableName")
	pkg.Assert(data.TableName == "", "table name cannot be empty！", 500)
	result, count, err := data.GetPage(db, pageSize, pageIndex)
	if err != nil {
		log.Errorf("GetPage error, %s", err.Error())
		e.Error(500, err, "")
		return
	}
	e.PageOK(result, count, pageIndex, pageSize, "查詢成功")
}
