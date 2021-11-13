package boot

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Migrate(db *gorm.DB) {
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
		global.GqaLog.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GqaLog.Info("register table success")
}
