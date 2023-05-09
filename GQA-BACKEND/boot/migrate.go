package boot

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		fmt.Println("Database migration failed!", zap.Any("error:", err))
		global.GqaLogger.Error("Database migration failed!", zap.Any("error:", err))
		os.Exit(0)
	}
	fmt.Println("Database migration succeeded!")
	global.GqaLogger.Info("Database migration succeeded!")
}
