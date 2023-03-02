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

type Currency struct {
	service.Service
}

// GetPage 查詢Currency列表
func (e *Currency) GetPage(c *dto.CurrencyGetPageReq, p *actions.DataPermission, list *[]models.Currency, count *int64) error {
	var err error
	var data models.Currency

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("CurrencyService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 查詢Currency單筆資料
func (e *Currency) Get(d *dto.CurrencyGetReq, p *actions.DataPermission, model *models.Currency) error {
	var data models.Currency

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查無資料或無權限")
		e.Log.Errorf("Service GetCurrency error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 新增Currency資料
func (e *Currency) Insert(c *dto.CurrencyInsertReq) error {
	var err error
	var data models.Currency
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CurrencyService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 更新Currency資料
func (e *Currency) Update(c *dto.CurrencyUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Currency{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("CurrencyService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限更新資料")
	}
	return nil
}

// Remove 刪除Currency資料
func (e *Currency) Remove(d *dto.CurrencyDeleteReq, p *actions.DataPermission) error {
	var data models.Currency

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveCurrency error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限刪除資料")
	}
	return nil
}
