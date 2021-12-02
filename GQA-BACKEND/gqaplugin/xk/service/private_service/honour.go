package private_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
)

func GetHonourList(getHonourList model.RequestHonourList) (err error, honour []model.GqaPluginXkHonour, total int64) {
	pageSize := getHonourList.PageSize
	offset := getHonourList.PageSize * (getHonourList.Page - 1)
	db := global.GqaDb.Model(&model.GqaPluginXkHonour{})
	var honourList []model.GqaPluginXkHonour
	//配置搜索
	if getHonourList.Title != ""{
		db = db.Where("title like ?", "%" + getHonourList.Title + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getHonourList.SortBy, getHonourList.Desc)).Preload("CreatedByUser").Find(&honourList).Error
	return err, honourList, total
}

func  EditHonour(toEditHonour model.GqaPluginXkHonour) (err error) {
	var honour model.GqaPluginXkHonour
	if err = global.GqaDb.Where("id = ?", toEditHonour.Id).First(&honour).Error; err != nil {
		return err
	}
	err = global.GqaDb.Updates(&toEditHonour).Error
	return err
}

func AddHonour(toAddHonour model.GqaPluginXkHonour) (err error) {
	err = global.GqaDb.Create(&toAddHonour).Error
	return err
}

func DeleteHonour(id uint) (err error) {
	var honour model.GqaPluginXkHonour
	if err = global.GqaDb.Where("id = ?", id).First(&honour).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&honour).Error
	return err
}

func QueryHonourById(id uint) (err error, honourInfo model.GqaPluginXkHonour) {
	var honour model.GqaPluginXkHonour
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&honour, "id = ?", id).Error
	return err, honour
}