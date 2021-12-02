package gqaplugin

import (
	//其他引用
	"github.com/gin-gonic/gin"
	// 1.插件引入方式1：github插件引入方式
	example "github.com/Junvary/gqa-plugin-example"
	xk "github.com/Junvary/gqa-plugin-xk"
	// 2.插件引入方式2：本地插件引入方式
	//"gin-quasar-admin/gqaplugin/example"
	//"gin-quasar-admin/gqaplugin/xk"
)

/*
	1.import插件(github模式、本地模式)
	2.最好为github模式引入方式提供别名
	3.插件填入下面的插件列表：pluginList
	4.本地插件开发完毕，可按xk方式提交单独仓库引用（上面方式1），即可去除本地插件包目录，如xk目录。
*/

var PluginList = []GqaPlugin{ //插件列表顺序填入
	example.PluginExample,
	xk.PluginXk,
}

/*
	下方为各种安装插件的函数
*/

type GqaPlugin interface { //插件需实现的接口
	PluginCode() string                                //插件编码，用于路由分组名
	PluginName() string                                //插件名称
	PluginVersion() string                             //插件版本
	PluginRemark() string                              //插件描述
	PluginRouterPublic(publicGroup *gin.RouterGroup)   //公开路由
	PluginRouterPrivate(privateGroup *gin.RouterGroup) //鉴权路由
	PluginMigrate() []interface{}                      //迁移数据库模型
	PluginData() []interface{ LoadData() (err error) } //初始化数据
}

func RegisterPluginRouter(PublicGroup, PrivateGroup *gin.RouterGroup) { //注册插件路由
	for _, p := range PluginList {
		PluginRouter(PublicGroup, PrivateGroup, p)
	}
}

func MigratePluginModel() []interface{} { //迁移插件数据库
	var model []interface{}
	for _, p := range PluginList {
		model = append(model,
			p.PluginMigrate()...,
		)
	}
	return model
}

func LoadPluginData() []interface{ LoadData() (err error) } { //初始化插件数据
	var data []interface{ LoadData() (err error) }
	for _, p := range PluginList {
		data = append(data,
			p.PluginData()...,
		)
	}
	return data
}

func PluginRouter(publicGroup, privateGroup *gin.RouterGroup, Plugin ...GqaPlugin) {
	//为插件单独提供路由分组，分组取名为：上面接口中的 PluginCode() 方法
	for i := range Plugin {
		PublicGroup := publicGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPublic(PublicGroup)
		PrivateGroup := privateGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPrivate(PrivateGroup)
	}
}
