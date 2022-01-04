package xk

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/router/private_router"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/router/public_router"
	"github.com/gin-gonic/gin"
)

var PluginXk = new(xk)

type xk struct{}

func (*xk) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-xk"
}

func (*xk) PluginName() string { //实现接口方法，插件名称
	return "系统开发科"
}

func (*xk) PluginVersion() string { //实现接口方法，插件版本
	return "v0.0.1"
}

func (*xk) PluginRemark() string { //实现接口方法，插件描述
	return "这是系统开发科插件。"
}

func (p *xk) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
	public_router.InitPublicRouter(publicGroup)
}

func (p *xk) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
	private_router.InitPrivateRouter(privateGroup)
}

func (p *xk) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	var ModelList = []interface{}{
		model.GqaPluginXkNews{},
		model.GqaPluginXkProject{},
		model.GqaPluginXkProjectDetail{},
		model.GqaPluginXkHonour{},
		model.GqaPluginXkDocument{},
		model.GqaPluginXkDownload{},
	}
	return ModelList
}

func (p *xk) PluginData() []interface{ LoadData() (err error) } { //实现接口方法，初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginXkSysApi,
		data.PluginXkSysCasbin,
		data.PluginXkSysMenu,
		data.PluginXkSysRoleMenu,
		data.PluginXkSysDict,
	}
	return DataList
}
