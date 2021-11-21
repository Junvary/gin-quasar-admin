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

func (*example) PluginCode() string { //PluginCode返回值：请使用 "plugin-"前缀开头
	return "plugin-example"
}

func (*example) PluginName() string {
	return "插件示例"
}

func (p *example) PluginRouterPublic(publicGroup *gin.RouterGroup) {
	//实现接口方法，公开路由初始化
	public_router.InitPublicRouter(publicGroup)
}

func (p *example) PluginRouterPrivate(privateGroup *gin.RouterGroup) {
	//实现接口方法，鉴权路由初始化
	private_router.InitPrivateRouter(privateGroup)
}

func (p *example) PluginMigrate() []interface{} {
	//实现接口方法，数据表初始化，填入插件中的数据库模型
	var ModelList = []interface{}{
		model.GqaPluginExampleNews{},
	}
	return ModelList
}

func (p *example) PluginData() []interface{ LoadData() (err error) } {
	//实现接口方法，初始化数据，填入插件中的初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginExampleNews,
		data.PluginExampleSysApi,
		data.PluginExampleSysCasbin,
	}
	return DataList
}
