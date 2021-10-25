package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var SysUser = new(sysUser)

type sysUser struct{}

var users = []system.SysUser{
	{
		GqaModel: global.GqaModel{
			Id:       1,
			Status:   "on",
			Sort:     1,
			Desc:     "我是超管我怕谁？",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		Uuid:     uuid.New(),
		Avatar:   "",
		Username: "admin",
		Password: "e10adc3949ba59abbe56e057f20f883e",
		Nickname: "我是超管",
		RealName: "超级管理员",
		Gender:   "u",
		Mobile:   "1234567890",
		Email:    "11111111111",
	},
}

func (a *sysUser) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysUser{}).RowsAffected == 1 {
			fmt.Println("\n[Mysql] --> sys_user 表的初始数据已存在！")
			global.GqaLog.Error("sys_user 表的初始数据已存在！")
			return nil
		}
		if err := tx.Create(&users).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("\n[Mysql] --> sys_user 表初始数据成功！")
		global.GqaLog.Error("sys_user 表初始数据成功！")
		return nil
	})
}
