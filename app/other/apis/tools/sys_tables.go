package tools

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"gorm.io/gorm"

	"go-admin/app/other/models/tools"
)

type SysTable struct {
	api.Api
}

// GetPage 分頁列表資料
// @Summary 分頁列表資料
// @Description 生成表分頁列表
// @Tags 工具 / 生成工具
// @Param tableName query string false "tableName / 資料表名稱"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页碼"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/page [get]
func (e SysTable) GetPage(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	var data tools.SysTables
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = pkg.StringToInt(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = pkg.StringToInt(index)
	}

	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫連接取得失敗")
		return
	}

	data.TBName = c.Request.FormValue("tableName")
	data.TableComment = c.Request.FormValue("tableComment")
	result, count, err := data.GetPage(db, pageSize, pageIndex)
	if err != nil {
		log.Errorf("GetPage error, %s", err.Error())
		e.Error(500, err, "")
		return
	}
	e.PageOK(result, count, pageIndex, pageSize, "查詢成功")
}

// Get
// @Summary 取得配置
// @Description 取得JSON
// @Tags 工具 / 生成工具
// @Param configKey path int true "configKey"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/info/{tableId} [get]
// @Security Bearer
func (e SysTable) Get(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫連接取得失敗")
		return
	}

	var data tools.SysTables
	data.TableId, _ = pkg.StringToInt(c.Param("tableId"))
	result, err := data.Get(db, true)
	if err != nil {
		log.Errorf("Get error, %s", err.Error())
		e.Error(500, err, "")
		return
	}

	mp := make(map[string]interface{})
	mp["list"] = result.Columns
	mp["info"] = result
	e.OK(mp, "")
}

func (e SysTable) GetSysTablesInfo(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫連接取得失敗")
		return
	}

	var data tools.SysTables
	if c.Request.FormValue("tableName") != "" {
		data.TBName = c.Request.FormValue("tableName")
	}
	result, err := data.Get(db, true)
	if err != nil {
		log.Errorf("Get error, %s", err.Error())
		e.Error(500, err, "抱歉未找到相关訊息")
		return
	}

	mp := make(map[string]interface{})
	mp["list"] = result.Columns
	mp["info"] = result
	e.OK(mp, "")
	//res.Data = mp
	//c.JSON(http.StatusOK, res.ReturnOK())
}

func (e SysTable) GetSysTablesTree(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫連接取得失敗")
		return
	}

	var data tools.SysTables
	result, err := data.GetTree(db)
	if err != nil {
		log.Errorf("GetTree error, %s", err.Error())
		e.Error(500, err, "抱歉未找到相关訊息")
		return
	}

	e.OK(result, "")
}

// Insert
// @Summary 新增表結構
// @Description 新增表結構
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param tables query string false "tableName / 資料表名稱"
// @Success 200 {string} string	"{"code": 200, "message": "新增成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "新增失敗"}"
// @Router /api/v1/sys/tables/info [post]
// @Security Bearer
func (e SysTable) Insert(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫連接取得失敗")
		return
	}

	tablesList := strings.Split(c.Request.FormValue("tables"), ",")
	for i := 0; i < len(tablesList); i++ {

		data, err := genTableInit(db, tablesList, i, c)
		if err != nil {
			log.Errorf("genTableInit error, %s", err.Error())
			e.Error(500, err, "")
			return
		}

		_, err = data.Create(db)
		if err != nil {
			log.Errorf("Create error, %s", err.Error())
			e.Error(500, err, "")
			return
		}
	}
	e.OK(nil, "新增成功")

}

func genTableInit(tx *gorm.DB, tablesList []string, i int, c *gin.Context) (tools.SysTables, error) {
	var data tools.SysTables
	var dbTable tools.DBTables
	var dbColumn tools.DBColumns
	data.TBName = tablesList[i]
	data.CreateBy = 0

	dbTable.TableName = data.TBName
	dbtable, err := dbTable.Get(tx)
	if err != nil {
		return data, err
	}

	dbColumn.TableName = data.TBName
	tablenamelist := strings.Split(dbColumn.TableName, "_")
	for i := 0; i < len(tablenamelist); i++ {
		strStart := string([]byte(tablenamelist[i])[:1])
		strend := string([]byte(tablenamelist[i])[1:])
		// 大驼峰表名 結構体使用
		data.ClassName += strings.ToUpper(strStart) + strend
		// 小驼峰表名 js函數名和權限標识使用
		if i == 0 {
			data.BusinessName += strings.ToLower(strStart) + strend
		} else {
			data.BusinessName += strings.ToUpper(strStart) + strend
		}
		//data.PackageName += strings.ToLower(strStart) + strings.ToLower(strend)
		//data.ModuleName += strings.ToLower(strStart) + strings.ToLower(strend)
	}
	//data.ModuleFrontName = strings.ReplaceAll(data.ModuleName, "_", "-")
	data.PackageName = "admin"
	data.TplCategory = "crud"
	data.Crud = true
	// 中横线表名稱，API路径、前端文件夾名稱和js名稱使用
	data.ModuleName = strings.Replace(data.TBName, "_", "-", -1)
	dbcolumn, err := dbColumn.GetList(tx)
	data.CreateBy = 0
	data.TableComment = dbtable.TableComment
	if dbtable.TableComment == "" {
		data.TableComment = data.ClassName
	}

	data.FunctionName = data.TableComment
	//data.BusinessName = data.ModuleName
	data.IsLogicalDelete = "1"
	data.LogicalDelete = true
	data.LogicalDeleteColumn = "is_del"
	data.IsActions = 2
	data.IsDataScope = 1
	data.IsAuth = 1

	data.FunctionAuthor = "wenjianzhang"
	for i := 0; i < len(dbcolumn); i++ {
		var column tools.SysColumns
		column.ColumnComment = dbcolumn[i].ColumnComment
		column.ColumnName = dbcolumn[i].ColumnName
		column.ColumnType = dbcolumn[i].ColumnType
		column.Sort = i + 1
		column.Insert = true
		column.IsInsert = "1"
		column.QueryType = "EQ"
		column.IsPk = "0"

		namelist := strings.Split(dbcolumn[i].ColumnName, "_")
		for i := 0; i < len(namelist); i++ {
			strStart := string([]byte(namelist[i])[:1])
			strend := string([]byte(namelist[i])[1:])
			column.GoField += strings.ToUpper(strStart) + strend
			if i == 0 {
				column.JsonField = strings.ToLower(strStart) + strend
			} else {
				column.JsonField += strings.ToUpper(strStart) + strend
			}
		}
		if strings.Contains(dbcolumn[i].ColumnKey, "PR") {
			column.IsPk = "1"
			column.Pk = true
			data.PkColumn = dbcolumn[i].ColumnName
			//column.GoField = strings.ToUpper(column.GoField)
			//column.JsonField = strings.ToUpper(column.JsonField)
			data.PkGoField = column.GoField
			data.PkJsonField = column.JsonField
		}
		column.IsRequired = "0"
		if strings.Contains(dbcolumn[i].IsNullable, "NO") {
			column.IsRequired = "1"
			column.Required = true
		}

		if strings.Contains(dbcolumn[i].ColumnType, "int") {
			if strings.Contains(dbcolumn[i].ColumnKey, "PR") {
				column.GoType = "int"
			} else {
				column.GoType = "string"
			}
			column.HtmlType = "input"
		} else if strings.Contains(dbcolumn[i].ColumnType, "timestamp") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
		} else if strings.Contains(dbcolumn[i].ColumnType, "datetime") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
		} else {
			column.GoType = "string"
			column.HtmlType = "input"
		}

		data.Columns = append(data.Columns, column)
	}
	return data, err
}

// Update
// @Summary 更新表結構
// @Description 更新表結構
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body tools.SysTables true "body"
// @Success 200 {string} string	"{"code": 200, "message": "新增成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "新增失敗"}"
// @Router /api/v1/sys/tables/info [put]
// @Security Bearer
func (e SysTable) Update(c *gin.Context) {
	var data tools.SysTables
	err := c.Bind(&data)
	pkg.HasError(err, "資料解析失敗", 500)

	e.Context = c
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫連接取得失敗")
		return
	}

	data.UpdateBy = 0
	result, err := data.Update(db)
	if err != nil {
		log.Errorf("Update error, %s", err.Error())
		e.Error(500, err, "")
		return
	}
	e.OK(result, "更新成功")
}

// Delete
// @Summary 刪除表結構
// @Description 刪除表結構
// @Tags 工具 / 生成工具
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "刪除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "刪除失敗"}"
// @Router /api/v1/sys/tables/info/{tableId} [delete]
func (e SysTable) Delete(c *gin.Context) {
	e.Context = c
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Errorf("get db connection error, %s", err.Error())
		e.Error(500, err, "資料庫連接取得失敗")
		return
	}

	var data tools.SysTables
	IDS := pkg.IdsStrToIdsIntGroup("tableId", c)
	_, err = data.BatchDelete(db, IDS)
	if err != nil {
		log.Errorf("BatchDelete error, %s", err.Error())
		e.Error(500, err, "刪除失敗")
		return
	}
	e.OK(nil, "刪除成功")
}
