package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceConfigFrontend struct{}

func (s *ServiceConfigFrontend) GetConfigFrontendList(getConfigFrontendList model.RequestGetConfigFrontendList) (err error, config interface{}, total int64) {
	pageSize := getConfigFrontendList.PageSize
	offset := getConfigFrontendList.PageSize * (getConfigFrontendList.Page - 1)
	db := global.GqaDb.Model(&model.SysConfigFrontend{})
	var configList []model.SysConfigFrontend
	// Search
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
	err = global.GqaDb.Save(&toEditConfigFrontend).Error
	return err
}

func (s *ServiceConfigFrontend) AddConfigFrontend(c *gin.Context, toAddConfigFrontend model.SysConfigFrontend) (err error) {
	var configFrontend model.SysConfigFrontend
	if !errors.Is(global.GqaDb.Where("config_item = ?", toAddConfigFrontend.ConfigItem).First(&configFrontend).Error, gorm.ErrRecordNotFound) {
		return errors.New(utils.GqaI18n(c, "AlreadyExist") + toAddConfigFrontend.ConfigItem)
	}
	err = global.GqaDb.Create(&toAddConfigFrontend).Error
	return err
}

func (s *ServiceConfigFrontend) DeleteConfigFrontendById(c *gin.Context, id uint) (err error) {
	var sysConfigFrontend model.SysConfigFrontend
	if err = global.GqaDb.Where("id = ?", id).First(&sysConfigFrontend).Error; err != nil {
		return err
	}
	if sysConfigFrontend.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n(c, "StableCantDo") + sysConfigFrontend.ConfigItem)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysConfigFrontend).Error
	return err
}
