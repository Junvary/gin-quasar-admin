package system

import (
	"github.com/gin-gonic/gin"
)

type RouterConfigBackend struct {}

func (r *RouterConfigBackend) InitRouterConfigBackend(Router *gin.RouterGroup) (R gin.IRoutes) {
	configGroup := Router.Group("config-backend")
	{
		//获取config列表
		configGroup.POST("config-backend-list", ApiConfigBackend.GetConfigBackendList)
		//编辑config
		configGroup.PUT("config-backend-edit", ApiConfigBackend.EditConfigBackend)
		//新增config
		configGroup.POST("config-backend-add", ApiConfigBackend.AddConfigBackend)
		//删除config
		configGroup.DELETE("config-backend-delete", ApiConfigBackend.DeleteConfigBackend)
	}
	return Router
}

