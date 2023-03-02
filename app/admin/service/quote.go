package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Quote struct {
	service.Service
}

// GetPage 查詢Quote列表
func (e *Quote) GetPage(c *dto.QuoteGetPageReq, p *actions.DataPermission, list *[]models.Quote, count *int64) error {
	var err error
	var data models.Quote

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("QuoteService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 查詢Quote單筆資料
func (e *Quote) Get(d *dto.QuoteGetReq, p *actions.DataPermission, model *models.Quote) error {
	var data models.Quote

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查無資料或無權限")
		e.Log.Errorf("Service GetQuote error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 新增Quote資料
func (e *Quote) Insert(c *dto.QuoteInsertReq) error {
	var err error
	var data models.Quote
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("QuoteService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 更新Quote資料
func (e *Quote) Update(c *dto.QuoteUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Quote{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("QuoteService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限更新資料")
	}
	return nil
}

// Remove 刪除Quote資料
func (e *Quote) Remove(d *dto.QuoteDeleteReq, p *actions.DataPermission) error {
	var data models.Quote

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveQuote error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限刪除資料")
	}
	return nil
}
