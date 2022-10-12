package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceMenu struct{}

func MenuList2MenuTree(menuList []model.SysMenu, pCode string) []model.SysMenu {
	var menuTree []model.SysMenu
	for _, v := range menuList {
		if v.ParentCode == pCode {
			v.Children = MenuList2MenuTree(menuList, v.Name)
			menuTree = append(menuTree, v)
		}
	}
	return menuTree
}

func (s *ServiceMenu) GetMenuList(requestMenuList model.RequestGetMenuList) (err error, menu interface{}, total int64) {
	pageSize := requestMenuList.PageSize
	offset := requestMenuList.PageSize * (requestMenuList.Page - 1)
	db := global.GqaDb.Model(&model.SysMenu{})
	var menuList []model.SysMenu
	//配置搜索
	if requestMenuList.Path != "" {
		db = db.Where("path like ?", "%"+requestMenuList.Path+"%")
	}
	if requestMenuList.Title != "" {
		db = db.Where("title like ?", "%"+requestMenuList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestMenuList.SortBy, requestMenuList.Desc)).Find(&menuList).Error
	menuTree := MenuList2MenuTree(menuList, "")
	return err, menuTree, total
}

func (s *ServiceMenu) EditMenu(toEditMenu model.SysMenu) (err error) {
	var sysMenu model.SysMenu
	if sysMenu.Stable == "yesNo_yes" {
		return errors.New("系统内置不允许编辑：" + toEditMenu.Title)
	}
	if err = global.GqaDb.Where("id = ?", toEditMenu.Id).First(&sysMenu).Error; err != nil {
		return err
	}
	//err = global.GqaDb.Updates(&toEditMenu).Error
	err = global.GqaDb.Save(&toEditMenu).Error
	return err
}

func (s *ServiceMenu) AddMenu(toAddMenu model.SysMenu) (err error) {
	err = global.GqaDb.Create(&toAddMenu).Error
	return err
}

func (s *ServiceMenu) DeleteMenuById(id uint) (err error) {
	var sysMenu model.SysMenu
	if sysMenu.Stable == "yesNo_yes" {
		return errors.New("系统内置不允许删除：" + sysMenu.Title)
	}
	if err = global.GqaDb.Where("id = ?", id).First(&sysMenu).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysMenu).Error
	return err
}

func (s *ServiceMenu) QueryMenuById(id uint) (err error, menuInfo model.SysMenu) {
	var menu model.SysMenu
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&menu, "id = ?", id).Error
	return err, menu
}
