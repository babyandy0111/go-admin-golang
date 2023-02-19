package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

func InitExamplesRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 無需認證
	examplesNoCheckRoleRouter(r)
	// 需要認證
	examplesCheckRoleRouter(r, authMiddleware)

	return r
}

func examplesNoCheckRoleRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

func examplesCheckRoleRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {
	v1 := r.Group("/api/v1")
	for _, f := range routerCheckRole {
		f(v1, authMiddleware)
	}
}
