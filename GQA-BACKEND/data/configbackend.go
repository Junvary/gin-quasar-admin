package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysConfigBackend = new(sysConfigBackend)

type sysConfigBackend struct{}

func (s *sysConfigBackend) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysConfigBackend{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData("SkipInsertWithData", "sys_config_backend"), count)
			global.GqaLogger.Warn(utils.GqaI18nWithData("SkipInsertWithData", "sys_config_backend"), zap.Any("count", count))
			return nil
		}
		if err := tx.Create(&sysConfigBackendData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData("TableInitSuccess", "sys_config_backend"))
		global.GqaLogger.Info(utils.GqaI18nWithData("TableInitSuccess", "sys_config_backend"))
		return nil
	})
}

var sysConfigBackendData = []model.SysConfigBackend{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户初始密码",
	}}, ConfigItem: "defaultPassword", ItemDefault: "123456"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "验证码字符数",
	}}, ConfigItem: "captchaKeyLong", ItemDefault: "4"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "验证码图片宽度",
	}}, ConfigItem: "captchaWidth", ItemDefault: "240"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "验证码图片高度",
	}}, ConfigItem: "captchaHeight", ItemDefault: "80"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt签发者",
	}}, ConfigItem: "jwtIssuer", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt签发者密钥",
	}}, ConfigItem: "jwtKey", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 7, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt过期时间(秒)",
	}}, ConfigItem: "jwtExpiresAt", ItemDefault: "14400"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 8, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt刷新时间(秒)",
	}}, ConfigItem: "jwtRefreshAt", ItemDefault: "86400"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 9, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "头像最大上传(M)",
	}}, ConfigItem: "avatarMaxSize", ItemDefault: "3"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "头像允许后缀",
	}}, ConfigItem: "avatarExt", ItemDefault: ".png,.jpg"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 11, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "文件最大上传(M)",
	}}, ConfigItem: "fileMaxSize", ItemDefault: "10"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 12, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "文件允许后缀",
	}}, ConfigItem: "fileExt", ItemDefault: ".png,.jpg,.docx,.xlsx,.txt,.doc,.xls,.zip"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 13, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Logo最大上传(M)",
	}}, ConfigItem: "logoMaxSize", ItemDefault: "2"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 14, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Logo允许后缀",
	}}, ConfigItem: "logoExt", ItemDefault: ".png,.jpg"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 15, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Favicon最大上传(M)",
	}}, ConfigItem: "faviconMaxSize", ItemDefault: "1"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 16, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Favicon允许后缀",
	}}, ConfigItem: "faviconExt", ItemDefault: ".ico"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 17, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站首页大图最大上传(M)",
	}}, ConfigItem: "bannerImageMaxSize", ItemDefault: "4"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 18, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站首页大图允许后缀",
	}}, ConfigItem: "bannerImageExt", ItemDefault: ".png,.jpg"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 19, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传头像保存位置",
	}}, ConfigItem: "uploadAvatarSavePath", ItemDefault: "avatar"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 20, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传文件保存位置",
	}}, ConfigItem: "uploadFileSavePath", ItemDefault: "file"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 21, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传logo保存位置",
	}}, ConfigItem: "uploadLogoSavePath", ItemDefault: "logo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 22, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传首页大图保存位置",
	}}, ConfigItem: "uploadBannerImageSavePath", ItemDefault: "banner"},
}
