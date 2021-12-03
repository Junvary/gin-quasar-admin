package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiUser struct{}

func (a *ApiUser) GetUserList(c *gin.Context) {
	var requestUserList system.RequestUserList
	if err := c.ShouldBindJSON(&requestUserList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, userList, total := ServiceUser.GetUserList(requestUserList); err != nil {
		global.GqaLog.Error("获取用户列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取用户列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  userList,
			Page:     requestUserList.Page,
			PageSize: requestUserList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiUser) EditUser(c *gin.Context) {
	var toEditUser system.SysUser
	if err := c.ShouldBindJSON(&toEditUser); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if toEditUser.Username == "admin" && toEditUser.Status == "off" {
		global.ErrorMessage("超级管理不能被禁用！", c)
		return
	}
	toEditUser.UpdatedBy = utils.GetUsername(c)
	if err := ServiceUser.EditUser(toEditUser); err != nil {
		global.GqaLog.Error("编辑用户失败！", zap.Any("err", err))
		global.ErrorMessage("编辑用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑用户成功！", c)
	}
}

func (a *ApiUser) AddUser(c *gin.Context) {
	var toAddUser system.RequestAddUser
	if err := c.ShouldBindJSON(&toAddUser); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addUser := &system.SysUser{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddUser.Status,
			Sort:      toAddUser.Sort,
			Remark:    toAddUser.Remark,
		},
		Avatar:   toAddUser.Avatar,
		Username: toAddUser.Username,
		Nickname: toAddUser.Nickname,
		RealName: toAddUser.RealName,
		Gender:   toAddUser.Gender,
		Mobile:   toAddUser.Mobile,
		Email:    toAddUser.Email,
	}
	if err := ServiceUser.AddUser(addUser); err != nil {
		global.GqaLog.Error("添加用户失败！", zap.Any("err", err))
		global.ErrorMessage("添加用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加用户成功！", c)
	}
}

func (a *ApiUser) DeleteUser(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	currentUsername := utils.GetUsername(c)
	err, currentUser := ServiceUser.GetUserByUsername(currentUsername)
	if err != nil {
		global.GqaLog.Error("获取用户信息失败！", zap.Any("err", err))
		global.ErrorMessage("获取用户信息失败，"+err.Error(), c)
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
	if err := ServiceUser.DeleteUser(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除用户失败！", zap.Any("err", err))
		global.ErrorMessage("删除用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除用户成功！", c)
	}
}

func (a *ApiUser) QueryUserById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, user := ServiceUser.QueryUserById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找用户失败！", zap.Any("err", err))
		global.ErrorMessage("查找用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": user}, "查找用户成功！", c)
	}
}

func (a *ApiUser) GetUserMenu(c *gin.Context) {
	err, menu := ServiceUser.GetUserMenu(c)
	if err != nil {
		global.ErrorMessage("获取用户菜单失败，"+err.Error(), c)
	}
	global.SuccessMessageData(gin.H{"records": menu}, "获取用户菜单成功！", c)
}

func (a *ApiUser) GetUserRole(c *gin.Context) {
	err, role := ServiceUser.GetUserRole(c)
	if err != nil {
		global.ErrorMessage("获取用户角色失败，"+err.Error(), c)
	}
	global.SuccessMessageData(gin.H{"records": role}, "获取用户角色成功！", c)
}

func (a *ApiUser)ChangePassword(c *gin.Context)  {
	var toChangePassword system.RequestChangePassword
	if err := c.ShouldBindJSON(&toChangePassword); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	username := utils.GetUsername(c)
	if err := ServiceUser.ChangePassword(username, toChangePassword); err != nil {
		global.GqaLog.Error("修改密码失败！", zap.Any("err", err))
		global.ErrorMessage("修改密码失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("修改密码成功！", c)
	}
}
