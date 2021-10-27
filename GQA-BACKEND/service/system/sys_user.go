package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceUser struct {
}

func (s *ServiceUser) GetUserList(pageInfo system.RequestPage) (err error, user interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&userList).Error
	return err, userList, total
}

func (s *ServiceUser) EditUser(user system.SysUser) (err error) {
	err = global.GqaDb.Updates(&user).Error
	return err
}

func (s *ServiceUser) AddUser(u system.SysUser) (err error) {
	var user system.SysUser
	if !errors.Is(global.GqaDb.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("此用户已存在：" + u.Username)
	}
	u.Uuid =uuid.New()
	u.Password = utils.Md5([]byte(global.GqaConfig.System.DefaultPassword))
	err = global.GqaDb.Create(&u).Error
	return err
}

func (s *ServiceUser) DeleteUser(id uint) (err error) {
	var user system.SysUser
	err = global.GqaDb.Where("id = ?", id).Delete(&user).Error
	return err
}

func (s *ServiceUser) GetUserByUsername(username string) (err error, userInfo system.SysUser) {
	var user system.SysUser
	err = global.GqaDb.First(&user, "username = ?", username).Error
	return err, user
}

func (s *ServiceUser) QueryUserById(id uint) (err error, userInfo system.SysUser) {
	var user system.SysUser
	err = global.GqaDb.First(&user, "id = ?", id).Error
	return err, user
}

func (s *ServiceUser) GetUserMenu(c *gin.Context) (err error, menu []system.SysMenu) {
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
