package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysConfig = new(sysConfig)

type sysConfig struct{}

var sysConfigData = []system.SysConfig{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "用户初始密码", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "defaultPassword", Default: "123456",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "验证码字符数", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "captchaKeyLong", Default: "6",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "验证码宽度", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "captchaWidth", Default: "240",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, Remark: "验证码高度", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "captchaHeight", Default: "80",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, Remark: "Jwt签发者", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "jwtIssuer", Default: "gin-quasar-admin",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, Remark: "Jwt签发密钥", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "jwtKey", Default: "gin-quasar-admin",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, Remark: "Jwt过期时间", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "jwtExpiresAt", Default: "604800",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 8, Remark: "头像最大上传（M）", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "avatarMaxSize", Default: "3",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 9, Remark: "头像允许后缀", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "avatarExt", Default: ".png,.jpg",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 10, Remark: "文件最大上传（M）", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "fileMaxSize", Default: "10",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 11, Remark: "文件允许后缀", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		GqaOption: "fileExt", Default: ".png,.jpg,.docx,.xlsx,.txt,.doc,.xls",
	},


}

func (s *sysConfig) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysConfig{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_config 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("sys_config 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysConfigData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_config 表初始数据成功！")
		global.GqaLog.Error("sys_config 表初始数据成功！")
		return nil
	})
}
