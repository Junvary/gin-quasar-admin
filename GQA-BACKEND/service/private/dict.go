package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceDict struct{}

func DictList2DictTree(dictList []model.SysDict, pCode string) []model.SysDict {
	var dictTree []model.SysDict
	for _, v := range dictList {
		if v.ParentCode == pCode {
			v.Children = DictList2DictTree(dictList, v.DictCode)
			dictTree = append(dictTree, v)
		}
	}
	return dictTree
}

func (s *ServiceDict) GetDictList(requestDictList model.RequestGetDictList) (err error, dict interface{}, total int64, parentCode string) {
	pageSize := requestDictList.PageSize
	offset := requestDictList.PageSize * (requestDictList.Page - 1)
	var db *gorm.DB
	if requestDictList.ParentCode == "" {
		db = global.GqaDb.Find(&model.SysDict{})
	} else {
		db = global.GqaDb.Where("parent_code = ?", requestDictList.ParentCode).Find(&model.SysDict{})
	}
	var dictList []model.SysDict
	//配置搜索
	if requestDictList.DictCode != "" {
		db = db.Where("dict_code like ?", "%"+requestDictList.DictCode+"%")
	}
	if requestDictList.DictLabel != "" {
		db = db.Where("dict_label like ?", "%"+requestDictList.DictLabel+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestDictList.SortBy, requestDictList.Desc)).Find(&dictList).Error
	dictTree := DictList2DictTree(dictList, "")
	return err, dictTree, total, requestDictList.ParentCode
}

func (s *ServiceDict) EditDict(c *gin.Context, toEditDict model.SysDict) (err error) {
	var sysDict model.SysDict
	if err = global.GqaDb.Where("id = ?", toEditDict.Id).First(&sysDict).Error; err != nil {
		return err
	}
	if sysDict.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n(c, "StableCantDo") + toEditDict.DictCode)
	}
	//err = global.GqaDb.Updates(&toEditDict).Error
	err = global.GqaDb.Save(&toEditDict).Error
	return err
}

func (s *ServiceDict) AddDict(c *gin.Context, toAddDict model.SysDict) (err error) {
	var dict model.SysDict
	if !errors.Is(global.GqaDb.Where("dict_code = ?", toAddDict.DictCode).First(&dict).Error, gorm.ErrRecordNotFound) {
		return errors.New(utils.GqaI18n(c, "AlreadyExist") + toAddDict.DictCode)
	}
	err = global.GqaDb.Create(&toAddDict).Error
	return err
}

func (s *ServiceDict) DeleteDictById(c *gin.Context, id uint) (err error) {
	var dict model.SysDict
	if err = global.GqaDb.Where("id = ?", id).First(&dict).Error; err != nil {
		return err
	}
	if dict.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n(c, "StableCantDo") + dict.DictCode)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&dict).Error
	return err
}

func (s *ServiceDict) QueryDictById(id uint) (err error, dictInfo model.SysDict) {
	var dict model.SysDict
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&dict, "id = ?", id).Error
	return err, dict
}
