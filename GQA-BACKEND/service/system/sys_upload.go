package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/utils"
	"mime/multipart"
	"strconv"
	"time"
)

type ServiceUpload struct {
}

func (s *ServiceUpload) UploadAvatar(username string, file *multipart.FileHeader) (err error, avatarUrl string) {
	createPath := global.GqaConfig.Upload.AvatarSavePath + "/" + username
	filename, err := utils.UploadFile(createPath, file)
	if err != nil{
		return err, ""
	}
	avatarUrl = createPath + "/" + filename
	return nil, avatarUrl
}

func (s *ServiceUpload) UploadFile(file *multipart.FileHeader) (err error, avatarUrl string) {
	createPath := global.GqaConfig.Upload.FileSavePath + "/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String()
	filename, err := utils.UploadFile(createPath, file)
	if err != nil{
		return err, ""
	}
	avatarUrl = createPath + "/" + filename
	return nil, avatarUrl
}
