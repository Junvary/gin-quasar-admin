package public

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterLogin struct {
}

func (r *RouterLogin) InitRouterLogin(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiLogin := api.GroupApiApp.ApiPublic.ApiLogin
	{
		Router.POST("login", apiLogin.Login)
	}
	return Router
}