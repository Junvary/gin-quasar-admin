package system

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiUpload struct {
}

func (a *ApiUpload) UploadAvatar(c *gin.Context) {
	username := utils.GetUsername(c)
	avatar, avatarHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GqaLog.Error("解析头像失败！", zap.Any("err", err))
		global.ErrorMessage("解析头像失败，"+err.Error(), c)
		return
	}
	err, avatarUrl := ServiceUpload.UploadAvatar(username, avatar, avatarHeader)
	if err != nil {
		global.GqaLog.Error("上传头像失败！", zap.Any("err", err))
		global.ErrorMessage("上传头像失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": avatarUrl}, c)
	}
}

func (a *ApiUpload) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GqaLog.Error("解析文件失败！", zap.Any("err", err))
		global.ErrorMessage("解析文件失败，"+err.Error(), c)
		return
	}
	err, fileUrl := ServiceUpload.UploadFile(file, fileHeader)
	if err != nil {
		global.GqaLog.Error("上传文件失败！", zap.Any("err", err))
		global.ErrorMessage("上传文件失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": fileUrl}, c)
	}
}

func (a *ApiUpload) UploadWebLogo(c *gin.Context) {
	logo, logoHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GqaLog.Error("解析文件失败！", zap.Any("err", err))
		global.ErrorMessage("解析文件失败，"+err.Error(), c)
		return
	}
	err, icoUrl := ServiceUpload.UploadWebLogo(logo, logoHeader)
	if err != nil {
		global.GqaLog.Error("上传网站Logo失败！", zap.Any("err", err))
		global.ErrorMessage("上传网站Logo失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": icoUrl}, c)
	}
}
