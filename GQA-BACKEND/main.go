package main

import (
	"gin-quasar-admin/boot"
	"gin-quasar-admin/global"
)

func main() {
	global.GqaViper = boot.Viper()
	global.GqaLog = boot.Zap()
	global.GqaDb = boot.Mysql()
	if global.GqaDb != nil {
		boot.Migrate(global.GqaDb)
		db, _ := global.GqaDb.DB()
		defer db.Close()
	}
	boot.Boot()
}