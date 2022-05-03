package private

import "github.com/gin-gonic/gin"

type RouterConfigFrontend struct{}

func (r *RouterConfigFrontend) InitRouterRouterConfigFrontend(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("config-frontend")
	{
		privateGroup.POST("get-config-frontend-list", apiPrivate.ApiConfigFrontend.GetConfigFrontendList)
		privateGroup.POST("edit-config-frontend", apiPrivate.ApiConfigFrontend.EditConfigFrontend)
		privateGroup.POST("add-config-frontend", apiPrivate.ApiConfigFrontend.AddConfigFrontend)
		privateGroup.POST("delete-config-frontend-by-id", apiPrivate.ApiConfigFrontend.DeleteConfigFrontendById)
	}
	return privateGroup
}
