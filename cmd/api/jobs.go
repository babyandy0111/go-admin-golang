package api

import "go-admin/app/jobs/router"

func init() {
	//註冊路由
	AppRouters = append(AppRouters, router.InitRouter)
}
