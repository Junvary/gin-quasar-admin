package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"gorm.io/gorm"
)

type ServiceConfigFrontend struct{}

func (s *ServiceConfigFrontend) GetConfigFrontendList(getConfigFrontendList model.RequestGetConfigFrontendList) (err error, config interface{}, total int64) {
	pageSize := getConfigFrontendList.PageSize
	offset := getConfigFrontendList.PageSize * (getConfigFrontendList.Page - 1)
	db := global.GqaDb.Model(&model.SysConfigFrontend{})
	var configList []model.SysConfigFrontend
	//配置搜索
	if getConfigFrontendList.ConfigItem != "" {
		db = db.Where("config_item like ?", "%"+getConfigFrontendList.ConfigItem+"%")
	}
	if getConfigFrontendList.Memo != "" {
		db = db.Where("memo like ?", "%"+getConfigFrontendList.Memo+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(getConfigFrontendList.SortBy, getConfigFrontendList.Desc)).Find(&configList).Error
	return err, configList, total
}

func (s *ServiceConfigFrontend) EditConfigFrontend(toEditConfigFrontend model.SysConfigFrontend) (err error) {
	// 因为前台只传 custom 字段，这里允许编辑
	err = global.GqaDb.Save(&toEditConfigFrontend).Error
	return err
}

func (s *ServiceConfigFrontend) AddConfigFrontend(toAddConfigFrontend model.SysConfigFrontend) (err error) {
	var configFrontend model.SysConfigFrontend
	if !errors.Is(global.GqaDb.Where("config_item = ?", toAddConfigFrontend.ConfigItem).First(&configFrontend).Error, gorm.ErrRecordNotFound) {
		return errors.New("此网站前台配置已存在：" + toAddConfigFrontend.ConfigItem)
	}
	err = global.GqaDb.Create(&toAddConfigFrontend).Error
	return err
}

func (s *ServiceConfigFrontend) DeleteConfigFrontendById(id uint) (err error) {
	var sysConfigFrontend model.SysConfigFrontend
	if sysConfigFrontend.Stable == "yesNo_yes" {
		return errors.New("系统内置不允许删除：" + sysConfigFrontend.ConfigItem)
	}
	if err = global.GqaDb.Where("id = ?", id).First(&sysConfigFrontend).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysConfigFrontend).Error
	return err
}
