package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"go-admin/app/admin/service/dto"
	"go-admin/common"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"

	"go-admin/common/global"
)

// LoggerToFile Log记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		// 开始時間
		startTime := time.Now()
		// 處理請求
		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				log.Warnf("copy body error, %s", err.Error())
				err = nil
			}
			rb, _ := ioutil.ReadAll(bf)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
			body = string(rb)
		}

		c.Next()
		url := c.Request.RequestURI
		if strings.Index(url, "logout") > -1 ||
			strings.Index(url, "login") > -1 {
			return
		}
		// 结束時間
		endTime := time.Now()
		if c.Request.Method == http.MethodOptions {
			return
		}

		rt, bl := c.Get("result")
		var result = ""
		if bl {
			rb, err := json.Marshal(rt)
			if err != nil {
				log.Warnf("json Marshal result error, %s", err.Error())
			} else {
				result = string(rb)
			}
		}

		st, bl := c.Get("status")
		var statusBus = 0
		if bl {
			statusBus = st.(int)
		}

		// 請求方式
		reqMethod := c.Request.Method
		// 請求路由
		reqUri := c.Request.RequestURI
		// 狀態碼
		statusCode := c.Writer.Status()
		// 請求IP
		clientIP := common.GetClientIP(c)
		// 执行時間
		latencyTime := endTime.Sub(startTime)
		// Log格式
		logData := map[string]interface{}{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}
		log.WithFields(logData).Info()

		if c.Request.Method != "OPTIONS" && config.LoggerConfig.EnabledDB && statusCode != 404 {
			SetDBOperLog(c, clientIP, statusCode, reqUri, reqMethod, latencyTime, body, result, statusBus)
		}
	}
}

// SetDBOperLog 写入操作Log表 fixme 该方法後续即将弃用
func SetDBOperLog(c *gin.Context, clientIP string, statusCode int, reqUri string, reqMethod string, latencyTime time.Duration, body string, result string, status int) {

	log := api.GetRequestLogger(c)
	l := make(map[string]interface{})
	l["_fullPath"] = c.FullPath()
	l["operUrl"] = reqUri
	l["operIp"] = clientIP
	l["operLocation"] = "" // pkg.GetLocation(clientIP, gaConfig.ExtConfig.AMap.Key)
	l["operName"] = user.GetUserName(c)
	l["requestMethod"] = reqMethod
	l["operParam"] = body
	l["operTime"] = time.Now()
	l["jsonResult"] = result
	l["latencyTime"] = latencyTime.String()
	l["statusCode"] = statusCode
	l["userAgent"] = c.Request.UserAgent()
	l["createBy"] = user.GetUserId(c)
	l["updateBy"] = user.GetUserId(c)
	if status == http.StatusOK {
		l["status"] = dto.OperaStatusEnabel
	} else {
		l["status"] = dto.OperaStatusDisable
	}
	q := sdk.Runtime.GetMemoryQueue(c.Request.Host)
	message, err := sdk.Runtime.GetStreamMessage("", global.OperateLog, l)
	if err != nil {
		log.Errorf("GetStreamMessage error, %s", err.Error())
		//Log报錯錯誤，不中断請求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}
