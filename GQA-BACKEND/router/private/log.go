package private

import "github.com/gin-gonic/gin"

type RouterLog struct{}

func (r *RouterLog) InitRouterRouterLog(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("log")
	{
		privateGroup.POST("get-log-login-list", apiPrivate.ApiLogLogin.GetLogLoginList)
		privateGroup.POST("delete-log-login-by-id", apiPrivate.ApiLogLogin.DeleteLogLoginById)
		privateGroup.POST("get-log-operation-list", apiPrivate.ApiLogOperation.GetLogOperationList)
		privateGroup.POST("delete-log-operation-by-id", apiPrivate.ApiLogOperation.DeleteLogOperationById)
	}
	return privateGroup
}
