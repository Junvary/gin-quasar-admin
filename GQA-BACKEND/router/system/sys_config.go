package system

import (
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {}

func (r *RouterConfig) InitRouterConfig(Router *gin.RouterGroup) (R gin.IRoutes) {
	configGroup := Router.Group("config")
	{
		//获取config列表
		configGroup.POST("config-list", ApiConfig.GetConfigList)
		//编辑config
		configGroup.PUT("config-edit", ApiConfig.EditConfig)
		//新增config
		configGroup.POST("config-add", ApiConfig.AddConfig)
		//删除config
		configGroup.DELETE("config-delete", ApiConfig.DeleteConfig)
	}
	return Router
}

