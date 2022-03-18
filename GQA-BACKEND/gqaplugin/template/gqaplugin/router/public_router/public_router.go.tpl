package public_router

import "github.com/gin-gonic/gin"

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
}

