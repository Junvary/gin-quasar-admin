package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	adapter "github.com/casbin/gorm-adapter/v3"
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

func (s *ServiceRole) GetRoleMenuList(roleCode *system.RequestRoleCode) (err error, menu []system.SysRoleMenu) {
	err = global.GqaDb.Where("sys_role_role_code=?", roleCode.RoleCode).Find(&menu).Error
	return err, menu
}

func (s *ServiceRole) EditRoleMenu(roleMenu *system.RequestRoleMenuEdit) (err error) {
	err = global.GqaDb.Where("sys_role_role_code=?", roleMenu.RoleCode).Delete(&system.SysRoleMenu{}).Error
	if err != nil{
		return err
	}
	if len(roleMenu.RoleMenu) != 0 {
		err = global.GqaDb.Model(&system.SysRoleMenu{}).Create(&roleMenu.RoleMenu).Error
		return err
	}
	return nil
}

func (s *ServiceRole) GetRoleApiList(roleCode *system.RequestRoleCode) (err error, api []adapter.CasbinRule) {
	err = global.GqaDb.Where("v0=?", roleCode.RoleCode).Find(&api).Error
	return err, api
}

func (s *ServiceRole) EditRoleApi(roleApi *system.RequestRoleApiEdit) (err error) {
	err = global.GqaDb.Where("v0=?", roleApi.RoleCode).Delete(&adapter.CasbinRule{}).Error
	if err != nil{
		return err
	}
	if len(roleApi.RoleApi) != 0 {
		err = global.GqaDb.Model(&adapter.CasbinRule{}).Create(&roleApi.RoleApi).Error
		return err
	}
	return nil
}
