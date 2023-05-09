package private

import (
	"encoding/json"
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
)

type ServiceUser struct{}

func (s *ServiceUser) GetUserList(requestUserList model.RequestGetUserList) (err error, user interface{}, total int64) {
	pageSize := requestUserList.PageSize
	offset := requestUserList.PageSize * (requestUserList.Page - 1)
	var userList []model.SysUser
	db := global.GqaDb.Model(&model.SysUser{})
	// Search
	if !requestUserList.WithAdmin {
		db = db.Where("username != ?", "admin")
	}
	if requestUserList.Username != "" {
		db = db.Where("username like ?", "%"+requestUserList.Username+"%")
	}
	if requestUserList.RealName != "" {
		db = db.Where("real_name like ?", "%"+requestUserList.RealName+"%")
	}
	err = db.Count(&total).Error
	db = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestUserList.SortBy, requestUserList.Desc))
	if requestUserList.DeptCode != "" {
		dept := model.SysDept{
			DeptCode: requestUserList.DeptCode,
		}
		err = db.Model(&dept).Preload("Role").Preload("Dept").Association("Staff").Find(&userList)
		total = int64(len(userList))
		return err, userList, total
	} else {
		err = db.Preload("Role").Preload("Dept").Find(&userList).Error
		return err, userList, total
	}
}

func (s *ServiceUser) EditUser(toEditUser model.SysUser) (err error) {
	var sysUser model.SysUser
	if sysUser.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n("StableCantDo") + toEditUser.Username)
	}
	if err = global.GqaDb.Where("id = ?", toEditUser.Id).First(&sysUser).Error; err != nil {
		return err
	}
	toEditUser.Password = sysUser.Password
	//err = global.GqaDb.Updates(&toEditUser).Error
	err = global.GqaDb.Save(&toEditUser).Error
	return err
}

func (s *ServiceUser) AddUser(toAddUser *model.SysUser) (err error) {
	var user model.SysUser
	if !errors.Is(global.GqaDb.Where("username = ?", toAddUser.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New(utils.GqaI18n("AlreadyExist") + toAddUser.Username)
	}
	defaultPassword := utils.GetConfigBackend("defaultPassword")
	if defaultPassword == "" {
		toAddUser.Password = utils.EncodeMD5("123456")
		err = global.GqaDb.Create(&toAddUser).Error
		return errors.New("successWithNoDefaultPassword")
	} else {
		toAddUser.Password = utils.EncodeMD5(defaultPassword)
		err = global.GqaDb.Create(&toAddUser).Error
		return err
	}
}

func (s *ServiceUser) DeleteUserById(id uint) (err error) {
	var sysUser model.SysUser
	if sysUser.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n("StableCantDo") + sysUser.Username)
	}
	if err = global.GqaDb.Where("id = ?", id).First(&sysUser).Error; err != nil {
		return err
	}
	if err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysUser).Error; err != nil {
		return err
	}
	var sysDeptUser model.SysDeptUser
	if err = global.GqaDb.Where("sys_user_username = ?", sysUser.Username).Delete(&sysDeptUser).Error; err != nil {
		return err
	}
	var sysUserRole model.SysUserRole
	err = global.GqaDb.Where("sys_user_username = ?", sysUser.Username).Delete(&sysUserRole).Error
	return err
}

func (s *ServiceUser) GetUserByUsername(username string) (err error, userInfo model.SysUser) {
	var user model.SysUser
	err = global.GqaDb.First(&user, "username = ?", username).Error
	return err, user
}

func (s *ServiceUser) QueryUserById(id uint) (err error, userInfo model.SysUser) {
	var user model.SysUser
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").
		Preload("Role").Preload("Dept").First(&user, "id = ?", id).Error
	return err, user
}

func (s *ServiceUser) ResetPassword(id uint) (err error) {
	var sysUser model.SysUser
	if err = global.GqaDb.Where("id = ?", id).First(&sysUser).Error; err != nil {
		return err
	}
	defaultPassword := utils.GetConfigBackend("defaultPassword")
	var pwd string
	if defaultPassword == "" {
		pwd = utils.EncodeMD5("123456")
	} else {
		pwd = utils.EncodeMD5(defaultPassword)
	}
	err = global.GqaDb.Model(&sysUser).Update("password", pwd).Error
	return err
}

func (s *ServiceUser) GetUserMenu(c *gin.Context) (err error, defaultPageList []string, menu []model.SysMenu, buttons []string) {
	username := utils.GetUsername(c)
	var user model.SysUser
	err = global.GqaDb.Preload("Role").Where("username=?", username).First(&user).Error
	if err != nil {
		return err, nil, nil, nil
	}
	var role []model.SysRole
	err = global.GqaDb.Model(&user).Association("Role").Find(&role)
	if err != nil {
		return err, nil, nil, nil
	}
	//获取角色默认页面列表
	var roleCodeList []string
	for _, v := range role {
		defaultPageList = append(defaultPageList, v.DefaultPage)
		// 获取角色code列表为获取按钮列表做准备
		roleCodeList = append(roleCodeList, v.RoleCode)
	}

	var menus []model.SysMenu
	err = global.GqaDb.Model(&role).Preload("Button").Association("Menu").Find(&menus)
	if err != nil {
		return err, nil, nil, nil
	}
	//获取所有按钮权限
	var buttonList []model.SysRoleButton
	err = global.GqaDb.Where("sys_role_role_code in ?", roleCodeList).Find(&buttonList).Error
	var buttonListString []string
	for _, v := range buttonList {
		buttonListString = append(buttonListString, v.SysButtonButtonCode)
	}
	//按钮权限去重
	buttons = utils.RemoveDuplicateElementFromSlice(buttonListString)
	//menus切片去重
	type distinctMenu []model.SysMenu
	resultMenu := map[string]bool{}
	for _, v := range menus {
		data, _ := json.Marshal(v)
		resultMenu[string(data)] = true
	}
	result := distinctMenu{}
	for mm := range resultMenu {
		var m model.SysMenu
		_ = json.Unmarshal([]byte(mm), &m)
		result = append(result, m)
	}
	//result切片排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Sort < result[j].Sort
	})
	return nil, defaultPageList, result, buttons
}

func (s *ServiceUser) ChangePassword(username string, toChangePassword model.RequestChangePassword) (err error) {
	if toChangePassword.NewPassword1 != toChangePassword.NewPassword2 {
		return errors.New("两次新密码不一致！")
	}
	var sysUser model.SysUser
	if err = global.GqaDb.Where("username = ?", username).First(&sysUser).Error; err != nil {
		return err
	}
	oldPassword := utils.EncodeMD5(toChangePassword.OldPassword)
	if sysUser.Password != oldPassword {
		return errors.New("旧密码错误！")
	}
	newPassword := utils.EncodeMD5(toChangePassword.NewPassword1)
	if oldPassword == newPassword {
		return errors.New("新旧密码是一样的！")
	}
	err = global.GqaDb.Model(&sysUser).Update("password", newPassword).Error
	return err
}

func (s *ServiceUser) ChangeNickname(username string, toChangeNickname model.RequestChangeNickname) (err error) {
	var sysUser model.SysUser
	if err = global.GqaDb.Where("username = ?", username).First(&sysUser).Error; err != nil {
		return err
	}
	err = global.GqaDb.Model(&sysUser).Update("nickname", toChangeNickname.Nickname).Error
	return err
}
