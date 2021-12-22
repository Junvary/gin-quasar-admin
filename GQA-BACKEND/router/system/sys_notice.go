package system

import "github.com/gin-gonic/gin"

type RouterNotice struct {}

func (r *RouterNotice) InitRouterNotice(Router *gin.RouterGroup) (R gin.IRoutes) {
	noticeGroup := Router.Group("notice")/*.Use(middleware.LogOperationHandler())*/
	{
		//获取消息列表
		noticeGroup.POST("notice-list", ApiNotice.GetNoticeList)
		//编辑消息信息
		noticeGroup.PUT("notice-edit", ApiNotice.EditNotice)
		//新增消息
		noticeGroup.POST("notice-add", ApiNotice.AddNotice)
		//删除消息
		noticeGroup.DELETE("notice-delete", ApiNotice.DeleteNotice)
		//根据ID查找消息
		noticeGroup.POST("notice-id", ApiNotice.QueryNoticeById)
		//根据ID查找消息并阅读
		noticeGroup.POST("notice-id-read", ApiNotice.QueryNoticeByIdRead)
		//发送消息
		noticeGroup.POST("notice-send", ApiNotice.SendNotice)
	}
	return Router
}
