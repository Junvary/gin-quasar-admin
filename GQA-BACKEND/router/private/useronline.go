package private

import "github.com/gin-gonic/gin"

type RouterUserOnline struct{}

func (r *RouterUserOnline) InitRouterRouterUserOnline(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("user-online")
	{
		privateGroup.POST("get-user-online-list", apiPrivate.ApiUserOnline.GetUserOnlineList)
		privateGroup.POST("kick-online-user", apiPrivate.ApiUserOnline.KickUserOnlineByUsername)
	}
	return privateGroup
}
