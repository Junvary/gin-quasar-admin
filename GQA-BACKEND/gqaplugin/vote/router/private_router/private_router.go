package private_router

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/api/private_api"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup)  {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	voteRouter := privateGroup.Use(middleware.LogOperationHandler())
	{
		//随机和固定投票人列表
		voteRouter.POST("voter-random-list", private_api.VoterRandomList)
		//获取随机投票人
		voteRouter.POST("voter-random-get", private_api.VoterRandomGet)
		//新增随机和固定投票人
		voteRouter.POST("voter-random-add", private_api.VoterRandomAdd)
		//删除随机和固定投票人
		voteRouter.DELETE("voter-random-delete", private_api.VoterRandomDelete)
		//检查是否可以投票
		voteRouter.POST("can-vote", private_api.CanVote)
		//投票结果列表
		voteRouter.POST("vote-result-list", private_api.VoteResultList)
		//投票结果图形
		voteRouter.POST("vote-result-chart", private_api.VoteResultChart)
		//投票
		voteRouter.POST("vote-handle", private_api.VoteHandle)
		//固定投票人列表
		voteRouter.POST("voter-list", private_api.VoterList)
		//新增固定投票人
		voteRouter.POST("voter-add", private_api.VoterAdd)
		//删除固定投票人
		voteRouter.DELETE("voter-delete", private_api.VoterDelete)
		//候选人列表
		voteRouter.POST("candidate-list", private_api.CandidateList)
		//新增候选人
		voteRouter.POST("candidate-add", private_api.CandidateAdd)
		//删除候选人
		voteRouter.DELETE("candidate-delete", private_api.CandidateDelete)
	}
}
