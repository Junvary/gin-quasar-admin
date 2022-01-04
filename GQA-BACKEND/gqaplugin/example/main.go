package example

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/router/private_router"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/router/public_router"
	"github.com/gin-gonic/gin"
)

var PluginExample = new(example)

type example struct{}

func (*example) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-example"
}

func (*example) PluginName() string { //实现接口方法，插件名称
	return "Plugin Example"
}

func (*example) PluginVersion() string { //实现接口方法，插件版本
	return "v0.0.1"
}

func (*example) PluginRemark() string { //实现接口方法，插件描述
	return "describe plugin example"
}

func (p *example) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
	public_router.InitPublicRouter(publicGroup)
}

func (p *example) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
	private_router.InitPrivateRouter(privateGroup)
}

func (p *example) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	var ModelList = []interface{}{
		model.GqaPluginExampleNews{},
	}
	return ModelList
}

func (p *example) PluginData() []interface{ LoadData() (err error) } { //实现接口方法，初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginExampleNews,
		data.PluginExampleSysApi,
		data.PluginExampleSysCasbin,
		data.PluginExampleSysMenu,
		data.PluginExampleSysRoleMenu,
	}
	return DataList
}
