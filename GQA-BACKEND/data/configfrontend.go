package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysConfigFrontend = new(sysConfigFrontend)

type sysConfigFrontend struct{}

func (s *sysConfigFrontend) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysConfigFrontend{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_config_frontend 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_config_frontend 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysConfigFrontendData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_config_frontend 表初始数据成功！")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_config_frontend 表初始数据成功！")
		return nil
	})
}

var sysConfigFrontendData = []model.SysConfigFrontend{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10001, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "首页大图",
	}}, ConfigItem: "bannerImage", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10002, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站主标题",
	}}, ConfigItem: "mainTitle", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10003, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站副标题",
	}}, ConfigItem: "subTitle", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10004, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站描述",
	}}, ConfigItem: "webDescribe", ItemDefault: "Lorem ipsum dolor sit amet consectetur adipisicing elit"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10005, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "登录页插件",
	}}, ConfigItem: "pluginLoginLayout", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10006, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Logo",
	}}, ConfigItem: "logo", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10007, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站favicon",
	}}, ConfigItem: "favicon", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10008, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "显示仓库入口",
	}}, ConfigItem: "showGit", ItemDefault: "yes"},
}
