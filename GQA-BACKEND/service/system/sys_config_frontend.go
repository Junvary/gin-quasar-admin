package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceConfigFrontend struct {}

func (s *ServiceConfigFrontend) GetConfigFrontendList(getConfigFrontendList system.RequestConfigFrontendList) (err error, role interface{}, total int64) {
	pageSize := getConfigFrontendList.PageSize
	offset := getConfigFrontendList.PageSize * (getConfigFrontendList.Page - 1)
	db := global.GqaDb.Model(&system.SysConfigFrontend{})
	var configList []system.SysConfigFrontend
	//配置搜索
	if getConfigFrontendList.GqaOption != ""{
		db = db.Where("gqa_option like ?", "%" + getConfigFrontendList.GqaOption + "%")
	}
	if getConfigFrontendList.Remark != ""{
		db = db.Where("remark like ?", "%" + getConfigFrontendList.Remark + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getConfigFrontendList.SortBy, getConfigFrontendList.Desc)).Find(&configList).Error
	return err, configList, total
}

func (s *ServiceConfigFrontend) EditConfigFrontend(toEditConfigFrontend system.SysConfigFrontend) (err error) {
	// 因为前台只传 custom 字段，这里允许编辑
	err = global.GqaDb.Save(&toEditConfigFrontend).Error
	return err
}

func (s *ServiceConfigFrontend) AddConfigFrontend(toAddConfigFrontend system.SysConfigFrontend) (err error) {
	var configFrontend system.SysConfigFrontend
	if !errors.Is(global.GqaDb.Where("gqa_option = ?", toAddConfigFrontend.GqaOption).First(&configFrontend).Error, gorm.ErrRecordNotFound) {
		return errors.New("此前台配置已存在：" + toAddConfigFrontend.GqaOption)
	}
	err = global.GqaDb.Create(&toAddConfigFrontend).Error
	return err
}

func (s *ServiceConfigFrontend) DeleteConfigFrontend(id uint) (err error) {
	var sysConfigFrontend system.SysConfigFrontend
	if err = global.GqaDb.Where("id = ?", id).First(&sysConfigFrontend).Error; err != nil {
		return err
	}
	if sysConfigFrontend.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + sysConfigFrontend.GqaOption)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysConfigFrontend).Error
	return err
}
