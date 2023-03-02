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

type Account struct {
	service.Service
}

// GetPage 查詢Account列表
func (e *Account) GetPage(c *dto.AccountGetPageReq, p *actions.DataPermission, list *[]models.Account, count *int64) error {
	var err error
	var data models.Account
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AccountService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 查詢Account單筆資料
func (e *Account) Get(d *dto.AccountGetReq, p *actions.DataPermission, model *models.Account) error {
	var data models.Account

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查無資料或無權限")
		e.Log.Errorf("Service GetAccount error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 新增Account資料
func (e *Account) Insert(c *dto.AccountInsertReq) error {
	var err error
	var data models.Account
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AccountService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 更新Account資料
func (e *Account) Update(c *dto.AccountUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Account{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AccountService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限更新資料")
	}
	return nil
}

// Remove 刪除Account資料
func (e *Account) Remove(d *dto.AccountDeleteReq, p *actions.DataPermission) error {
	var data models.Account

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAccount error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限刪除資料")
	}
	return nil
}
