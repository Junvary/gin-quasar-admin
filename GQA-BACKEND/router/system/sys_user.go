package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterUser struct {
}

func (r *RouterUser) InitRouterUser(Router *gin.RouterGroup) (R gin.IRoutes) {
	userGroup := Router.Group("user")
	apiUser := api.GroupApiApp.ApiSystem.ApiUser
	{
		// 获取用户列表
		userGroup.POST("user-list", apiUser.GetUserList)
		// 编辑用户信息
		userGroup.PUT("user-edit", apiUser.EditUser)
		// 新增用户
		userGroup.POST("user-add", apiUser.AddUser)
		// 删除用户
		userGroup.DELETE("user-delete", apiUser.DeleteUser)
		// 根据ID查找用户
		userGroup.POST("user-id", apiUser.QueryUserById)
		// 获取用户的菜单
		userGroup.GET("user-menu", apiUser.GetUserMenu)
		// 获取用户的角色列表
		userGroup.GET("user-role", apiUser.GetUserRole)
	}
	return Router
}
