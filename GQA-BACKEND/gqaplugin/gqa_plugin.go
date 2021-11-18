package gqaplugin

import (
	"gin-quasar-admin/gqaplugin/xk"
	"github.com/gin-gonic/gin"
)

type GqaPlugin interface {
	PluginCode() string                                //插件代号，用于路由分组名
	PluginName() string                                //插件名称
	PluginRouterPublic(publicGroup *gin.RouterGroup)   //公开路由
	PluginRouterPrivate(privateGroup *gin.RouterGroup) //鉴权路由
	PluginMigrate() []interface{}                      //迁移数据库模型
	PluginData() []interface{ LoadData() (err error) }       //初始化数据
}

func PluginRouter(publicGroup, privateGroup *gin.RouterGroup, Plugin ...GqaPlugin) {
	for i := range Plugin {
		PublicGroup := publicGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPublic(PublicGroup)
		PrivateGroup := privateGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPrivate(PrivateGroup)
	}
}

func RegisterPluginRouter(PublicGroup, PrivateGroup *gin.RouterGroup) {
	PluginRouter(PublicGroup, PrivateGroup, xk.PluginXk)
	//顺序扩展其他插件
	//PluginRouter(PublicGroup, PrivateGroup, ??.Plugin??)
}

func MigratePluginModel() []interface{} {
	var model []interface{}
	model = append(model,
		xk.PluginXk.PluginMigrate()...,
		//顺序添加其他插件
		//??.Plugin??.PluginMigrate(),
	)
	return model
}

func LoadPluginData() []interface{LoadData() (err error)} {
	var data []interface{LoadData() (err error)}
	data = append(data,
		xk.PluginXk.PluginData()...,
		//顺序添加其他插件
		//??.Plugin??.PluginData()...,
	)
	return data
}
