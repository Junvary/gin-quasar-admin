package public_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
)

func GetHonour(getHonourList model.RequestHonourList) (err error, honour []model.GqaPluginXkHonour, total int64) {
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
