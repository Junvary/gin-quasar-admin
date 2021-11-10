package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceConfig struct {
}

func (s *ServiceConfig) GetConfigList(pageInfo global.RequestPage) (err error, role interface{}, total int64) {
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

func (s *ServiceConfig) EditConfig(toEditConfig system.SysConfig) (err error) {
	// 因为前台只传 custom 字段，这里允许编辑
	//var sysConfig system.SysConfig
	//if err = global.GqaDb.Where("id = ?", editConfig.Id).First(&sysConfig).Error; err != nil {
	//	return err
	//}
	//if sysConfig.Stable {
	//	return errors.New("系统内置不允许编辑：" + editConfig.GqaOption)
	//}
	err = global.GqaDb.Updates(&toEditConfig).Error
	return err
}

func (s *ServiceConfig) AddConfig(toAddConfig system.SysConfig) (err error) {
	var config system.SysConfig
	if !errors.Is(global.GqaDb.Where("gqa_option = ?", toAddConfig.GqaOption).First(&config).Error, gorm.ErrRecordNotFound) {
		return errors.New("此配置已存在：" + toAddConfig.GqaOption)
	}
	err = global.GqaDb.Create(&toAddConfig).Error
	return err
}

func (s *ServiceConfig) DeleteConfig(id uint) (err error) {
	var sysConfig system.SysConfig
	if err = global.GqaDb.Where("id = ?", id).First(&sysConfig).Error; err != nil {
		return err
	}
	if sysConfig.Stable == "yes" {
		return errors.New("系统内置不允许删除！" + sysConfig.GqaOption)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysConfig).Error
	return err
}
