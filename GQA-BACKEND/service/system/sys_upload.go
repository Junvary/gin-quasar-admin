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
	var avatarConfig system.SysConfig
	err = global.GqaDb.First(&avatarConfig, "gqa_option = ?", "avatarMaxSize").Error
	if err != nil{
		return err, ""
	}
	var maxSizeString string
	if avatarConfig.Custom != ""{
		maxSizeString = avatarConfig.Custom
	}else {
		maxSizeString = avatarConfig.Default
	}
	if !utils.CheckFileSize(avatar, maxSizeString){
		return errors.New("头像大小超出限制！"), ""
	}
	createPath := global.GqaConfig.Upload.AvatarSavePath + "/" + username
	filename, err := utils.UploadFile(createPath, avatarHeader)
	if err != nil{
		return err, ""
	}
	avatarUrl = createPath + "/" + filename
	return nil, avatarUrl
}

func (s *ServiceUpload) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (err error, avatarUrl string) {
	var fileConfig system.SysConfig
	err = global.GqaDb.First(&fileConfig, "gqa_option = ?", "fileMaxSize").Error
	if err != nil{
		return err, ""
	}
	var maxSizeString string
	if fileConfig.Custom != ""{
		maxSizeString = fileConfig.Custom
	}else {
		maxSizeString = fileConfig.Default
	}
	if !utils.CheckFileSize(file, maxSizeString){
		return errors.New("文件大小超出限制！"), ""
	}
	createPath := global.GqaConfig.Upload.FileSavePath + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	filename, err := utils.UploadFile(createPath, fileHeader)
	if err != nil{
		return err, ""
	}
	avatarUrl = createPath + "/" + filename
	return nil, avatarUrl
}
