package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
}

func (r *RouterConfig) InitRouterConfig(Router *gin.RouterGroup) (R gin.IRoutes) {
	configGroup := Router.Group("config")
	apiConfig := api.GroupApiApp.ApiSystem.ApiConfig
	{
		// 获取config列表
		configGroup.POST("config-list", apiConfig.GetConfigList)
		// 编辑config
		configGroup.PUT("config-edit", apiConfig.EditConfig)
		// 新增config
		configGroup.POST("config-add", apiConfig.AddConfig)
		// 删除config
		configGroup.DELETE("config-delete", apiConfig.DeleteConfig)
	}
	return Router
}

