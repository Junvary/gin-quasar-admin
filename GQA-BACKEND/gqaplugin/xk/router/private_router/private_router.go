package private_router

import "github.com/gin-gonic/gin"

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	{
		privateGroup.GET("", func(context *gin.Context) {
			context.JSON(200, "看不到")
		})
	}
}
