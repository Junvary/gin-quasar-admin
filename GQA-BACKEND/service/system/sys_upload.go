package system

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"mime/multipart"
	"strconv"
	"time"
)

type ServiceUpload struct {
}

func (s *ServiceUpload) UploadAvatar(username string, avatar multipart.File, avatarHeader *multipart.FileHeader) (err error, avatarUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("avatarMaxSize")
	if maxSizeString == ""{
		return errors.New("找不到头像大小配置！"), ""
	}
	if !utils.CheckFileSize(avatar, maxSizeString, "M"){
		return errors.New("头像大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("avatarExt")
	if extListString == ""{
		return errors.New("找不到头像后缀配置！"), ""
	}
	if !utils.CheckFileExt(avatarHeader, extListString){
		return errors.New("头像后缀不被允许！"), ""
	}
	createPath := global.GqaConfig.Upload.AvatarSavePath + "/" + username
	createUrl := global.GqaConfig.Upload.AvatarUrl + "/" + username
	filename, err := utils.UploadFile(createPath, avatarHeader)
	if err != nil{
		return err, ""
	}
	avatarUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, avatarUrl
}

func (s *ServiceUpload) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("fileMaxSize")
	if maxSizeString == ""{
		return errors.New("没有找到文件大小配置！"), ""
	}
	if !utils.CheckFileSize(file, maxSizeString, "M"){
		return errors.New("文件大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("fileExt")
	if extListString == ""{
		return errors.New("没有找到文件后缀配置！"), ""
	}
	if !utils.CheckFileExt(fileHeader, extListString){
		return errors.New("文件后缀不被允许！"), ""
	}
	createPath := global.GqaConfig.Upload.FileSavePath + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	createUrl := global.GqaConfig.Upload.FileUrl  + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	filename, err := utils.UploadFile(createPath, fileHeader)
	if err != nil{
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}

func (s *ServiceUpload) UploadBannerImage(img multipart.File, bannerImage *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("bannerImageMaxSize")
	if maxSizeString == ""{
		return errors.New("没有找到网站首页大图大小配置！"), ""
	}
	if !utils.CheckFileSize(img, maxSizeString, "M"){
		return errors.New("网站首页大图大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("bannerImageExt")
	if extListString == ""{
		return errors.New("没有找到网站首页大图后缀配置！"), ""
	}
	if !utils.CheckFileExt(bannerImage, extListString){
		return errors.New("网站首页大图后缀不被允许！"), ""
	}
	createPath := global.GqaConfig.Upload.WebLogoSavePath
	createUrl := global.GqaConfig.Upload.WebLogoUrl
	filename, err := utils.UploadFile(createPath, bannerImage)
	if err != nil{
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}

func (s *ServiceUpload) UploadWebLogo(logo multipart.File, logoHeader *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("webLogoMaxSize")
	if maxSizeString == ""{
		return errors.New("没有找到网站Logo大小配置！"), ""
	}
	if !utils.CheckFileSize(logo, maxSizeString, "M"){
		return errors.New("网站Logo大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("webLogoExt")
	if extListString == ""{
		return errors.New("没有找到网站Logo后缀配置！"), ""
	}
	if !utils.CheckFileExt(logoHeader, extListString){
		return errors.New("网站Logo后缀不被允许！"), ""
	}
	createPath := global.GqaConfig.Upload.WebLogoSavePath
	createUrl := global.GqaConfig.Upload.WebLogoUrl
	filename, err := utils.UploadFile(createPath, logoHeader)
	if err != nil{
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}

func (s *ServiceUpload) UploadHeaderLogo(logo multipart.File, logoHeader *multipart.FileHeader) (err error, fileUrl string) {
	// 检查文件大小
	maxSizeString := utils.GetConfigBackend("headerLogoMaxSize")
	if maxSizeString == ""{
		return errors.New("没有找到标签页Logo大小配置！"), ""
	}
	if !utils.CheckFileSize(logo, maxSizeString, "M"){
		return errors.New("标签页Logo大小超出限制！"), ""
	}
	// 检查文件后缀
	extListString := utils.GetConfigBackend("headerLogoExt")
	if extListString == ""{
		return errors.New("没有找到标签页Logo后缀配置！"), ""
	}
	if !utils.CheckFileExt(logoHeader, extListString){
		return errors.New("标签页Logo后缀不被允许！"), ""
	}
	createPath := global.GqaConfig.Upload.WebLogoSavePath
	createUrl := global.GqaConfig.Upload.WebLogoUrl
	filename, err := utils.UploadFile(createPath, logoHeader)
	if err != nil{
		return err, ""
	}
	fileUrl = "gqa-upload:" + createUrl + "/" + filename
	return nil, fileUrl
}
