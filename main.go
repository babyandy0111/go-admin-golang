package main

import (
	"go-admin/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Element UI的前後端分离權限管理系統的API文档
// @description 新增qq群: 521386980 进入技术交流群 請先star，谢谢！
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
