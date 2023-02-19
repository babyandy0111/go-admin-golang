package api

import "go-admin/app/other/router"

func init() {
	//註冊路由
	AppRouters = append(AppRouters, router.InitRouter)
}
