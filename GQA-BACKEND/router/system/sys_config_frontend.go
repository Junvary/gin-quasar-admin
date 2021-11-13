package system

import (
	"github.com/gin-gonic/gin"
)

type RouterConfigFrontend struct {}

func (r *RouterConfigFrontend) InitRouterConfigFrontend(Router *gin.RouterGroup) (R gin.IRoutes) {
	configGroup := Router.Group("config-frontend")
	{
		//获取config列表
		configGroup.POST("config-frontend-list", ApiConfigFrontend.GetConfigFrontendList)
		//编辑config
		configGroup.PUT("config-frontend-edit", ApiConfigFrontend.EditConfigFrontend)
		//新增config
		configGroup.POST("config-frontend-add", ApiConfigFrontend.AddConfigFrontend)
		//删除config
		configGroup.DELETE("config-frontend-delete", ApiConfigFrontend.DeleteConfigFrontend)
	}
	return Router
}

