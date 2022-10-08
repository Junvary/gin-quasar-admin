package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"gorm.io/gorm"
)

type ServiceConfigBackend struct{}

func (s *ServiceConfigBackend) GetConfigBackendList(getConfigBackendList model.RequestGetConfigBackendList) (err error, config interface{}, total int64) {
	pageSize := getConfigBackendList.PageSize
	offset := getConfigBackendList.PageSize * (getConfigBackendList.Page - 1)
	db := global.GqaDb.Model(&model.SysConfigBackend{})
	var configList []model.SysConfigBackend
	//配置搜索
	if getConfigBackendList.ConfigItem != "" {
		db = db.Where("config_item like ?", "%"+getConfigBackendList.ConfigItem+"%")
	}
	if getConfigBackendList.Memo != "" {
		db = db.Where("memo like ?", "%"+getConfigBackendList.Memo+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(getConfigBackendList.SortBy, getConfigBackendList.Desc)).Find(&configList).Error
	return err, configList, total
}

func (s *ServiceConfigBackend) EditConfigBackend(toEditConfigBackend model.SysConfigBackend) (err error) {
	// 因为前台只传 custom 字段，这里允许编辑
	err = global.GqaDb.Save(&toEditConfigBackend).Error
	return err
}

func (s *ServiceConfigBackend) AddConfigBackend(toAddConfigBackend model.SysConfigBackend) (err error) {
	var configBackend model.SysConfigBackend
	if !errors.Is(global.GqaDb.Where("config_item = ?", toAddConfigBackend.ConfigItem).First(&configBackend).Error, gorm.ErrRecordNotFound) {
		return errors.New("此后台配置已存在：" + toAddConfigBackend.ConfigItem)
	}
	err = global.GqaDb.Create(&toAddConfigBackend).Error
	return err
}

func (s *ServiceConfigBackend) DeleteConfigBackendById(id uint) (err error) {
	var sysConfigBackend model.SysConfigBackend
	if err = global.GqaDb.Where("id = ?", id).First(&sysConfigBackend).Error; err != nil {
		return err
	}
	if sysConfigBackend.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + sysConfigBackend.ConfigItem)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysConfigBackend).Error
	return err
}
