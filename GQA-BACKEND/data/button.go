package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysButton = new(sysButton)

type sysButton struct{}

func (s *sysButton) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysButton{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_button 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_button 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysButtonData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_button 表初始数据成功！")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_button 表初始数据成功！")
		return nil
	})
}

var sysButtonData = []model.SysButton{
	{MenuName: "api", ButtonName: "添加", ButtonCode: "api:add"},
}
