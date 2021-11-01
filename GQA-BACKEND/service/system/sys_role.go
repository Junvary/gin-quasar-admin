package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

type ServiceRole struct {
	
}

func (s *ServiceRole)GetRoleList(pageInfo global.RequestPage) (err error, role interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysRole{})
	var roleList []system.SysRole
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(pageInfo.SortBy, pageInfo.Desc)).Find(&roleList).Error
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

func (s *ServiceRole) GetRoleApiList(roleCode *system.RequestRoleCode) (err error, api [][]string) {
	api = global.GqaCasbin.GetFilteredPolicy(0, roleCode.RoleCode)
	return err, api
}

func (s *ServiceRole) EditRoleApi(roleApi *system.RequestRoleApiEdit) (err error) {
	_, _ = global.GqaCasbin.RemoveFilteredPolicy(0, roleApi.RoleCode)
	var newPolicy [][]string
	for _, v:= range roleApi.Policy{
		sc := system.SysCasbin{
			Ptype: "p",
			RoleCode: v.RoleCode,
			Path: v.Path,
			Method: v.Method,
		}
		newPolicy = append(newPolicy, []string{sc.RoleCode, sc.Path, sc.Method})
	}
	if len(newPolicy) != 0{
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

func (s *ServiceRole) RemoveRoleUser(toDeleteRoleUser *system.RequestRoleUser) (err error) {
	var roleUser system.SysUserRole
	err = global.GqaDb.Where("sys_role_role_code = ? and sys_user_id = ?", toDeleteRoleUser.RoleCode, toDeleteRoleUser.UserId).Delete(&roleUser).Error
	return err
}

func (s *ServiceRole) AddRoleUser(toAddRoleUser *system.RequestRoleUserAdd) (err error) {
	var roleUser []system.RequestRoleUser
	for _, r := range toAddRoleUser.UserId{
		ur := system.RequestRoleUser{
			UserId: r,
			RoleCode: toAddRoleUser.RoleCode,
		}
		roleUser = append(roleUser, ur)
	}
	err = global.GqaDb.Model(&system.SysUserRole{}).Save(&roleUser).Error
	return err
}

