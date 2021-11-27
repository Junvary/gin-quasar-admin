package public_router

import (
	"gin-quasar-admin/gqaplugin/example/api/public_api"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		publicGroup.GET("news-list", public_api.GetNews)
	}
}
