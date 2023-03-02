package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDashboardRouter)
}

// registerArticleRouter
func registerDashboardRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Dashboard{}
	r := v1.Group("/dashboard").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetSalesByM)
		r.GET("GetSalesByM", api.GetSalesByM)
		r.GET("GetSalesTop20", api.GetSalesTop20)
		r.GET("GetSalesByMAccount", api.GetSalesByMAccount)
		r.GET("GetSalesByProduct", api.GetSalesByProduct)
	}
}
