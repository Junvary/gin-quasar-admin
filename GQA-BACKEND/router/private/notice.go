package private

import "github.com/gin-gonic/gin"

type RouterNotice struct{}

func (r *RouterNotice) InitRouterRouterNotice(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("notice")
	{
		privateGroup.POST("get-notice-list", apiPrivate.ApiNotice.GetNoticeList)
		privateGroup.POST("add-notice", apiPrivate.ApiNotice.AddNotice)
		privateGroup.POST("delete-notice-by-id", apiPrivate.ApiNotice.DeleteNoticeById)
		privateGroup.POST("query-notice-by-id", apiPrivate.ApiNotice.QueryNoticeById)
		privateGroup.POST("query-notice-read-by-id", apiPrivate.ApiNotice.QueryNoticeReadById)
		privateGroup.POST("send-notice", apiPrivate.ApiNotice.SendNotice)
	}
	return privateGroup
}
