package boot

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin"
	"gin-quasar-admin/model/system"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Migrate(db *gorm.DB) {
	//迁移Gin-Quasar-Admin数据库
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysRole{},
		system.SysUserRole{},
		system.SysMenu{},
		system.SysRoleMenu{},
		system.SysDept{},
		system.SysDict{},
		system.SysApi{},
		gormadapter.CasbinRule{},
		system.SysConfigBackend{},
		system.SysConfigFrontend{},
	)
	if err != nil {
		global.GqaLog.Error("迁移【Gin-Quasar-admin】数据库失败！", zap.Any("err", err))
		os.Exit(0)
	}
	//迁移GQA-Plugin数据库
	err = db.AutoMigrate(gqaplugin.MigratePluginModel())
	if err != nil {
		global.GqaLog.Error("迁移【GQA-Plugin】数据库失败！", zap.Any("err", err))
		os.Exit(0)
	}
	global.GqaLog.Info("迁移数据库成功！")
}
