package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
)

type ServiceMenu struct {
}

func (s *ServiceMenu) GetUserMenu(c *gin.Context) (err error, menu []system.SysMenu) {
	username := utils.GetUsername(c)
	var user system.SysUser
	err = global.GqaDb.Preload("Role").Where("username=?", username).First(&user).Error
	if err != nil {
		return err, nil
	}
	var role []system.SysRole
	err = global.GqaDb.Model(&user).Association("Role").Find(&role)
	if err != nil {
		return err, nil
	}
	var menus []system.SysMenu
	err = global.GqaDb.Model(&role).Association("Menu").Find(&menus)
	if err != nil {
		return err, nil
	}
	return nil, menus
}

func (s *ServiceMenu) GetMenuList(pageInfo system.RequestPage) (err error, menu interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysMenu{})
	var menuList []system.SysMenu
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&menuList).Error
	return err, menuList, total
}

func (s *ServiceMenu) EditMenu(menu system.SysMenu) (err error) {
	err = global.GqaDb.Updates(&menu).Error
	return err
}

func (s *ServiceMenu) AddMenu(m system.SysMenu) (err error) {
	err = global.GqaDb.Create(&m).Error
	return err
}

func (s *ServiceMenu) DeleteMenu(id uint) (err error) {
	var menu system.SysMenu
	err = global.GqaDb.Where("id = ?", id).Delete(&menu).Error
	return err
}

func (s *ServiceMenu) QueryMenuById(id uint) (err error, menuInfo system.SysMenu) {
	var menu system.SysMenu
	err = global.GqaDb.First(&menu, "id = ?", id).Error
	return err, menu
}
