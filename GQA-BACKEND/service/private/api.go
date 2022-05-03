package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceApi struct{}

func (s *ServiceApi) GetApiList(getApiList model.RequestGetApiList) (err error, role interface{}, total int64) {
	pageSize := getApiList.PageSize
	offset := getApiList.PageSize * (getApiList.Page - 1)
	db := global.GqaDb.Model(&model.SysApi{})
	var apiList []model.SysApi
	//配置搜索
	if getApiList.ApiGroup != "" {
		db = db.Where("api_group like ?", "%"+getApiList.ApiGroup+"%")
	}
	if getApiList.ApiMethod != "" {
		db = db.Where("api_method like ?", "%"+getApiList.ApiMethod+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(getApiList.SortBy, getApiList.Desc)).Find(&apiList).Error
	return err, apiList, total
}

func (s *ServiceApi) EditApi(toEditApi model.SysApi) (err error) {
	// 因为前台只传 custom 字段，这里允许编辑
	err = global.GqaDb.Save(&toEditApi).Error
	return err
}

func (s *ServiceApi) AddApi(toAddApi model.SysApi) (err error) {
	err = global.GqaDb.Create(&toAddApi).Error
	return err
}

func (s *ServiceApi) DeleteApiById(id uint) (err error) {
	var sysApi model.SysApi
	if err = global.GqaDb.Where("id = ?", id).First(&sysApi).Error; err != nil {
		return err
	}
	if sysApi.Stable == "yes" {
		return errors.New("系统内置不允许删除")
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysApi).Error
	return err
}

func (s *ServiceApi) QueryApiById(id uint) (err error, roleInfo model.SysApi) {
	var role model.SysApi
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&role, "id = ?", id).Error
	return err, role
}
