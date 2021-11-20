package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/example/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginExampleNews = new(pluginExampleNews)

type pluginExampleNews struct{}

var newsData = []model.GqaPluginExampleNews{
	{GqaModel: global.GqaModel{Stable: "no", Status: "on", Sort: 1, Remark: "新闻1", CreatedAt: time.Now(), CreatedBy: "admin"},
		Title: "这是测试新闻1的标题", Content: "这是测试新闻1的内容",
	},
}

func (p *pluginExampleNews) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.GqaPluginExampleNews{}).Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> gqa_plugin_example_news 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("gqa_plugin_example_news 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&newsData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> gqa_plugin_example_news 表初始数据成功！")
		global.GqaLog.Error("gqa_plugin_example_news 表初始数据成功！")
		return nil
	})
}
