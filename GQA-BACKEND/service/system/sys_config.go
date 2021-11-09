package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceConfig struct {
}

func (s *ServiceConfig)GetConfigList(pageInfo global.RequestPage) (err error, role interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysConfig{})
	var configList []system.SysConfig
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(pageInfo.SortBy, pageInfo.Desc)).Find(&configList).Error
	return err, configList, total
}

func (s *ServiceConfig) EditConfig(config system.SysConfig) (err error) {
	err = global.GqaDb.Updates(&config).Error
	return err
}

func (s *ServiceConfig) AddConfig(d system.SysConfig) (err error) {
	var config system.SysConfig
	if !errors.Is(global.GqaDb.Where("gqa_option = ?", d.GqaOption).First(&config).Error, gorm.ErrRecordNotFound) {
		return errors.New("此配置已存在：" + d.GqaOption)
	}
	err = global.GqaDb.Create(&d).Error
	return err
}