package public_router

import (
	"github.com/Junvary/gqa-plugin-xk/api/public_api"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	{
		//获取news列表
		publicGroup.POST("news-list", public_api.GetNews)
		//获取project列表
		publicGroup.POST("project-list", public_api.GetProject)
		//获取武器库语言饼状图
		publicGroup.POST("weapon-language-list", public_api.GetWeaponLanguage)
		//获取武器库节点饼状图
		publicGroup.POST("weapon-node-list", public_api.GetWeaponNode)
		//获取honour列表
		publicGroup.POST("honour-list", public_api.GetHonour)
		//获取document列表
		publicGroup.POST("document-list", public_api.GetDocument)
		//获取download列表
		publicGroup.POST("download-list", public_api.GetDownload)
	}
}
