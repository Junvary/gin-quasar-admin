package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var PluginExampleTestData = new(exportData)

type exportData struct{}

func (s *exportData) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.PluginExampleTestData{}).Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> example插件testData表存在数据，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-plugins] --> example插件testData表存在数据，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Create(CreateData()).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> example插件初始数据进入 testData 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-plugins] --> example插件初始数据进入 testData 表成功！")
		return nil
	})
}

func CreateData() (d []model.PluginExampleTestData) {
	for i := 1; i < 500; i++ {
		d = append(d, model.PluginExampleTestData{
			GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{
				GqaModel: gqaGlobal.GqaModel{
					Sort:   uint(100 + i),
					Stable: "",
					Status: "",
					Memo:   "",
				},
			},
			Column1: "第1列内容" + strconv.Itoa(i),
			Column2: "第2列内容" + strconv.Itoa(i),
			Column3: "第3列内容" + strconv.Itoa(i),
			Column4: "第4列内容" + strconv.Itoa(i),
			Column5: "第5列内容" + strconv.Itoa(i),
		})
	}
	return d
}
