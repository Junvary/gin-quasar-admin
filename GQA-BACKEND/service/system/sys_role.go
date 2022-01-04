package system

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"gorm.io/gorm"
)

type ServiceRole struct {
}

func (s *ServiceRole) GetRoleList(requestRoleList system.RequestRoleList) (err error, role interface{}, total int64) {
	pageSize := requestRoleList.PageSize
	offset := requestRoleList.PageSize * (requestRoleList.Page - 1)
	db := global.GqaDb.Model(&system.SysRole{})
	var roleList []system.SysRole
	//配置搜索
	if requestRoleList.RoleCode != ""{
		db = db.Where("role_code like ?", "%" + requestRoleList.RoleCode + "%")
	}
	if requestRoleList.RoleName != ""{
		db = db.Where("role_name like ?", "%" + requestRoleList.RoleName + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(requestRoleList.SortBy, requestRoleList.Desc)).Find(&roleList).Error
	return err, roleList, total
}

func (s *ServiceRole) EditRole(toEditRole system.SysRole) (err error) {
	var sysRole system.SysRole
	if err = global.GqaDb.Where("id = ?", toEditRole.Id).First(&sysRole).Error; err != nil {
		return err
	}
	if sysRole.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + toEditRole.RoleCode)
	}
	err = global.GqaDb.Updates(&toEditRole).Error
	return err
}

func (s *ServiceRole) AddRole(toAddRole system.SysRole) (err error) {
	var role system.SysRole
	if !errors.Is(global.GqaDb.Where("role_code = ?", toAddRole.RoleCode).First(&role).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色已存在：" + toAddRole.RoleCode)
	}
	err = global.GqaDb.Create(&toAddRole).Error
	return err
}

func (s *ServiceRole) DeleteRole(id uint) (err error) {
	var sysRole system.SysRole
	if err = global.GqaDb.Where("id = ?", id).First(&sysRole).Error; err != nil {
		return err
	}
	if sysRole.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + sysRole.RoleCode)
	}
	roleCode := sysRole.RoleCode
	// 删除 casbin_rule 表的权限
	_, _ = global.GqaCasbin.RemoveFilteredPolicy(0, roleCode)
	// 删除 sys_role 表的数据
	err = global.GqaDb.Unscoped().Delete(&sysRole).Error
	if err != nil {
		return err
	}
	// 删除 sys_user_role 表的对应关系
	err = global.GqaDb.Where("sys_role_role_code = ?", roleCode).Delete(&system.SysUserRole{}).Error
	if err != nil {
		return err
	}
	// 删除 sys_role_menu 表的对应关系
	err = global.GqaDb.Where("sys_role_role_code = ?", roleCode).Delete(&system.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	return err
}

func (s *ServiceRole) QueryRoleById(id uint) (err error, roleInfo system.SysRole) {
	var role system.SysRole
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&role, "id = ?", id).Error
	return err, role
}

func (s *ServiceRole) GetRoleMenuList(roleCode *system.RequestRoleCode) (err error, menu []system.SysRoleMenu) {
	err = global.GqaDb.Where("sys_role_role_code=?", roleCode.RoleCode).Find(&menu).Error
	return err, menu
}

func (s *ServiceRole) EditRoleMenu(toEditRoleMenu *system.RequestRoleMenuEdit) (err error) {
	err = global.GqaDb.Where("sys_role_role_code=?", toEditRoleMenu.RoleCode).Delete(&system.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	if len(toEditRoleMenu.RoleMenu) != 0 {
		err = global.GqaDb.Model(&system.SysRoleMenu{}).Create(&toEditRoleMenu.RoleMenu).Error
		return err
	}
	return nil
}

func (s *ServiceRole) GetRoleApiList(roleCode *system.RequestRoleCode) (err error, api [][]string) {
	api = global.GqaCasbin.GetFilteredPolicy(0, roleCode.RoleCode)
	return err, api
}

func (s *ServiceRole) EditRoleApi(toEditRoleApi *system.RequestRoleApiEdit) (err error) {
	_, _ = global.GqaCasbin.RemoveFilteredPolicy(0, toEditRoleApi.RoleCode)
	var newPolicy [][]string
	for _, v := range toEditRoleApi.Policy {
		sc := system.SysCasbin{
			Ptype: "p",
			V0:    toEditRoleApi.RoleCode,
			V1:    v.V1,
			V2:    v.V2,
		}
		newPolicy = append(newPolicy, []string{sc.V0, sc.V1, sc.V2})
	}
	if len(newPolicy) != 0 {
		_, _ = global.GqaCasbin.AddPolicies(newPolicy)
	}
	return nil
}

func (s *ServiceRole) QueryUserByRole(roleCode *system.RequestRoleCode) (err error, user []system.SysUser) {
	role := system.SysRole{
		RoleCode: roleCode.RoleCode,
	}
	var userList []system.SysUser
	err = global.GqaDb.Model(&role).Association("User").Find(&userList)
	return err, userList
}

func (s *ServiceRole) RemoveRoleUser(toRemoveRoleUser *system.RequestRoleUser) (err error) {
	var roleUser system.SysUserRole
	if toRemoveRoleUser.Username == "admin" && toRemoveRoleUser.RoleCode == "super-admin"{
		return errors.New("抱歉，你不能把超级管理员从超级管理员组中移除！")
	}
	err = global.GqaDb.Where("sys_role_role_code = ? and sys_user_username = ?", toRemoveRoleUser.RoleCode, toRemoveRoleUser.Username).Delete(&roleUser).Error
	return err
}

func (s *ServiceRole) AddRoleUser(toAddRoleUser *system.RequestRoleUserAdd) (err error) {
	var roleUser []system.RequestRoleUser
	for _, r := range toAddRoleUser.Username {
		ur := system.RequestRoleUser{
			Username:   r,
			RoleCode: toAddRoleUser.RoleCode,
		}
		roleUser = append(roleUser, ur)
	}
	if len(roleUser) != 0 {
		err = global.GqaDb.Model(&system.SysUserRole{}).Save(&roleUser).Error
		return err
	}else{
		return errors.New("本次操作没有影响！")
	}
}

func (s *ServiceRole) EditRoleDeptDataPermission(toEditRoleDeptDataPermission *system.RequestRoleDeptDataPermission) (err error) {
	var sysRole system.SysRole
	if err = global.GqaDb.Where("role_code = ?", toEditRoleDeptDataPermission.RoleCode).First(&sysRole).Error; err != nil {
		return err
	}
	if sysRole.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + toEditRoleDeptDataPermission.RoleCode)
	}
	sysRole.DeptDataPermissionType = toEditRoleDeptDataPermission.DeptDataPermissionType
	sysRole.DeptDataPermissionCustom = toEditRoleDeptDataPermission.DeptDataPermissionCustom
	err = global.GqaDb.Save(&sysRole).Error
	return err
}
