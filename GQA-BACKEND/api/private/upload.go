package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiUpload struct{}

func (a *ApiUpload) UploadAvatar(c *gin.Context) {
	username := utils.GetUsername(c)
	avatar, avatarHeader, err := c.Request.FormFile("file")
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "ParseFileFailed")+err.Error(), c)
		return
	}
	err, avatarUrl := servicePrivate.ServiceUpload.UploadAvatar(username, avatar, avatarHeader)
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "UploadFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": avatarUrl}, c)
	}
}

func (a *ApiUpload) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "ParseFileFailed")+err.Error(), c)
		return
	}
	err, fileUrl := servicePrivate.ServiceUpload.UploadFile(file, fileHeader)
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "UploadFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": fileUrl}, c)
	}
}

func (a *ApiUpload) UploadBannerImage(c *gin.Context) {
	img, bannerImage, err := c.Request.FormFile("file")
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "ParseFileFailed")+err.Error(), c)
		return
	}
	err, bannerImageUrl := servicePrivate.ServiceUpload.UploadBannerImage(img, bannerImage)
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "UploadFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": bannerImageUrl}, c)
	}
}

func (a *ApiUpload) UploadLogo(c *gin.Context) {
	logo, logoHeader, err := c.Request.FormFile("file")
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "ParseFileFailed")+err.Error(), c)
		return
	}
	err, icoUrl := servicePrivate.ServiceUpload.UploadLogo(logo, logoHeader)
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "UploadFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": icoUrl}, c)
	}
}

func (a *ApiUpload) UploadFavicon(c *gin.Context) {
	logo, logoHeader, err := c.Request.FormFile("file")
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "ParseFileFailed")+err.Error(), c)
		return
	}
	err, icoUrl := servicePrivate.ServiceUpload.UploadFavicon(logo, logoHeader)
	if err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "UploadFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": icoUrl}, c)
	}
}
