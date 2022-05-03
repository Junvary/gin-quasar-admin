package public

import "github.com/gin-gonic/gin"

type RouterWebsocket struct{}

func (r *RouterWebsocket) InitRouterWebSocket(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		//获取用户列表
		publicGroup.GET("ws/:username", apiPublic.ApiWebSocket.WebSocket)
	}
	return publicGroup
}
