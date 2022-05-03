package private

import (
	"github.com/gin-gonic/gin"
)

type RouterGenPlugin struct{}

func (r *RouterGenPlugin) InitRouterGenPlugin(Router *gin.RouterGroup) (R gin.IRoutes) {
	genPluginGroup := Router.Group("gen-plugin") /*.Use(middleware.LogOperationHandler())*/
	{
		//生成插件
		genPluginGroup.POST("gen-plugin", apiPrivate.ApiGenPlugin.GenPlugin)
	}
	return Router
}
