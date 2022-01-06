package system

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"gorm.io/gorm"
)

type ServiceApi struct {}

func (s *ServiceApi) GetApiList(requestApiList system.RequestApiList) (err error, api interface{}, total int64) {
	pageSize := requestApiList.PageSize
	offset := requestApiList.PageSize * (requestApiList.Page - 1)
	db := global.GqaDb.Model(&system.SysApi{})
	var apiList []system.SysApi
	//配置搜索
	if requestApiList.ApiGroup != "" {
		db = db.Where("api_group like ?", "%"+requestApiList.ApiGroup+"%")
	}
	if requestApiList.ApiPath != "" {
		db = db.Where("api_path like ?", "%"+requestApiList.ApiPath+"%")
	}
	if requestApiList.ApiMethod != "" {
		db = db.Where("api_method like ?", "%"+requestApiList.ApiMethod+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(requestApiList.SortBy, requestApiList.Desc)).Find(&apiList).Error
	return err, apiList, total
}

func (s *ServiceApi) EditApi(toEditApi system.SysApi) (err error) {
	var sysApi system.SysApi
	if err = global.GqaDb.Where("id = ?", toEditApi.Id).First(&sysApi).Error; err != nil {
		return err
	}
	if sysApi.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + sysApi.ApiPath)
	}
	//err = global.GqaDb.Updates(&toEditApi).Error
	err = global.GqaDb.Save(&toEditApi).Error
	return err
}

func (s *ServiceApi) AddApi(toAddApi system.SysApi) (err error) {
	var api system.SysApi
	if !errors.Is(global.GqaDb.Where("api_path = ? and api_method = ?", toAddApi.ApiPath, toAddApi.ApiMethod).First(&api).Error, gorm.ErrRecordNotFound) {
		return errors.New("此配置已存在：" + toAddApi.ApiPath)
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
		return errors.New("系统内置不允许删除：" + sysApi.ApiPath)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysApi).Error
	return err
}
