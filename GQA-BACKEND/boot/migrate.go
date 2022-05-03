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
		fmt.Println("迁移 gin-quasar-admin 数据库失败！", zap.Any("error:", err))
		global.GqaLogger.Error("迁移 gin-quasar-admin 数据库失败！", zap.Any("error:", err))
		os.Exit(0)
	}
	fmt.Println("迁移 gin-quasar-admin 数据库成功！")
	global.GqaLogger.Info("迁移 gin-quasar-admin 数据库成功！")
}
