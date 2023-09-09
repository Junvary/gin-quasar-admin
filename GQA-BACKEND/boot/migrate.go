package boot

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"gorm.io/gorm"
	"os"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		fmt.Println("Database migration failed!", err.Error())
		global.GqaSLogger.Error("Database migration failed!", "err", err.Error())
		os.Exit(0)
	}
	fmt.Println("Database migration succeeded!")
	global.GqaSLogger.Info("Database migration succeeded!")
}
