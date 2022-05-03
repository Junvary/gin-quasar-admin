package public

import "github.com/gin-gonic/gin"

type RouterConfigBackend struct{}

func (r *RouterConfigBackend) InitRouterConfigBackend(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		publicGroup.POST("get-config-backend-all", apiPublic.ApiConfigBackend.GetConfigBackendAll)
	}
	return publicGroup
}
