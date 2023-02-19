package api

import "go-admin/app/jobs/router"

func init() {
	//注册路由
	AppRouters = append(AppRouters, router.InitRouter)
}
