package public

import "github.com/gin-gonic/gin"

type RouterConfigFrontend struct{}

func (r *RouterConfigFrontend) InitRouterConfigFrontend(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		publicGroup.POST("get-config-frontend-all", apiPublic.ApiConfigFrontend.GetConfigFrontendAll)
	}
	return publicGroup
}
