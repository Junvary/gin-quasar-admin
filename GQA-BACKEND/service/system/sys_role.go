package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceRole struct {
	
}

func (s *ServiceRole)GetRoleList(pageInfo system.RequestPage) (err error, role interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysRole{})
	var roleList []system.SysRole
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&roleList).Error
	return err, roleList, total
}

func (s *ServiceRole) EditRole(role system.SysRole) (err error) {
	err = global.GqaDb.Updates(&role).Error
	return err
}

func (s *ServiceRole) AddRole(r system.SysRole) (err error) {
	var role system.SysRole
	if !errors.Is(global.GqaDb.Where("role_code = ?", r.RoleCode).First(&role).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色已存在：" + r.RoleCode)
	}
	err = global.GqaDb.Create(&r).Error
	return err
}

func (s *ServiceRole) DeleteRole(id uint) (err error) {
	var role system.SysRole
	err = global.GqaDb.Where("id = ?", id).Delete(&role).Error
	return err
}

func (s *ServiceRole) QueryRoleById(id uint) (err error, roleInfo system.SysRole) {
	var role system.SysRole
	err = global.GqaDb.First(&role, "id = ?", id).Error
	return err, role
}
