package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiUser struct {
}

func (a *ApiUser) GetUserList(c *gin.Context) {
	var pageInfo system.RequestPage
	_ = c.ShouldBindJSON(&pageInfo)
	if err, userList, total := service.GroupServiceApp.ServiceSystem.GetUserList(pageInfo); err != nil {
		global.GqaLog.Error("获取用户列表失败：", zap.Any("err", err))
		global.ErrorMessage("获取用户列表失败，" + err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:     userList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiUser) EditUser(c *gin.Context) {
	var toEditUser system.SysUser
	_ = c.ShouldBindJSON(&toEditUser)
	if toEditUser.Username == "admin" && toEditUser.Status == "off"{
		global.ErrorMessage("超级管理不能被禁用！", c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.EditUser(toEditUser); err != nil {
		global.GqaLog.Error("编辑用户失败!", zap.Any("err", err))
		global.ErrorMessage("编辑用户失败，" + err.Error(), c)
	} else {
		global.SuccessMessage("编辑用户成功！", c)
	}
}

func (a *ApiUser) AddUser(c *gin.Context) {
	var toAddUser system.RequestAddUser
	_ = c.ShouldBindJSON(&toAddUser)
	addUser := &system.SysUser{
		Avatar:   toAddUser.Avatar,
		Username: toAddUser.Username,
		Nickname: toAddUser.Nickname,
		RealName: toAddUser.RealName,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddUser(*addUser); err != nil {
		global.GqaLog.Error("添加用户失败！", zap.Any("err", err))
		global.ErrorMessage("添加用户失败！"+err.Error(), c)
	} else {
		global.SuccessMessage("添加用户成功！", c)
	}
}

func (a *ApiUser) DeleteUser(c *gin.Context) {
	var toDeleteId system.RequestDelete
	_ = c.ShouldBindJSON(&toDeleteId)
	currentUsername := utils.GetUsername(c)
	err, currentUser := service.GroupServiceApp.ServiceSystem.GetUserByUsername(currentUsername)
	if err != nil {
		global.GqaLog.Error("获取用户信息失败！", zap.Any("err", err))
		global.ErrorMessage("获取用户信息失败，" + err.Error(), c)
		return
	}
	if currentUser.Id == toDeleteId.Id {
		global.ErrorMessage("你不能删除自己！", c)
		return
	}
	// 初始化时 admin 的 Id 为 1，这里就这样判断了，可以增加更多的逻辑。
	if toDeleteId.Id == 1 {
		global.ErrorMessage("超级管理员不能被删除！", c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.DeleteUser(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除用户失败！", zap.Any("err", err))
		global.ErrorMessage("删除用户失败，" + err.Error(), c)
	} else {
		global.SuccessMessage("删除用户成功！", c)
	}
}

func (a *ApiUser) QueryUserById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	_ = c.ShouldBindJSON(&toQueryId)
	if err, user := service.GroupServiceApp.ServiceSystem.QueryUserById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找用户失败！", zap.Any("err", err))
		global.ErrorMessage("查找用户失败，" + err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": user}, "查找用户成功！", c)
	}
}

func (a *ApiUser) GetUserMenu(c *gin.Context) {
	err, menu := service.GroupServiceApp.ServiceSystem.GetUserMenu(c)
	if err != nil {
		global.ErrorMessage("获取用户菜单失败，" + err.Error(), c)
	}
	global.SuccessData(system.ResponseMenu{
		Menu: menu,
	}, c)
}

func (a *ApiUser) GetUserRole(c *gin.Context) {
	err, role := service.GroupServiceApp.ServiceSystem.GetUserRole(c)
	if err != nil {
		global.ErrorMessage("获取用户角色失败，" + err.Error(), c)
	}
	global.SuccessData(system.ResponseRole{
		Role: role,
	}, c)
}
