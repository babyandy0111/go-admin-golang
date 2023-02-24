package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDashboardRouter)
}

// registerArticleRouter
func registerDashboardRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Dashboard{}
	r := v1.Group("/dashboard").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetSalesByM)
		r.GET("GetSalesByM", actions.PermissionAction(), api.GetSalesByM)
		r.GET("GetSalesTop20", actions.PermissionAction(), api.GetSalesTop20)
		r.GET("GetSalesByMAccount", actions.PermissionAction(), api.GetSalesByMAccount)
		r.GET("GetSalesByProduct", actions.PermissionAction(), api.GetSalesByProduct)
	}
}
