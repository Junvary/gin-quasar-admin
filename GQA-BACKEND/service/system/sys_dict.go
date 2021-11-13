package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceDict struct {
}

func (s *ServiceDict) GetDictList(requestDictList system.RequestDictList) (err error, role interface{}, total int64, parentId uint) {
	pageSize := requestDictList.PageSize
	offset := requestDictList.PageSize * (requestDictList.Page - 1)
	db := global.GqaDb.Where("parent_id = ?", requestDictList.ParentId).Find(&system.SysDict{})
	var dictList []system.SysDict
	//配置搜索
	if requestDictList.Value != ""{
		db = db.Where("value like ?", "%" + requestDictList.Value + "%")
	}
	if requestDictList.Label != ""{
		db = db.Where("label like ?", "%" + requestDictList.Label + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(requestDictList.SortBy, requestDictList.Desc)).Find(&dictList).Error
	return err, dictList, total, requestDictList.ParentId
}

func (s *ServiceDict) EditDict(toEditDict system.SysDict) (err error) {
	var sysDict system.SysDict
	if err = global.GqaDb.Where("id = ?", toEditDict.Id).First(&sysDict).Error; err != nil {
		return err
	}
	if sysDict.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + toEditDict.Value)
	}
	err = global.GqaDb.Updates(&toEditDict).Error
	return err
}

func (s *ServiceDict) AddDict(toAddDict system.SysDict) (err error) {
	var dict system.SysDict
	if !errors.Is(global.GqaDb.Where("value = ?", toAddDict.Value).First(&dict).Error, gorm.ErrRecordNotFound) {
		return errors.New("此字典已存在：" + toAddDict.Value)
	}
	err = global.GqaDb.Create(&toAddDict).Error
	return err
}

func (s *ServiceDict) DeleteDict(id uint) (err error) {
	var dict system.SysDict
	if err = global.GqaDb.Where("id = ?", id).First(&dict).Error; err != nil {
		return err
	}
	if dict.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + dict.Value)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&dict).Error
	return err
}

func (s *ServiceDict) QueryDictById(id uint) (err error, dictInfo system.SysDict) {
	var dict system.SysDict
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&dict, "id = ?", id).Error
	return err, dict
}
