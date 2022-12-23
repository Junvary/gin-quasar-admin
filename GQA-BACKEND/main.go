package main

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/boot"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
)

func main() {
	global.GqaViper = boot.Viper()
	global.GqaLogger = boot.Zap()
	global.GqaDb = boot.Mysql()
	if global.GqaDb != nil {
		boot.Migrate(global.GqaDb)
		db, _ := global.GqaDb.DB()
		defer db.Close()
	}
	global.GqaCron = boot.Cron()
	boot.Boot()
}
