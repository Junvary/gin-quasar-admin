package gqaplugin

/*
    // =============================== 引入方式 ===============================
	// 1. 本地引入方式：（可以git忽略掉，只共本地开发使用，不在主仓库中追踪）
	// 如："github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xtkfk"
    // 2. github引入方式：（本地开发完善后，可单独提取成仓库，推荐）
	// 如："github.com/Junvary/gqa-plugin-xtkfk/gqaplugin/xtkfk"
    // ========================== 开发插件、接入插件 ============================
    // 开发插件需实现 GqaPlugin 接口，完成后填入 PluginList 切片中即可
	// ========================================================================
*/

import (
	"github.com/gin-gonic/gin"
)

var PluginList = []GqaPlugin{ //插件加入此切片即可
	//xtkfk.PluginXtkfk,
}

/*
	===============================下方为各种安装插件的函数===============================
*/

type GqaPlugin interface { //插件需实现的接口
	PluginCode() string                                //插件编码，用于路由分组名
	PluginName() string                                //插件名称
	PluginVersion() string                             //插件版本
	PluginMemo() string                                //插件描述
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
