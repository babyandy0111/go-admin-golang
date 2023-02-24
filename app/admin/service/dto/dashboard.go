package dto

import (
	"go-admin/common/dto"
)

type DashboardReq struct {
	dto.Pagination `search:"-"`
	Year           string `form:"year" comment:"年份"`
	Currency       string `form:"currency" comment:"幣別名稱"`
}
