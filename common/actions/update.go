package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/common/dto"
	"go-admin/common/models"
)

// UpdateAction 通用更新動作
func UpdateAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		req := control.Generate()
		//更新操作
		err = req.Bind(c)
		if err != nil {
			response.Error(c, http.StatusUnprocessableEntity, err, "參數驗證失敗")
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			response.Error(c, 500, err, "model生成失敗")
			return
		}
		object.SetUpdateBy(user.GetUserId(c))

		//資料權限檢查
		p := GetPermissionFromContext(c)

		db = db.WithContext(c).Scopes(
			Permission(object.TableName(), p),
		).Where(req.GetId()).Updates(object)
		if err = db.Error; err != nil {
			log.Errorf("MsgID[%s] Update error: %s", msgID, err)
			response.Error(c, 500, err, "更新失敗")
			return
		}
		if db.RowsAffected == 0 {
			response.Error(c, http.StatusForbidden, nil, "無權限更新該資料")
			return
		}
		response.OK(c, object.GetId(), "更新成功")
		c.Next()
	}
}
