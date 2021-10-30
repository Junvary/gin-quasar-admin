package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiRole struct {
}

func (a *ApiRole) GetRoleList(c *gin.Context) {
	var pageInfo system.RequestPage
	_ = c.ShouldBindJSON(&pageInfo)
	if err, roleList, total := service.GroupServiceApp.ServiceSystem.GetRoleList(pageInfo); err != nil {
		global.GqaLog.Error("获取角色列表失败：", zap.Any("err", err))
		global.ErrorMessage("获取角色列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  roleList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiRole) EditRole(c *gin.Context) {
	var toEditRole system.SysRole
	_ = c.ShouldBindJSON(&toEditRole)
	if err := service.GroupServiceApp.ServiceSystem.EditRole(toEditRole); err != nil {
		global.GqaLog.Error("编辑角色失败!", zap.Any("err", err))
		global.ErrorMessage("编辑角色失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑角色成功！", c)
	}
}

func (a *ApiRole) AddRole(c *gin.Context) {
	var toAddRole system.RequestAddRole
	_ = c.ShouldBindJSON(&toAddRole)
	addRole := &system.SysRole{
		RoleCode: toAddRole.RoleCode,
		RoleName: toAddRole.RoleName,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddRole(*addRole); err != nil {
		global.GqaLog.Error("添加角色失败！", zap.Any("err", err))
		global.ErrorMessage("添加角色失败！"+err.Error(), c)
	} else {
		global.SuccessMessage("添加角色成功！", c)
	}
}

func (a *ApiRole) DeleteRole(c *gin.Context) {
	var toDeleteId system.RequestDelete
	_ = c.ShouldBindJSON(&toDeleteId)
	if err := service.GroupServiceApp.ServiceSystem.DeleteRole(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除角色失败！", zap.Any("err", err))
		global.ErrorMessage("删除角色失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除角色成功！", c)
	}
}

func (a *ApiRole) QueryRoleById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	_ = c.ShouldBindJSON(&toQueryId)
	if err, role := service.GroupServiceApp.ServiceSystem.QueryRoleById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找角色失败！", zap.Any("err", err))
		global.ErrorMessage("查找角色失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": role}, "查找角色成功！", c)
	}
}

func (a *ApiRole) GetRoleMenuList(c *gin.Context) {
	var roleCode system.RequestRoleCode
	_ = c.ShouldBindJSON(&roleCode)
	if err, menuList := service.GroupServiceApp.ServiceSystem.GetRoleMenuList(&roleCode); err != nil {
		global.GqaLog.Error("获取角色菜单列表失败：", zap.Any("err", err))
		global.ErrorMessage("获取角色菜单列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": menuList}, c)
	}
}

func (a *ApiRole) EditRoleMenu(c *gin.Context) {
	var roleMenu system.RequestRoleMenuEdit
	_ = c.ShouldBindJSON(&roleMenu)
	if roleMenu.RoleCode == "super-admin" {
		global.ErrorMessage("超级管理员角色不允许编辑！", c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.EditRoleMenu(&roleMenu); err != nil {
		global.GqaLog.Error("编辑角色菜单失败：", zap.Any("err", err))
		global.ErrorMessage("编辑角色菜单失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑角色菜单成功！", c)
	}
}

func (a *ApiRole) GetRoleApiList(c *gin.Context) {
	var roleCode system.RequestRoleCode
	_ = c.ShouldBindJSON(&roleCode)
	if err, apiList := service.GroupServiceApp.ServiceSystem.GetRoleApiList(&roleCode); err != nil {
		global.GqaLog.Error("获取角色API列表失败：", zap.Any("err", err))
		global.SuccessMessage("获取角色API列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": apiList}, c)
	}
}

func (a *ApiRole) EditRoleApi(c *gin.Context) {
	var roleApi system.RequestRoleApiEdit
	_ = c.ShouldBindJSON(&roleApi)
	if roleApi.RoleCode == "super-admin" {
		global.ErrorMessage("超级管理员角色不允许编辑！", c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.EditRoleApi(&roleApi); err != nil {
		global.GqaLog.Error("编辑角色API失败：", zap.Any("err", err))
		global.ErrorMessage("编辑角色API失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑角色API成功！", c)
	}
}

func (a *ApiRole) QueryUserByRole(c *gin.Context) {
	var roleCode system.RequestRoleCode
	_ = c.ShouldBindJSON(&roleCode)
	if err, userList := service.GroupServiceApp.ServiceSystem.QueryUserByRole(&roleCode); err != nil {
		global.GqaLog.Error("查找用户失败：", zap.Any("err", err))
		global.ErrorMessage("查找用户失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": userList}, c)
	}
}

func (a *ApiRole) RemoveRoleUser(c *gin.Context) {
	var toDeleteRoleUser system.RequestRoleUser
	_ = c.ShouldBindJSON(&toDeleteRoleUser)
	if toDeleteRoleUser.UserId == 1 && toDeleteRoleUser.RoleCode == "super-admin" {
		global.ErrorMessage("抱歉，你不能把超级管理员从超级管理员组中移除！", c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.RemoveRoleUser(&toDeleteRoleUser); err != nil {
		global.GqaLog.Error("移除用户失败！", zap.Any("err", err))
		global.ErrorMessage("移除用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("移除用户成功！", c)
	}
}

func (a *ApiRole) AddRoleUser(c *gin.Context) {
	var toAddRoleUser system.RequestRoleUserAdd
	_ = c.ShouldBindJSON(&toAddRoleUser)
	if err := service.GroupServiceApp.ServiceSystem.AddRoleUser(&toAddRoleUser); err != nil {
		global.GqaLog.Error("添加角色失败！", zap.Any("err", err))
		global.ErrorMessage("添加角色失败！"+err.Error(), c)
	} else {
		global.SuccessMessage("添加角色成功！", c)
	}
}
