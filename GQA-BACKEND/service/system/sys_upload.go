package system

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"mime/multipart"
	"strconv"
	"time"
)

type ServiceUpload struct {
}

func (s *ServiceUpload) UploadAvatar(username string, avatar multipart.File, avatarHeader *multipart.FileHeader) (err error, avatarUrl string) {
	// 检查文件大小
	var avatarSizeConfig system.SysConfig
	err = global.GqaDb.First(&avatarSizeConfig, "gqa_option = ?", "avatarMaxSize").Error
	if err != nil{
		return err, ""
	}
	var maxSizeString string
	if avatarSizeConfig.Custom != ""{
		maxSizeString = avatarSizeConfig.Custom
	}else {
		maxSizeString = avatarSizeConfig.Default
	}
	if !utils.CheckFileSize(avatar, maxSizeString){
		return errors.New("头像大小超出限制！"), ""
	}
	// 检查文件后缀
	var avatarExtConfig system.SysConfig
	err = global.GqaDb.First(&avatarExtConfig, "gqa_option = ?", "avatarExt").Error
	if err != nil{
		return err, ""
	}
	var extListString string
	if avatarExtConfig.Custom != ""{
		extListString = avatarExtConfig.Custom
	}else {
		extListString = avatarExtConfig.Default
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
	var fileSizeConfig system.SysConfig
	err = global.GqaDb.First(&fileSizeConfig, "gqa_option = ?", "fileMaxSize").Error
	if err != nil{
		return err, ""
	}
	var maxSizeString string
	if fileSizeConfig.Custom != ""{
		maxSizeString = fileSizeConfig.Custom
	}else {
		maxSizeString = fileSizeConfig.Default
	}
	if !utils.CheckFileSize(file, maxSizeString){
		return errors.New("文件大小超出限制！"), ""
	}
	// 检查文件后缀
	var fileExtConfig system.SysConfig
	err = global.GqaDb.First(&fileExtConfig, "gqa_option = ?", "fileExt").Error
	if err != nil{
		return err, ""
	}
	var extListString string
	if fileExtConfig.Custom != ""{
		extListString = fileExtConfig.Custom
	}else {
		extListString = fileExtConfig.Default
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
