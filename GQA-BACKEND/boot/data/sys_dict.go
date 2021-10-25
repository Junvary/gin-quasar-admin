package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
	"time"
)

var SysDict = new(sysDict)

type sysDict struct{}

var dictDetail = []system.SysDict{
	{
		GqaModel: global.GqaModel{
			Id:       1,
			Status:   "on",
			Sort:     1,
			Desc:     "这是一个父字典",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 0,
		Value: "gender",
		Label: "性别",
	},
	{
		GqaModel: global.GqaModel{
			Id:       2,
			Status:   "on",
			Sort:     1,
			Desc:     "这是男",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 1,
		Value: "m",
		Label: "男",
	},
	{
		GqaModel: global.GqaModel{
			Id:       3,
			Status:   "on",
			Sort:     2,
			Desc:     "这是女",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 1,
		Value: "f",
		Label: "女",
	},
	{
		GqaModel: global.GqaModel{
			Id:       4,
			Status:   "on",
			Sort:     3,
			Desc:     "这是保密",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 1,
		Value: "u",
		Label: "保密",
	},
	{
		GqaModel: global.GqaModel{
			Id:       5,
			Status:   "on",
			Sort:     2,
			Desc:     "这是是否状态",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 0,
		Value: "status",
		Label: "是否状态",
	},
	{
		GqaModel: global.GqaModel{
			Id:       6,
			Status:   "on",
			Sort:     1,
			Desc:     "这是是",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 5,
		Value: "on",
		Label: "是",
	},
	{
		GqaModel: global.GqaModel{
			Id:       7,
			Status:   "on",
			Sort:     2,
			Desc:     "这是否",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId: 5,
		Value: "off",
		Label: "否",
	},
}

func (a *sysDict) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysDict{}).RowsAffected == 4 {
			fmt.Println("\n[Mysql] --> sys_dict 表的初始数据已存在！")
			global.GqaLog.Error("sys_dict 表的初始数据已存在！")
			return nil
		}
		if err := tx.Create(&dictDetail).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("\n[Mysql] --> sys_dict 表初始数据成功！")
		global.GqaLog.Error("sys_dict 表初始数据成功！")
		return nil
	})
}
