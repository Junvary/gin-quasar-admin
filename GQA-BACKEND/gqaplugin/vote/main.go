package vote

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/router/private_router"
	"github.com/gin-gonic/gin"
)

var PluginVote = new(vote)

type vote struct{}

func (*vote) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-vote"
}

func (*vote) PluginName() string { //实现接口方法，插件名称
	return "投票系统"
}

func (*vote) PluginVersion() string { //实现接口方法，插件版本
	return "v0.0.1"
}

func (*vote) PluginRemark() string { //实现接口方法，插件描述
	return "这是投票系统"
}

func (p *vote) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
	//public_router.InitPublicRouter(publicGroup)
}

func (p *vote) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
	private_router.InitPrivateRouter(privateGroup)
}

func (p *vote) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	var ModelList = []interface{}{
		model.GqaPluginVoteVoter{},
		model.GqaPluginVoteCandidate{},
		model.GqaPluginVoteScoreResult{},
		model.GqaPluginVoteVoterRandom{},
	}
	return ModelList
}

func (p *vote) PluginData() []interface{ LoadData() (err error) } { //实现接口方法，初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginVoteSysApi,
		data.PluginVoteSysCasbin,
		data.PluginVoteSysMenu,
		data.PluginVoteSysRoleMenu,
		data.PluginVoteSysDict,
	}
	return DataList
}
