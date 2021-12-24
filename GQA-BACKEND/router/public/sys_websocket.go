package public

import "github.com/gin-gonic/gin"

type RouterWebSocket struct {}

func (r *RouterWebSocket) InitRouterWebSocket(Router *gin.RouterGroup) (R gin.IRoutes) {
	webSocketGroup := Router.Group("")
	{
		//获取用户列表
		webSocketGroup.GET("ws/:username", ApiWebSocket.WebSocket)
	}
	return Router
}
