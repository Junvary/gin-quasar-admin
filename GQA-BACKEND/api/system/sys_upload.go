package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/service"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiUpload struct {
}

func (a *ApiUpload) UploadAvatar(c *gin.Context) {
	username := utils.GetUsername(c)
	_, fileHeader, err := c.Request.FormFile("file")
	if err!= nil{
		global.GqaLog.Error("解析头像失败！", zap.Any("err", err))
		global.ErrorMessage("解析头像失败，"+err.Error(), c)
		return
	}
	err, avatarUrl := service.GroupServiceApp.ServiceSystem.UploadAvatar(username, fileHeader)
	if err != nil {
		global.GqaLog.Error("上传头像失败！", zap.Any("err", err))
		global.ErrorMessage("上传头像失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": avatarUrl}, c)
	}
}

func (a *ApiUpload) UploadFile(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	if err!= nil{
		global.GqaLog.Error("解析文件失败！", zap.Any("err", err))
		global.ErrorMessage("解析文件失败，"+err.Error(), c)
		return
	}
	err, fileUrl := service.GroupServiceApp.ServiceSystem.UploadFile(fileHeader)
	if err != nil {
		global.GqaLog.Error("上传文件失败！", zap.Any("err", err))
		global.ErrorMessage("上传文件失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": fileUrl}, c)
	}
}
