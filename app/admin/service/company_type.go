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

type CompanyType struct {
	service.Service
}

// GetPage 查詢CompanyType列表
func (e *CompanyType) GetPage(c *dto.CompanyTypeGetPageReq, p *actions.DataPermission, list *[]models.CompanyType, count *int64) error {
	var err error
	var data models.CompanyType

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("CompanyTypeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 查詢CompanyType單筆資料
func (e *CompanyType) Get(d *dto.CompanyTypeGetReq, p *actions.DataPermission, model *models.CompanyType) error {
	var data models.CompanyType

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查無資料或無權限")
		e.Log.Errorf("Service GetCompanyType error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 新增CompanyType資料
func (e *CompanyType) Insert(c *dto.CompanyTypeInsertReq) error {
	var err error
	var data models.CompanyType
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CompanyTypeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 更新CompanyType資料
func (e *CompanyType) Update(c *dto.CompanyTypeUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.CompanyType{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("CompanyTypeService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限更新資料")
	}
	return nil
}

// Remove 刪除CompanyType資料
func (e *CompanyType) Remove(d *dto.CompanyTypeDeleteReq, p *actions.DataPermission) error {
	var data models.CompanyType

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveCompanyType error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限刪除資料")
	}
	return nil
}
