package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceApi struct{}

func (s *ServiceApi) GetApiList(getApiList model.RequestGetApiList) (err error, api interface{}, total int64) {
	pageSize := getApiList.PageSize
	offset := getApiList.PageSize * (getApiList.Page - 1)
	db := global.GqaDb.Model(&model.SysApi{})
	var apiList []model.SysApi
	// Search
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

func (s *ServiceApi) EditApi(c *gin.Context, toEditApi model.SysApi) (err error) {
	if toEditApi.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n(c, "StableCantDo"))
	}
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var oldApi model.SysApi
		if err = tx.Where("id = ?", toEditApi.Id).First(&oldApi).Error; err != nil {
			return err
		}
		if err = tx.Save(&toEditApi).Error; err != nil {
			return err
		}
		var oldRoleApiList []model.SysRoleApi
		if err = tx.Where("api_group = ? and api_method = ? and api_path = ?", oldApi.ApiGroup, oldApi.ApiMethod, oldApi.ApiPath).
			Find(&oldRoleApiList).Error; err != nil {
			return err
		}
		if err = tx.Where("api_group = ? and api_method = ? and api_path = ?", oldApi.ApiGroup, oldApi.ApiMethod, oldApi.ApiPath).
			Delete(&model.SysRoleApi{}).Error; err != nil {
			return err
		}
		var newRoleApiList []model.SysRoleApi
		for _, oldRoleApi := range oldRoleApiList {
			newRoleApiList = append(newRoleApiList, model.SysRoleApi{
				RoleCode:  oldRoleApi.RoleCode,
				ApiGroup:  oldApi.ApiGroup,
				ApiMethod: oldApi.ApiMethod,
				ApiPath:   oldApi.ApiPath,
			})
		}
		if err = tx.Create(&newRoleApiList).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *ServiceApi) AddApi(toAddApi model.SysApi) (err error) {
	err = global.GqaDb.Create(&toAddApi).Error
	return err
}

func (s *ServiceApi) DeleteApiById(c *gin.Context, id uint) (err error) {
	var sysApi model.SysApi
	if err = global.GqaDb.Where("id = ?", id).First(&sysApi).Error; err != nil {
		return err
	}
	if sysApi.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n(c, "StableCantDo"))
	}
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("id = ?", id).Unscoped().Delete(&sysApi).Error; err != nil {
			return err
		}
		if err = tx.Where("api_group = ? and api_method = ? and api_path = ?", sysApi.ApiGroup, sysApi.ApiMethod, sysApi.ApiPath).
			Delete(&model.SysRoleApi{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *ServiceApi) QueryApiById(id uint) (err error, apiInfo model.SysApi) {
	var api model.SysApi
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&api, "id = ?", id).Error
	return err, api
}
