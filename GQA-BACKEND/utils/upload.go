package utils

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func UploadFile(createPath string, fileHeader *multipart.FileHeader) (filename string, err error) {
	// 读取文件后缀
	ext := path.Ext(fileHeader.Filename)

	// 读取文件名并加密
	name := strings.TrimSuffix(fileHeader.Filename, ext)
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
	f, openError := fileHeader.Open()
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

func CheckFileSize(file multipart.File, maxSizeString string) bool {
	content, _ := ioutil.ReadAll(file)
	size := len(content)
	maxSize, _ := strconv.Atoi(maxSizeString)
	// 转换成M
	if size > maxSize*1024*1024 {
		return false
	}
	return true
}

func CheckFileExt(fileHeader *multipart.FileHeader, extListString string)bool  {
	ext := path.Ext(fileHeader.Filename)
	extList := strings.Split(extListString, ",")
	extMap := make(map[string]struct{}, len(extList))
	for _, v := range extList{
		extMap[v] = struct{}{}
	}
	_, ok := extMap[ext]
	return ok
}
