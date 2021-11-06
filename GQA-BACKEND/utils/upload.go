package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func UploadFile(createPath string, file *multipart.FileHeader) (filename string, err error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)

	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = EncodeMD5(name)

	// 拼接新文件名
	filename = name + "_" + time.Now().Format("20060102150405") + ext

	// 创建upload/username的目录
	if err = os.MkdirAll(createPath, os.ModePerm); err != nil {
		return "", err
	}

	// 拼接路径和文件名
	filepath := createPath + "/" + filename

	// 读取文件
	f, openError := file.Open()
	if err != nil {
		return "", openError
	}
	defer f.Close()

	// 创建目录
	out, createErr := os.Create(filepath)
	if createErr != nil {
		return "", createErr
	}
	defer out.Close()

	// 拷贝文件
	if _, err = io.Copy(out, f); err != nil {
		return "", err
	}

	return filename, nil
}

