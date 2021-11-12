package public

import (
	"github.com/gin-gonic/gin"
)

type RouterLogin struct {}

func (r *RouterLogin) InitRouterLogin(Router *gin.RouterGroup) (R gin.IRoutes) {
	{
		//登录入口
		Router.POST("login", ApiLogin.Login)
	}
	return Router
}