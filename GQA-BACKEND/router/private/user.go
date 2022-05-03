package private

import "github.com/gin-gonic/gin"

type RouterUser struct{}

func (r *RouterUser) InitRouterRouterUser(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("user")
	{
		privateGroup.POST("get-user-menu", apiPrivate.ApiUser.GetUserMenu)
		privateGroup.POST("get-user-list", apiPrivate.ApiUser.GetUserList)
		privateGroup.POST("edit-user", apiPrivate.ApiUser.EditUser)
		privateGroup.POST("add-user", apiPrivate.ApiUser.AddUser)
		privateGroup.POST("delete-user-by-id", apiPrivate.ApiUser.DeleteUserById)
		privateGroup.POST("query-user-by-id", apiPrivate.ApiUser.QueryUserById)
		privateGroup.POST("reset-password", apiPrivate.ApiUser.ResetPassword)
		privateGroup.POST("change-password", apiPrivate.ApiUser.ChangePassword)
		privateGroup.POST("change-nickname", apiPrivate.ApiUser.ChangeNickname)
	}
	return privateGroup
}
