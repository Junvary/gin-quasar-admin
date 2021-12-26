package system

import (
	"github.com/gin-gonic/gin"
)

type RouterUser struct {}

func (r *RouterUser) InitRouterUser(Router *gin.RouterGroup) (R gin.IRoutes) {
	userGroup := Router.Group("user")/*.Use(middleware.LogOperationHandler())*/
	{
		//获取用户列表
		userGroup.POST("user-list", ApiUser.GetUserList)
		//编辑用户信息
		userGroup.PUT("user-edit", ApiUser.EditUser)
		//新增用户
		userGroup.POST("user-add", ApiUser.AddUser)
		//删除用户
		userGroup.DELETE("user-delete", ApiUser.DeleteUser)
		//根据ID查找用户
		userGroup.POST("user-id", ApiUser.QueryUserById)
		//获取用户的菜单
		userGroup.GET("user-menu", ApiUser.GetUserMenu)
		//用户修改密码
		userGroup.POST("user-change-password", ApiUser.ChangePassword)
		//用户修改昵称
		userGroup.POST("user-change-nickname", ApiUser.ChangeNickname)
	}
	return Router
}
