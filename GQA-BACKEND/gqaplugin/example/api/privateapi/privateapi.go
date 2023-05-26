package privateapi

import (
	elPath "github.com/Junvary/erleng/path"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"time"
)

func GetTestDataList(c *gin.Context) {
	var getTestDataList model.RequestGetTestDataList
	if err := gqaModel.RequestShouldBindJSON(c, &getTestDataList); err != nil {
		return
	}
	if err, exportData, total := privateservice.GetTestDataList(getTestDataList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取导出数据列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取导出数据列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  exportData,
			Page:     getTestDataList.Page,
			PageSize: getTestDataList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditTestData(c *gin.Context) {
	var toEditTestData model.PluginExampleTestData
	if err := gqaModel.RequestShouldBindJSON(c, &toEditTestData); err != nil {
		return
	}
	toEditTestData.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditTestData(toEditTestData, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑导出数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑导出数据失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑导出数据成功！", c)
	}
}

func AddTestData(c *gin.Context) {
	var toAddTestData model.RequestAddTestData
	if err := gqaModel.RequestShouldBindJSON(c, &toAddTestData); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddTestData.Status,
			Sort:      toAddTestData.Sort,
			Memo:      toAddTestData.Memo,
		},
	}
	addTestData := &model.PluginExampleTestData{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Column1:                           toAddTestData.Column1,
		Column2:                           toAddTestData.Column2,
		Column3:                           toAddTestData.Column3,
		Column4:                           toAddTestData.Column4,
		Column5:                           toAddTestData.Column5,
	}
	if err := privateservice.AddTestData(*addTestData, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加导出数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加导出数据失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加导出数据成功！", c)
	}
}

func DeleteTestDataById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteTestDataById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除导出数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除导出数据失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除导出数据成功！", c)
	}
}

func QueryTestDataById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryTestDataById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找导出数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找导出数据失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找导出数据成功！", c)
	}
}

func DownloadTemplateTestData(c *gin.Context) {
	var filename model.RequestFilename
	if err := gqaModel.RequestShouldBindJSON(c, &filename); err != nil {
		return
	}
	err := elPath.CreatePath(gqaGlobal.GqaConfig.System.TemplatePath)
	if err != nil {
		gqaGlobal.GqaLogger.Error("创建模板文件夹失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("创建模板文件夹失败，"+err.Error(), c)
	}
	templateFile := gqaGlobal.GqaConfig.System.TemplatePath + "/" + filename.Filename
	_, err = os.Stat(templateFile)
	if err != nil {
		gqaGlobal.GqaLogger.Error("模板文件不存在！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("模板文件不存在，"+err.Error(), c)
	}
	c.File(templateFile)
}

func ExportTestData(c *gin.Context) {
	var getTestDataList model.RequestGetTestDataList
	if err := gqaModel.RequestShouldBindJSON(c, &getTestDataList); err != nil {
		return
	}
	err := elPath.CreatePath(gqaGlobal.GqaConfig.System.ExportPath)
	if err != nil {
		gqaGlobal.GqaLogger.Error("创建文件夹失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("创建文件夹失败，"+err.Error(), c)
	}
	filePath := gqaGlobal.GqaConfig.System.ExportPath + "/GqaExport-" + time.Now().Format("20060102150405") + ".xlsx"
	if err := privateservice.ExportTestData(getTestDataList, filePath, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("导出数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("导出数据失败！"+err.Error(), c)
	} else {
		c.File(filePath)
	}
}

func ImportTestData(c *gin.Context) {
	username := gqaUtils.GetUsername(c)
	_, avatarHeader, err := c.Request.FormFile("file")
	if err != nil {
		gqaGlobal.GqaLogger.Error("解析文件失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("解析文件失败，"+err.Error(), c)
		return
	}
	err = elPath.CreatePath(gqaGlobal.GqaConfig.System.ImportPath)
	if err != nil {
		gqaGlobal.GqaLogger.Error("创建导入文件夹失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("创建导入文件夹失败，"+err.Error(), c)
	}
	filename := username + "-" + time.Now().Format("20060102150405") + ".xlsx"
	err = c.SaveUploadedFile(avatarHeader, gqaGlobal.GqaConfig.System.ImportPath+"/"+filename)
	if err != nil {
		return
	}
	if err = privateservice.ImportTestData(filename); err != nil {
		gqaGlobal.GqaLogger.Error("导入数据失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("导入数据失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("导入数据成功！", c)
	}
}
