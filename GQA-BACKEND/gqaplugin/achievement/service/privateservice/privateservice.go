package privateservice

import (
	"errors"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCategoryList(getCategoryList model.RequestGetCategoryList, username string) (err error, records []model.PluginAchievementCategory, total int64) {
	pageSize := getCategoryList.PageSize
	offset := getCategoryList.PageSize * (getCategoryList.Page - 1)
	var recordList []model.PluginAchievementCategory
	db := gqaGlobal.GqaDb.Model(&model.PluginAchievementCategory{})

	//配置搜索
	if getCategoryList.Category != "" {
		db = db.Where("category like ?", "%"+getCategoryList.Category+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return err, recordList, 0
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getCategoryList.SortBy, getCategoryList.Desc)).Preload("CreatedByUser").Find(&recordList).Error
	return err, recordList, total
}

func EditCategory(toEditCategory model.PluginAchievementCategory, username string) (err error) {
	db := gqaGlobal.GqaDb.Model(&model.PluginAchievementCategory{})
	var record model.PluginAchievementCategory
	if err = db.Where("id = ?", toEditCategory.Id).First(&record).Error; err != nil {
		return err
	}
	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEditCategory.GqaModel),
		map[string]interface{}{
			"Category": toEditCategory.Category,
			"Code":     toEditCategory.Code,
			"Name":     toEditCategory.Name,
		})).Error
	return err
}

func AddCategory(c *gin.Context, toAddCategory model.PluginAchievementCategory, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(c, username, gqaGlobal.GqaDb.Model(&model.PluginAchievementCategory{})); err != nil {
		return err
	}
	err = db.Create(&toAddCategory).Error
	return err
}

func DeleteCategoryById(c *gin.Context, id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(c, username, gqaGlobal.GqaDb.Model(&model.PluginAchievementCategory{})); err != nil {
		return err
	}
	var record model.PluginAchievementCategory
	if err = db.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&record).Error
	return err
}

func QueryCategoryById(c *gin.Context, id uint, username string) (err error, recordInfo model.PluginAchievementCategory) {
	var record model.PluginAchievementCategory
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(c, username, gqaGlobal.GqaDb.Model(&model.PluginAchievementCategory{})); err != nil {
		return err, record
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&record, "id = ?", id).Error
	return err, record
}

func GetObtainList(getObtainList model.RequestGetObtainList, username string) (err error, records []model.PluginAchievementObtain, total int64) {
	pageSize := getObtainList.PageSize
	offset := getObtainList.PageSize * (getObtainList.Page - 1)
	var recordList []model.PluginAchievementObtain
	db := gqaGlobal.GqaDb.Model(&model.PluginAchievementObtain{})

	//配置搜索
	if getObtainList.CategoryCode != "" {
		db = db.Where("category_code like ?", "%"+getObtainList.CategoryCode+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return err, recordList, 0
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getObtainList.SortBy, getObtainList.Desc)).Preload("CreatedByUser").Find(&recordList).Error
	return err, recordList, total
}

func ObtainFind(obtainFind model.ObtainFind) (find bool) {
	var obtainGet model.PluginAchievementObtain
	err := gqaGlobal.GqaDb.Where("category_code = ? and created_by = ?", obtainFind.CategoryCode, obtainFind.Username).First(&obtainGet).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return false
	} else {
		return true
	}
}

func ObtainAchievement(toObtain model.PluginAchievementObtain) (err error) {
	err = gqaGlobal.GqaDb.Create(&toObtain).Error
	return err
}
