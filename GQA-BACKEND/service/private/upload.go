package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"mime/multipart"
	"strconv"
	"time"
)

type ServiceUpload struct{}

func (s *ServiceUpload) UploadAvatar(username string, avatar multipart.File, avatarHeader *multipart.FileHeader) (err error, avatarUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("avatarMaxSize")
	if maxSizeString == "" {
		return errors.New("找不到头像大小配置！"), ""
	}
	if !utils.CheckFileSize(avatar, maxSizeString, "M") {
		return errors.New("头像大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("avatarExt")
	if extListString == "" {
		return errors.New("找不到头像后缀配置！"), ""
	}
	if !utils.CheckFileExt(avatarHeader, extListString) {
		return errors.New("头像后缀不被允许！"), ""
	}
	createPath := global.GqaServeUpload + "/" + utils.GetConfigBackend("uploadAvatarSavePath") + "/" + username
	createUrl := global.GqaServeAvatar + "/" + username
	filename, err := utils.UploadFile(createPath, avatarHeader)
	if err != nil {
		return err, ""
	}
	avatarUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, avatarUrl
}

func (s *ServiceUpload) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("fileMaxSize")
	if maxSizeString == "" {
		return errors.New("没有找到文件大小配置！"), ""
	}
	if !utils.CheckFileSize(file, maxSizeString, "M") {
		return errors.New("文件大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("fileExt")
	if extListString == "" {
		return errors.New("没有找到文件后缀配置！"), ""
	}
	if !utils.CheckFileExt(fileHeader, extListString) {
		return errors.New("文件后缀不被允许！"), ""
	}
	createPath := global.GqaServeUpload + "/" + utils.GetConfigBackend("uploadFileSavePath") + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	createUrl := global.GqaServeFile + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	filename, err := utils.UploadFile(createPath, fileHeader)
	if err != nil {
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}

func (s *ServiceUpload) UploadBannerImage(img multipart.File, bannerImage *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("bannerImageMaxSize")
	if maxSizeString == "" {
		return errors.New("没有找到网站首页大图大小配置！"), ""
	}
	if !utils.CheckFileSize(img, maxSizeString, "M") {
		return errors.New("网站首页大图大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("bannerImageExt")
	if extListString == "" {
		return errors.New("没有找到网站首页大图后缀配置！"), ""
	}
	if !utils.CheckFileExt(bannerImage, extListString) {
		return errors.New("网站首页大图后缀不被允许！"), ""
	}
	createPath := global.GqaServeUpload + "/" + utils.GetConfigBackend("uploadBannerImageSavePath")
	createUrl := global.GqaServeBanner
	filename, err := utils.UploadFile(createPath, bannerImage)
	if err != nil {
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}

func (s *ServiceUpload) UploadLogo(logo multipart.File, logoHeader *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("logoMaxSize")
	if maxSizeString == "" {
		return errors.New("没有找到网站Logo大小配置！"), ""
	}
	if !utils.CheckFileSize(logo, maxSizeString, "M") {
		return errors.New("网站Logo大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("logoExt")
	if extListString == "" {
		return errors.New("没有找到网站Logo后缀配置！"), ""
	}
	if !utils.CheckFileExt(logoHeader, extListString) {
		return errors.New("网站Logo后缀不被允许！"), ""
	}
	createPath := global.GqaServeUpload + "/" + utils.GetConfigBackend("uploadLogoSavePath")
	createUrl := global.GqaServeLogo
	filename, err := utils.UploadFile(createPath, logoHeader)
	if err != nil {
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}

func (s *ServiceUpload) UploadFavicon(logo multipart.File, logoHeader *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("faviconMaxSize")
	if maxSizeString == "" {
		return errors.New("没有找到favicon大小配置！"), ""
	}
	if !utils.CheckFileSize(logo, maxSizeString, "M") {
		return errors.New("favicon大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("faviconExt")
	if extListString == "" {
		return errors.New("没有找到favicon后缀配置！"), ""
	}
	if !utils.CheckFileExt(logoHeader, extListString) {
		return errors.New("favicon后缀不被允许！"), ""
	}
	createPath := global.GqaServeUpload + "/" + utils.GetConfigBackend("uploadLogoSavePath")
	createUrl := global.GqaServeLogo
	filename, err := utils.UploadFile(createPath, logoHeader)
	if err != nil {
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}
