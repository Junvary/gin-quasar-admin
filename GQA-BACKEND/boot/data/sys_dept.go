package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
	"time"
)

var SysDept = new(sysDept)

type sysDept struct{}

var department = []system.SysDept{
	{
		GqaModel: global.GqaModel{
			Id:       1,
			Status:   "on",
			Sort:     1,
			Desc:     "这是Gin-Quasar-Admin部门",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 0,
		DeptCode: "gin-quasar-admin",
		DeptName: "Gin-Quasar-Admin部门",
		OwnerId: 1,
	},
	{
		GqaModel: global.GqaModel{
			Id:       2,
			Status:   "on",
			Sort:     1,
			Desc:     "这是Gin部门",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 1,
		DeptCode: "gin",
		DeptName: "Gin部门",
		OwnerId: 1,
	},
	{
		GqaModel: global.GqaModel{
			Id:       3,
			Status:   "on",
			Sort:     2,
			Desc:     "这是Quasar部门",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 1,
		DeptCode: "quasar",
		DeptName: "Quasar部门",
		OwnerId: 1,
	},
}

func (a *sysDept) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysDept{}).RowsAffected == 3 {
			fmt.Println("\n[Mysql] --> sys_dept 表的初始数据已存在！")
			global.GqaLog.Error("sys_dept 表的初始数据已存在！")
			return nil
		}
		if err := tx.Create(&department).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("\n[Mysql] --> sys_dept 表初始数据成功！")
		global.GqaLog.Error("sys_dept 表初始数据成功！")
		return nil
	})
}
