package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (s *ServiceMenu) EditMenu(c *gin.Context, toEditMenu model.SysMenu) (err error) {
	var sysMenu model.SysMenu
	if err = global.GqaDb.Where("id = ?", toEditMenu.Id).First(&sysMenu).Error; err != nil {
		return err
	}
	if sysMenu.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n(c, "StableCantDo") + toEditMenu.Title)
	}
	if sysMenu.Name != toEditMenu.Name {
		return errors.New(utils.GqaI18n(c, "EditFailed") + sysMenu.Name)
	}
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		//先删除关联button表中menu_name的记录
		var menuButton model.SysButton
		if err = tx.Where("menu_name = ?", toEditMenu.Name).Delete(&menuButton).Error; err != nil {
			return err
		}
		err = tx.Save(&toEditMenu).Error
		return err
	})
}

func (s *ServiceMenu) AddMenu(toAddMenu model.SysMenu) (err error) {
	err = global.GqaDb.Create(&toAddMenu).Error
	return err
}

func (s *ServiceMenu) DeleteMenuById(c *gin.Context, id uint) (err error) {
	var sysMenu model.SysMenu
	if err = global.GqaDb.Where("id = ?", id).First(&sysMenu).Error; err != nil {
		return err
	}
	if sysMenu.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n(c, "StableCantDo") + sysMenu.Title)
	}
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("id = ?", id).Unscoped().Delete(&sysMenu).Error; err != nil {
			return err
		}
		if err = tx.Where("sys_menu_name = ?", sysMenu.Name).Delete(&model.SysRoleMenu{}).Error; err != nil {
			return err
		}
		if err = tx.Where("menu_name = ?", sysMenu.Name).Delete(&model.SysButton{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *ServiceMenu) QueryMenuById(id uint) (err error, menuInfo model.SysMenu) {
	var menu model.SysMenu
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Button").First(&menu, "id = ?", id).Error
	return err, menu
}
