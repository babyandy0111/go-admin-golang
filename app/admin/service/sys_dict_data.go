package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
)

type SysDictData struct {
	service.Service
}

func (e *SysDictData) GetPage(c *dto.SysDictDataGetPageReq, list *[]models.SysDictData, count *int64) error {
	var err error
	var data models.SysDictData

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *SysDictData) Get(d *dto.SysDictDataGetReq, model *models.SysDictData) error {
	var err error
	var data models.SysDictData

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查無資料或無權限")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *SysDictData) Insert(c *dto.SysDictDataInsertReq) error {
	var err error
	var data = new(models.SysDictData)
	c.Generate(data)
	err = e.Orm.Create(data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

func (e *SysDictData) Update(c *dto.SysDictDataUpdateReq) error {
	var err error
	var model = models.SysDictData{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.Orm.Save(model)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("無權限更新該資料")

	}
	return nil
}

func (e *SysDictData) Remove(c *dto.SysDictDataDeleteReq) error {
	var err error
	var data models.SysDictData

	db := e.Orm.Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("無權限刪除該資料")
		return err
	}
	return nil
}

func (e *SysDictData) GetAll(c *dto.SysDictDataGetPageReq, list *[]models.SysDictData) error {
	var err error
	var data models.SysDictData

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
