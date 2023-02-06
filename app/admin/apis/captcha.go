package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"
)

type System struct {
	api.Api
}

// GenerateCaptchaHandler 獲取验证碼
// @Summary 獲取验证碼
// @Description 獲取验证碼
// @Tags 登陆
// @Success 200 {object} response.Response{data=string,id=string,msg=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/captcha [get]
func (e System) GenerateCaptchaHandler(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Error(500, err, "服务初始化失敗！")
		return
	}
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		e.Logger.Errorf("DriverDigitFunc error, %s", err.Error())
		e.Error(500, err, "验证碼獲取失敗")
		return
	}
	e.Custom(gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
