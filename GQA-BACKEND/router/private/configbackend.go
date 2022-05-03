package private

import "github.com/gin-gonic/gin"

type RouterConfigBackend struct{}

func (r *RouterConfigBackend) InitRouterRouterConfigBackend(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("config-backend")
	{
		privateGroup.POST("get-config-backend-list", apiPrivate.ApiConfigBackend.GetConfigBackendList)
		privateGroup.POST("edit-config-backend", apiPrivate.ApiConfigBackend.EditConfigBackend)
		privateGroup.POST("add-config-backend", apiPrivate.ApiConfigBackend.AddConfigBackend)
		privateGroup.POST("delete-config-backend-by-id", apiPrivate.ApiConfigBackend.DeleteConfigBackendById)
	}
	return privateGroup
}
