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

var SysUser = new(sysUser)

type sysUser struct{}

func (s *sysUser) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysUser{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData("SkipInsertWithData", "sys_user"), count)
			global.GqaLogger.Warn(utils.GqaI18nWithData("SkipInsertWithData", "sys_user"), zap.Any("count", count))
			return nil
		}
		if err := tx.Create(&sysUserData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData("TableInitSuccess", "sys_user"))
		global.GqaLogger.Info(utils.GqaI18nWithData("TableInitSuccess", "sys_user"))
		return nil
	})
}

var sysUserData = []model.SysUser{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Be the change you want to see in the world.",
	}}, Username: "admin", Nickname: "Mr.J", RealName: "SuperAdmin", Password: "e10adc3949ba59abbe56e057f20f883e",
		Avatar: "", Gender: "gender_unknown", Mobile: "1234567890", Email: "11111111111"},
}
