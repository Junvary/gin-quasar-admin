package system

import (
	"github.com/gin-gonic/gin"
)

type RouterGenPlugin struct{}

func (r *RouterGenPlugin) InitRouterGenPlugin(Router *gin.RouterGroup) (R gin.IRoutes) {
	genPluginGroup := Router.Group("gen-plugin") /*.Use(middleware.LogOperationHandler())*/
	{
		//获取api列表
		genPluginGroup.POST("gen-plugin", ApiGenPlugin.GenPlugin)
	}
	return Router
}
