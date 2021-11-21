package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysDept = new(sysDept)

type sysDept struct{}

var sysDeptData = []system.SysDept{
	// 父级设置ID
	{GqaModel: global.GqaModel{Id: 1, Stable: "yes", Status: "on", Sort: 1, Remark: "这是Gin-Quasar-Admin部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentId: 0, DeptCode: "gin-quasar-admin", DeptName: "Gin-Quasar-Admin部门", OwnerId: 1,
	},

	// 子级内容：
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是Gin部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentId: 1, DeptCode: "gin", DeptName: "Gin部门", OwnerId: 1,
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是Quasar部门", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentId: 1, DeptCode: "quasar", DeptName: "Quasar部门", OwnerId: 1,
	},
}

func (s *sysDept) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysDept{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_dept 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_dept 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDeptData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_dept 表初始数据成功！")
		global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_dept 表初始数据成功！")
		return nil
	})
}
