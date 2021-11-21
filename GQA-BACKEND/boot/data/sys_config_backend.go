package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysConfigBackend = new(sysConfigBackend)

type sysConfigBackend struct{}

var sysConfigBackendData = []system.SysConfigBackend{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "Casbin模型", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "casbinModel", Default: "./config/casbin_model.conf",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "用户初始密码", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "defaultPassword", Default: "123456",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "验证码字符数", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "captchaKeyLong", Default: "4",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, Remark: "验证码宽度", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "captchaWidth", Default: "240",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, Remark: "验证码高度", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "captchaHeight", Default: "80",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, Remark: "Jwt签发者", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "jwtIssuer", Default: "gin-quasar-admin",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, Remark: "Jwt签发密钥", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "jwtKey", Default: "gin-quasar-admin",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 8, Remark: "Jwt过期时间", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "jwtExpiresAt", Default: "604800",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 9, Remark: "头像最大上传（M）", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "avatarMaxSize", Default: "3",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 10, Remark: "头像允许后缀", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "avatarExt", Default: ".png,.jpg",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 11, Remark: "文件最大上传（M）", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "fileMaxSize", Default: "10",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 12, Remark: "文件允许后缀", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "fileExt", Default: ".png,.jpg,.docx,.xlsx,.txt,.doc,.xls",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 13, Remark: "网站Logo最大上传（M）", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "webLogoMaxSize", Default: "2",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 14, Remark: "网站Logo允许后缀", CreatedAt: time.Now(), CreatedBy: "admin"},
		GqaOption: "webLogoExt", Default: ".ico,.png,.jpg",
	},
}

func (s *sysConfigBackend) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysConfigBackend{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_config_backend 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_config_backend 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysConfigBackendData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_config_backend 表初始数据成功！")
		global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_config_backend 表初始数据成功！")
		return nil
	})
}
