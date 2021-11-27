package private_router

import (
	"gin-quasar-admin/gqaplugin/example/api/private_api"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		//获取news列表
		privateGroup.POST("news-list", private_api.GetNews)
	}
}
