package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
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
	// Search
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
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestMenuList.SortBy, requestMenuList.Desc)).
		Preload("Button").Find(&menuList).Error
	menuTree := MenuList2MenuTree(menuList, "")
	return err, menuTree, total
}

func (s *ServiceMenu) EditMenu(toEditMenu model.SysMenu) (err error) {
	var sysMenu model.SysMenu
	if sysMenu.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n("StableCantDo") + toEditMenu.Title)
	}
	if err = global.GqaDb.Where("id = ?", toEditMenu.Id).First(&sysMenu).Error; err != nil {
		return err
	}
	//先删除关联button表中menu_name的记录
	var menuButton model.SysButton
	if err = global.GqaDb.Where("menu_name = ?", toEditMenu.Name).Delete(&menuButton).Error; err != nil {
		return err
	}
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
		return errors.New(utils.GqaI18n("StableCantDo") + sysMenu.Title)
	}
	if err = global.GqaDb.Where("id = ?", id).First(&sysMenu).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysMenu).Error
	return err
}

func (s *ServiceMenu) QueryMenuById(id uint) (err error, menuInfo model.SysMenu) {
	var menu model.SysMenu
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Button").First(&menu, "id = ?", id).Error
	return err, menu
}
