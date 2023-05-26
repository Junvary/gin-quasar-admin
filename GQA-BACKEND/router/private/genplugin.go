package private

import (
	"github.com/gin-gonic/gin"
)

type RouterGenPlugin struct{}

func (r *RouterGenPlugin) InitRouterGenPlugin(Router *gin.RouterGroup) (R gin.IRoutes) {
	genPluginGroup := Router.Group("gen-plugin") /*.Use(middleware.LogOperationHandler())*/
	{
		//已生成插件的列表
		genPluginGroup.POST("get-gen-plugin-list", apiPrivate.ApiGenPlugin.GetGenPluginList)
		//生成插件
		genPluginGroup.POST("gen-plugin", apiPrivate.ApiGenPlugin.GenPlugin)
		//删除插件
		genPluginGroup.POST("delete-gen-plugin-by-id", apiPrivate.ApiGenPlugin.DeleteGenPluginById)
		//下载插件
		genPluginGroup.POST("download-gen-plugin-by-id", apiPrivate.ApiGenPlugin.DownloadGenPluginById)
	}
	return Router
}
