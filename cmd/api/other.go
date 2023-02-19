package api

import "go-admin/app/other/router"

func init() {
	//注册路由
	AppRouters = append(AppRouters, router.InitRouter)
}
