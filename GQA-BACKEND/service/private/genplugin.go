package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	genPath = "gqa-gen-plugin/"
)

type ServiceGenPlugin struct{}

type tplData struct {
	template      *template.Template
	locationPath  string
	genGoFilePath string
}

func (s *ServiceGenPlugin) GenPlugin(genPluginStruct *model.SysGenPlugin) (err error) {
	dataList, fileList, makeDirList, err := s.PrepareData(genPluginStruct)
	if err != nil {
		return err
	}
	if err = utils.CheckAndCreatePath(makeDirList...); err != nil {
		return err
	}
	// 生成文件
	for _, value := range dataList {
		f, err := os.OpenFile(value.genGoFilePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, genPluginStruct); err != nil {
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(genPath); err != nil {
			return
		}
	}()
	if err = utils.ZipFiles("./gqa-gen-plugin.zip", fileList, ".", "."); err != nil {
		return err
	}
	return nil
}

func (s *ServiceGenPlugin) PrepareData(genPluginStruct *model.SysGenPlugin) (dataList []tplData, fileList []string, makeDirList []string, err error) {
	var basePath = global.GqaConfig.System.GenPluginPath
	tplFileList, err := s.GetAllTplFile(basePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	makeDirList = make([]string, 0, len(tplFileList))

	for _, value := range tplFileList {
		temp, err := template.ParseFiles(value)
		if err != nil {
			return nil, nil, nil, err
		}
		dataList = append(dataList, tplData{locationPath: value, template: temp})
	}
	//basePath: template/gqaplugintemplate
	//locationPath: 如 template/gqaplugintemplate/gqaplugin/main.go.tpl
	for index, value := range dataList {
		// tplInTemplatePath 类似 gqaplugin/api/privateapi/privateapi.go.tpl
		// tplInTemplatePath 类似 gqaplugin/main.go.tpl
		// tplInTemplatePath 类似 readme.txt.tpl
		tplInTemplatePath := strings.TrimPrefix(value.locationPath, basePath+"/")

		if tplInTemplatePath == "readme.txt.tpl" {
			//创建到打包文件夹的第一层目录中
			dataList[index].genGoFilePath = filepath.Join(genPath, "help.txt")
			continue
		}

		tplInTemplatePath = tplInTemplatePath[:strings.Index(tplInTemplatePath, "/")] + "/" +
			genPluginStruct.PluginCode + "/" +
			tplInTemplatePath[strings.Index(tplInTemplatePath, "/"):]

		if lastSeparator := strings.LastIndex(tplInTemplatePath, "/"); lastSeparator != -1 {
			// trueFileName 类似 privateapi.go
			// trueFileName 类似 main.go
			trueFileName := strings.TrimSuffix(tplInTemplatePath[lastSeparator+1:], ".tpl")
			firstDotIndex := strings.Index(trueFileName, ".")
			if firstDotIndex != -1 {
				dataList[index].genGoFilePath = filepath.Join(genPath, tplInTemplatePath[:lastSeparator], trueFileName)
			}
		}

		if lastSeparator := strings.LastIndex(dataList[index].genGoFilePath, string(os.PathSeparator)); lastSeparator != -1 {
			makeDirList = append(makeDirList, dataList[index].genGoFilePath[:lastSeparator])
		}
	}
	for _, value := range dataList {
		fileList = append(fileList, value.genGoFilePath)
	}
	return dataList, fileList, makeDirList, err
}

func (s *ServiceGenPlugin) GetAllTplFile(basePath string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(basePath)
	for _, file := range files {
		if file.IsDir() {
			fileList, err = s.GetAllTplFile(basePath+"/"+file.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(file.Name(), ".tpl") {
				fileList = append(fileList, basePath+"/"+file.Name())
			}
		}
	}
	return fileList, err
}
