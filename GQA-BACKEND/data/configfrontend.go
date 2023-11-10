package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var SysConfigFrontend = new(sysConfigFrontend)

type sysConfigFrontend struct{}

func (s *sysConfigFrontend) LoadData(c *gin.Context) error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysConfigFrontend{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_config_frontend"), count)
			global.GqaSLogger.Warn(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_config_frontend"), "has_count", count)
			return nil
		}
		if err := tx.Create(&sysConfigFrontendData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_config_frontend"))
		global.GqaSLogger.Info(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_config_frontend"))
		return nil
	})
}

var sysConfigFrontendData = []model.SysConfigFrontend{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "首页大图",
	}}, ConfigItem: "bannerImage", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站主标题",
	}}, ConfigItem: "mainTitle", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站副标题",
	}}, ConfigItem: "subTitle", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站描述",
	}}, ConfigItem: "webDescribe", ItemDefault: "Lorem ipsum dolor sit amet consectetur adipisicing elit"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "登录页风格",
	}}, ConfigItem: "loginLayoutStyle", ItemDefault: "loginLayoutStyle_login"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "登录页插件",
	}}, ConfigItem: "pluginLoginLayout", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 7, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Logo",
	}}, ConfigItem: "logo", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 8, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站favicon",
	}}, ConfigItem: "favicon", ItemDefault: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 9, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "显示仓库入口",
	}}, ConfigItem: "showGit", ItemDefault: "yesNo_yes"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 10, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "星星颜色",
	}}, ConfigItem: "starColor", ItemDefault: "#00b4ff"},
}
