package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiUser struct{}

func (a *ApiUser) GetUserList(c *gin.Context) {
	var requestUserList model.RequestGetUserList
	if err := model.RequestShouldBindJSON(c, &requestUserList); err != nil {
		return
	}
	if err, userList, total := servicePrivate.ServiceUser.GetUserList(requestUserList); err != nil {
		global.GqaLogger.Error("获取用户列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取用户列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  userList,
			Page:     requestUserList.Page,
			PageSize: requestUserList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiUser) EditUser(c *gin.Context) {
	var toEditUser model.SysUser
	if err := model.RequestShouldBindJSON(c, &toEditUser); err != nil {
		return
	}
	if toEditUser.Username == "admin" && toEditUser.Status == "off" {
		model.ResponseErrorMessage("超级管理不能被禁用！", c)
		return
	}
	toEditUser.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceUser.EditUser(toEditUser); err != nil {
		global.GqaLogger.Error("编辑用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑用户失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑用户成功！")
		model.ResponseSuccessMessage("编辑用户成功！", c)
	}
}

func (a *ApiUser) AddUser(c *gin.Context) {
	var toAddUser model.RequestAddUser
	if err := model.RequestShouldBindJSON(c, &toAddUser); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddUser.Status,
			Sort:      toAddUser.Sort,
			Memo:      toAddUser.Memo,
		},
	}
	addUser := &model.SysUser{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Avatar:                            toAddUser.Avatar,
		Username:                          toAddUser.Username,
		Nickname:                          toAddUser.Nickname,
		RealName:                          toAddUser.RealName,
		Gender:                            toAddUser.Gender,
		Mobile:                            toAddUser.Mobile,
		Email:                             toAddUser.Email,
	}
	if err := servicePrivate.ServiceUser.AddUser(addUser); err != nil {
		if err.Error() == "successWithNoDefaultPassword" {
			model.ResponseSuccessMessage("添加用户成功！但未找到配置的默认密码，此用户的密码设置为：123456", c)
		} else {
			global.GqaLogger.Error("添加用户失败！", zap.Any("err", err))
			model.ResponseErrorMessage("添加用户失败，"+err.Error(), c)
		}
	} else {
		model.ResponseSuccessMessage("添加用户成功！", c)
	}
}

func (a *ApiUser) DeleteUserById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	currentUsername := utils.GetUsername(c)
	err, currentUser := servicePrivate.ServiceUser.GetUserByUsername(currentUsername)
	if err != nil {
		global.GqaLogger.Error("获取用户信息失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取用户信息失败，"+err.Error(), c)
		return
	}
	if currentUser.Id == toDeleteId.Id {
		global.GqaLogger.Error(utils.GetUsername(c) + "你不能删除自己！")
		model.ResponseErrorMessage("你不能删除自己！", c)
		return
	}
	// 初始化时 admin 的 Id 为 1，这里就这样判断了，可以增加更多的逻辑。
	if toDeleteId.Id == 1 {
		global.GqaLogger.Error(utils.GetUsername(c) + "超级管理员不能被删除！")
		model.ResponseErrorMessage("超级管理员不能被删除！", c)
		return
	}
	if err := servicePrivate.ServiceUser.DeleteUserById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GetUsername(c)+"删除用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除用户失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除用户成功！")
		model.ResponseSuccessMessage("删除用户成功！", c)
	}
}

func (a *ApiUser) QueryUserById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, user := servicePrivate.ServiceUser.QueryUserById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找用户失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": user}, "查找用户成功！", c)
	}
}

func (a *ApiUser) ResetPassword(c *gin.Context) {
	var toResetPasswordId model.RequestQueryById
	if err := c.ShouldBindJSON(&toResetPasswordId); err != nil {
		global.GqaLogger.Error("模型绑定失败！", zap.Any("err", err))
		model.ResponseErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := servicePrivate.ServiceUser.ResetPassword(toResetPasswordId.Id); err != nil {
		global.GqaLogger.Error("重置用户密码失败！", zap.Any("err", err))
		model.ResponseErrorMessage("重置用户密码失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("重置用户密码成功！", c)
	}
}

func (a *ApiUser) GetUserMenu(c *gin.Context) {
	err, defaultPageList, menu, buttons := servicePrivate.ServiceUser.GetUserMenu(c)
	if err != nil {
		model.ResponseErrorMessage("获取用户菜单失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": menu, "buttons": buttons, "default_page_list": defaultPageList}, "获取用户菜单成功！", c)
	}
}

func (a *ApiUser) ChangePassword(c *gin.Context) {
	var toChangePassword model.RequestChangePassword
	if err := model.RequestShouldBindJSON(c, &toChangePassword); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err := servicePrivate.ServiceUser.ChangePassword(username, toChangePassword); err != nil {
		global.GqaLogger.Error("修改密码失败！", zap.Any("err", err))
		model.ResponseErrorMessage("修改密码失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "修改密码成功！")
		model.ResponseSuccessMessage("修改密码成功！", c)
	}
}

func (a *ApiUser) ChangeNickname(c *gin.Context) {
	var toChangeNickname model.RequestChangeNickname
	if err := model.RequestShouldBindJSON(c, &toChangeNickname); err != nil {
		return
	}
	if toChangeNickname.Nickname == "" {
		global.GqaLogger.Error("修改昵称失败！不能为空！")
		model.ResponseErrorMessage("修改昵称失败！不能为空！", c)
	}
	username := utils.GetUsername(c)
	if err := servicePrivate.ServiceUser.ChangeNickname(username, toChangeNickname); err != nil {
		global.GqaLogger.Error("修改昵称失败！", zap.Any("err", err))
		model.ResponseErrorMessage("修改昵称失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "修改昵称成功！")
		model.ResponseSuccessMessage("修改昵称成功！", c)
	}
}
