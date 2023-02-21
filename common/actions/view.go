package actions

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"gorm.io/gorm"

	"go-admin/common/dto"
	"go-admin/common/models"
)

// ViewAction 通用詳情動作
func ViewAction(control dto.Control, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		//查看詳情
		req := control.Generate()
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

		var rsp interface{}
		if f != nil {
			rsp = f()
		} else {
			rsp, _ = req.GenerateM()
		}

		//資料權限檢查
		p := GetPermissionFromContext(c)

		err = db.Model(object).WithContext(c).Scopes(
			Permission(object.TableName(), p),
		).Where(req.GetId()).First(rsp).Error

		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, nil, "查無資料或無權限")
			return
		}
		if err != nil {
			log.Errorf("MsgID[%s] View error: %s", msgID, err)
			response.Error(c, 500, err, "查看失敗")
			return
		}
		response.OK(c, rsp, "查詢成功")
		c.Next()
	}
}
