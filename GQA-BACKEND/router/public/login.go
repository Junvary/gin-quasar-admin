package public

import "github.com/gin-gonic/gin"

type RouterLogin struct{}

func (r *RouterLogin) InitRouterLogin(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		publicGroup.POST("login", apiPublic.ApiLogin.Login)
	}
	return publicGroup
}
