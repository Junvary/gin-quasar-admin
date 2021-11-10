package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceApi struct {
}

func (s *ServiceApi) GetApiList(pageInfo global.RequestPage) (err error, api interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysApi{})
	var apiList []system.SysApi
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(pageInfo.SortBy, pageInfo.Desc)).Find(&apiList).Error
	return err, apiList, total
}

func (s *ServiceApi) EditApi(toEditApi system.SysApi) (err error) {
	var sysApi system.SysApi
	if err = global.GqaDb.Where("id = ?", toEditApi.Id).First(&sysApi).Error; err != nil {
		return err
	}
	if sysApi.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + sysApi.Path)
	}
	err = global.GqaDb.Updates(&sysApi).Error
	return err
}

func (s *ServiceApi) AddApi(toAddApi system.SysApi) (err error) {
	var api system.SysApi
	if !errors.Is(global.GqaDb.Where("path = ? and method = ?", toAddApi.Path, toAddApi.Method).First(&api).Error, gorm.ErrRecordNotFound) {
		return errors.New("此配置已存在：" + toAddApi.Path)
	}
	err = global.GqaDb.Create(&api).Error
	return err
}

func (s *ServiceApi) DeleteApi(id uint) (err error) {
	var sysApi system.SysApi
	if err = global.GqaDb.Where("id = ?", id).First(&sysApi).Error; err != nil {
		return err
	}
	if sysApi.Stable == "yes" {
		return errors.New("系统内置不允许删除！" + sysApi.Path)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysApi).Error
	return err
}
