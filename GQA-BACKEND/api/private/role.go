package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiRole struct{}

func (a *ApiRole) GetRoleList(c *gin.Context) {
	var requestRoleList model.RequestGetRoleList
	if err := model.RequestShouldBindJSON(c, &requestRoleList); err != nil {
		return
	}
	if err, roleList, total := servicePrivate.ServiceRole.GetRoleList(requestRoleList); err != nil {
		global.GqaLogger.Error("获取角色列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取角色列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  roleList,
			Page:     requestRoleList.Page,
			PageSize: requestRoleList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiRole) EditRole(c *gin.Context) {
	var toEditRole model.SysRole
	if err := model.RequestShouldBindJSON(c, &toEditRole); err != nil {
		return
	}
	toEditRole.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceRole.EditRole(toEditRole); err != nil {
		global.GqaLogger.Error("编辑角色失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑角色失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑角色成功！")
		model.ResponseSuccessMessage("编辑角色成功！", c)
	}
}

func (a *ApiRole) AddRole(c *gin.Context) {
	var toAddRole model.RequestAddRole
	if err := model.RequestShouldBindJSON(c, &toAddRole); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddRole.Status,
			Sort:      toAddRole.Sort,
			Memo:      toAddRole.Memo,
		},
	}
	addRole := &model.SysRole{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		RoleCode:                          toAddRole.RoleCode,
		RoleName:                          toAddRole.RoleName,
	}
	if err := servicePrivate.ServiceRole.AddRole(*addRole); err != nil {
		global.GqaLogger.Error("添加角色失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加角色失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加角色成功！", c)
	}
}

func (a *ApiRole) DeleteRoleById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.DeleteRoleById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除角色失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除角色失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除角色成功！")
		model.ResponseSuccessMessage("删除角色成功！", c)
	}
}

func (a *ApiRole) QueryRoleById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, role := servicePrivate.ServiceRole.QueryRoleById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找角色失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找角色失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": role}, "查找角色成功！", c)
	}
}

func (a *ApiRole) GetRoleMenuList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, menuList := servicePrivate.ServiceRole.GetRoleMenuList(&roleCode); err != nil {
		global.GqaLogger.Error("获取角色菜单列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取角色菜单列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": menuList}, c)
	}
}

func (a *ApiRole) EditRoleMenu(c *gin.Context) {
	var roleMenu model.RequestRoleMenuEdit
	if err := model.RequestShouldBindJSON(c, &roleMenu); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.EditRoleMenu(&roleMenu); err != nil {
		global.GqaLogger.Error("编辑角色菜单失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑角色菜单失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑角色菜单成功！")
		model.ResponseSuccessMessage("编辑角色菜单成功！", c)
	}
}

func (a *ApiRole) GetRoleApiList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, apiList := servicePrivate.ServiceRole.GetRoleApiList(&roleCode); err != nil {
		global.GqaLogger.Error("获取角色API列表失败！", zap.Any("err", err))
		model.ResponseSuccessMessage("获取角色API列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": apiList}, c)
	}
}

func (a *ApiRole) EditRoleApi(c *gin.Context) {
	var roleApi model.RequestEditRoleApi
	if err := model.RequestShouldBindJSON(c, &roleApi); err != nil {
		return
	}
	//if roleApi.RoleCode == "super-admin" {
	//	model.ResponseErrorMessage("超级管理员角色不允许编辑！", c)
	//	return
	//}
	if err := servicePrivate.ServiceRole.EditRoleApi(&roleApi); err != nil {
		global.GqaLogger.Error("编辑角色API失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑角色API失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑角色API成功！")
		model.ResponseSuccessMessage("编辑角色API成功！", c)
	}
}

func (a *ApiRole) QueryUserByRole(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, userList := servicePrivate.ServiceRole.QueryUserByRole(&roleCode); err != nil {
		global.GqaLogger.Error("查找角色用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找角色用户失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": userList}, c)
	}
}

func (a *ApiRole) RemoveRoleUser(c *gin.Context) {
	var toDeleteRoleUser model.RequestRoleUser
	if err := model.RequestShouldBindJSON(c, &toDeleteRoleUser); err != nil {
		return
	}
	if toDeleteRoleUser.Username == "admin" && toDeleteRoleUser.RoleCode == "super-admin" {
		model.ResponseErrorMessage("抱歉，你不能把超级管理员从超级管理员组中移除！", c)
		return
	}
	if err := servicePrivate.ServiceRole.RemoveRoleUser(&toDeleteRoleUser); err != nil {
		global.GqaLogger.Error("移除角色用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("移除角色用户失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "移除角色用户成功！")
		model.ResponseSuccessMessage("移除角色用户成功！", c)
	}
}

func (a *ApiRole) AddRoleUser(c *gin.Context) {
	var toAddRoleUser model.RequestRoleUserAdd
	if err := model.RequestShouldBindJSON(c, &toAddRoleUser); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.AddRoleUser(&toAddRoleUser); err != nil {
		global.GqaLogger.Error("添加角色用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加角色用户失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加角色用户成功！", c)
	}
}

func (a *ApiRole) EditRoleDeptDataPermission(c *gin.Context) {
	var toEditRoleDeptDataPermission model.RequestRoleDeptDataPermission
	if err := model.RequestShouldBindJSON(c, &toEditRoleDeptDataPermission); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.EditRoleDeptDataPermission(&toEditRoleDeptDataPermission); err != nil {
		global.GqaLogger.Error("编辑角色部门数据权限失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑角色部门数据权限失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑角色部门数据权限成功！")
		model.ResponseSuccessMessage("编辑角色部门数据权限成功！", c)
	}
}

func (a *ApiRole) GetRoleButtonList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, buttonList := servicePrivate.ServiceRole.GetRoleButtonList(&roleCode); err != nil {
		global.GqaLogger.Error("获取角色按钮列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取角色按钮列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": buttonList}, c)
	}
}
