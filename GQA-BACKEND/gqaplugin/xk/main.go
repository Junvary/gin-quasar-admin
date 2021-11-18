package xk

import (
	"gin-quasar-admin/gqaplugin/xk/data"
	"gin-quasar-admin/gqaplugin/xk/model"
	"gin-quasar-admin/gqaplugin/xk/router/private_router"
	"gin-quasar-admin/gqaplugin/xk/router/public_router"
	"github.com/gin-gonic/gin"
)

var PluginXk = new(xk)

type xk struct{}

func (*xk) PluginCode() string {
	return "xk"
}

func (*xk) PluginName() string {
	return "哈哈哈"
}

func (p *xk) PluginRouterPublic(publicGroup *gin.RouterGroup) {
	//实现接口方法，公开路由初始化
	public_router.InitPublicRouter(publicGroup)
}

func (p *xk) PluginRouterPrivate(privateGroup *gin.RouterGroup) {
	//实现接口方法，鉴权路由初始化
	private_router.InitPrivateRouter(privateGroup)
}

func (p *xk) PluginMigrate() []interface{} {
	//实现接口方法，数据表初始化，填入插件中的数据库模型
	var ModelList = []interface{}{
		model.GqaPluginXkNews{},
	}
	return ModelList
}

func (p *xk) PluginData() []interface{ LoadData() (err error) } {
	//实现接口方法，初始化数据，填入插件中的初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginXkNews,
	}
	return DataList
}
