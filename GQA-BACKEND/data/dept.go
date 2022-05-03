package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysDept = new(sysDept)

type sysDept struct{}

func (s *sysDept) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysDept{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_dept 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_dept 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDeptData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_dept 表初始数据成功！")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_dept 表初始数据成功！")
		return nil
	})
}

var sysDeptData = []model.SysDept{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10001, CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是Gin-Quasar-Admin部门",
	}}, DeptCode: "gin-quasar-admin", DeptName: "Gin-Quasar-Admin", Leader: "admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是Gin部门",
	}}, DeptCode: "gin", DeptName: "Gin", Leader: "admin", ParentCode: "gin-quasar-admin"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是Quasar部门",
	}}, DeptCode: "quasar", DeptName: "Quasar", Leader: "admin", ParentCode: "gin-quasar-admin"},
}
