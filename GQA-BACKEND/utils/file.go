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
	// Read file suffix
	ext := path.Ext(fileHeader.Filename)

	// Read file name and encrypt
	name := strings.TrimSuffix(fileHeader.Filename, ext)
	name = EncodeMD5(name)

	// Splice new file names
	filename = name + "_" + time.Now().Format("20060102150405") + ext

	// Create the directory of upload/username
	if err = os.MkdirAll(createPath, os.ModePerm); err != nil {
		return "", err
	}

	// Splicing path and file name
	filepath := createPath + "/" + filename

	// read file
	f, openError := fileHeader.Open()
	if err != nil {
		return "", openError
	}
	defer f.Close()

	// create dir
	out, createErr := os.Create(filepath)
	if createErr != nil {
		return "", createErr
	}
	defer out.Close()

	// copy file
	if _, err = io.Copy(out, f); err != nil {
		return "", err
	}

	return filename, nil
}

func CheckFileSize(file multipart.File, maxSizeString string, unit string) bool {
	content, _ := ioutil.ReadAll(file)
	size := len(content)
	maxSize, _ := strconv.Atoi(maxSizeString)
	if unit == "M" {
		// ==> M
		if size > maxSize*1024*1024 {
			return false
		}
	} else {
		// ==> K
		if size > maxSize*1024 {
			return false
		}
	}
	return true
}

func CheckFileExt(fileHeader *multipart.FileHeader, extListString string) bool {
	ext := path.Ext(fileHeader.Filename)
	extList := strings.Split(extListString, ",")
	extMap := make(map[string]struct{}, len(extList))
	for _, v := range extList {
		extMap[v] = struct{}{}
	}
	_, ok := extMap[ext]
	return ok
}
