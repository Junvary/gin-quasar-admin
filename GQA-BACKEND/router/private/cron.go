package private

import "github.com/gin-gonic/gin"

type RouterCron struct{}

func (r *RouterCron) InitRouterRouterCron(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("cron")
	{
		privateGroup.POST("get-cron-list", apiPrivate.ApiCron.GetCronList)
		privateGroup.POST("start-cron", apiPrivate.ApiCron.StartCron)
		privateGroup.POST("stop-cron", apiPrivate.ApiCron.StopCron)
	}
	return privateGroup
}
