package achievement

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/router/privaterouter"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/router/publicrouter"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var PluginAchievement = new(Achievement)

type Achievement struct{}

func (*Achievement) PluginCode() string { //实现接口方法，插件编码。返回值：请使用 "plugin-"前缀开头。
	return "plugin-achievement"
}

func (*Achievement) PluginSort() uint { //实现接口方法，插件排序
	return data.PluginSort
}

func (*Achievement) PluginName() string { //实现接口方法，插件名称
	return "成就系统"
}

func (*Achievement) PluginVersion() string { //实现接口方法，插件版本
	return "v0.0.1"
}

func (*Achievement) PluginMemo() string { //实现接口方法，插件描述
	return "这是成就系统插件"
}

func (*Achievement) PluginRouterPublic(publicGroup *gin.RouterGroup) { //实现接口方法，公开路由初始化
	publicrouter.InitPublicRouter(publicGroup)
}

func (*Achievement) PluginRouterPrivate(privateGroup *gin.RouterGroup) { //实现接口方法，鉴权路由初始化
	privaterouter.InitPrivateRouter(privateGroup)
}

func (*Achievement) PluginMigrate() []interface{} { //实现接口方法，迁移插件数据表
	var ModelList = []interface{}{

		model.PluginAchievementCategory{},

		model.PluginAchievementObtain{},
	}
	return ModelList
}

func (*Achievement) PluginData() []interface {
	LoadData(c *gin.Context) (err error)
} { //实现接口方法，初始化数据
	var DataList = []interface {
		LoadData(c *gin.Context) (err error)
	}{
		data.PluginAchievementSysApi,
		data.PluginAchievementSysRoleApi,
		data.PluginAchievementSysMenu,
		data.PluginAchievementSysRoleMenu,
		data.PluginAchievementCategory,
	}
	return DataList
}

func (p *Achievement) PluginCron() ([]gqaModel.SysCron, map[uuid.UUID]func()) {
	return nil, nil
}
