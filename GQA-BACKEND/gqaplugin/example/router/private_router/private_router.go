package private_router

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/api/public_api"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		//获取news列表
		privateGroup.GET("news-list", public_api.GetNews)
	}
}
