package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
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
			fmt.Println("[Gin-Quasar-Admin] --> sys_config_backend 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_config_backend 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysConfigBackendData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_config_backend 表初始数据成功！")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_config_backend 表初始数据成功！")
		return nil
	})
}

var sysConfigBackendData = []model.SysConfigBackend{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10001, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户初始密码",
	}}, ConfigItem: "defaultPassword", ItemDefault: "123456"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10002, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "验证码字符数",
	}}, ConfigItem: "captchaKeyLong", ItemDefault: "4"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10003, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "验证码图片宽度",
	}}, ConfigItem: "captchaWidth", ItemDefault: "240"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10004, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "验证码图片高度",
	}}, ConfigItem: "captchaHeight", ItemDefault: "80"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10005, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt签发者",
	}}, ConfigItem: "jwtIssuer", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10006, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt签发者密钥",
	}}, ConfigItem: "jwtKey", ItemDefault: "Gin-Quasar-Admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10007, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt过期时间(秒)",
	}}, ConfigItem: "jwtExpiresAt", ItemDefault: "14400"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10008, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Jwt刷新时间(秒)",
	}}, ConfigItem: "jwtRefreshAt", ItemDefault: "86400"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10009, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "头像最大上传(M)",
	}}, ConfigItem: "avatarMaxSize", ItemDefault: "3"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10010, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "头像允许后缀",
	}}, ConfigItem: "avatarExt", ItemDefault: ".png,.jpg"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10011, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "文件最大上传(M)",
	}}, ConfigItem: "fileMaxSize", ItemDefault: "10"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10012, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "文件允许后缀",
	}}, ConfigItem: "fileExt", ItemDefault: ".png,.jpg,.docx,.xlsx,.txt,.doc,.xls,.zip"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10013, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Logo最大上传(M)",
	}}, ConfigItem: "logoMaxSize", ItemDefault: "2"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10014, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Logo允许后缀",
	}}, ConfigItem: "logoExt", ItemDefault: ".png,.jpg"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10015, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Favicon最大上传(M)",
	}}, ConfigItem: "faviconMaxSize", ItemDefault: "1"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10016, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站Favicon允许后缀",
	}}, ConfigItem: "faviconExt", ItemDefault: ".ico"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10017, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站首页大图最大上传(M)",
	}}, ConfigItem: "bannerImageMaxSize", ItemDefault: "4"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10018, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "网站首页大图允许后缀",
	}}, ConfigItem: "bannerImageExt", ItemDefault: ".png,.jpg"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10019, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传头像保存位置",
	}}, ConfigItem: "uploadAvatarSavePath", ItemDefault: "avatar"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10020, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传文件保存位置",
	}}, ConfigItem: "uploadFileSavePath", ItemDefault: "file"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10021, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传logo保存位置",
	}}, ConfigItem: "uploadLogoSavePath", ItemDefault: "logo"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10022, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "上传首页大图保存位置",
	}}, ConfigItem: "uploadBannerImageSavePath", ItemDefault: "banner"},
}
