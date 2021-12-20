package system

import "github.com/gin-gonic/gin"

type RouterLog struct{}

func (r *RouterLog) InitRouterLog(Router *gin.RouterGroup) (R gin.IRoutes) {
	logGroup := Router.Group("log")
	{
		//获取登录日志列表
		logGroup.POST("log-login-list", ApiLogLogin.GetLogLoginList)
		//删除登录日志
		logGroup.DELETE("log-login-delete", ApiLogLogin.DeleteLogLogin)
		//获取操作日志列表
		logGroup.POST("log-operation-list", ApiLogOperation.GetLogOperationList)
		//删除操作日志
		logGroup.DELETE("log-operation-delete", ApiLogOperation.DeleteLogOperation)
	}
	return Router
}
