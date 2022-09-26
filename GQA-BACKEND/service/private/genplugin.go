package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
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
	// 按模板准备tpl文件
	dataList, err := s.ChangeToTplData(genPluginStruct)
	if err != nil {
		return err
	}
	// 转换前后台文件，创建目录
	dataListNew, fileListNew, makeDirListNew := s.GenDataNew(genPluginStruct, dataList)
	if err = utils.CheckAndCreatePath(makeDirListNew...); err != nil {
		return err
	}
	// 生成前后端文件
	if err = s.GenFrontendAndBackendFile(genPluginStruct, dataListNew); err != nil {
		return err
	}
	defer func() { // 移除中间文件
		if err := os.RemoveAll(genPath); err != nil {
			return
		}
	}()
	if err = utils.ZipFiles("./gqa-gen-plugin.zip", fileListNew, ".", "."); err != nil {
		return err
	}
	return nil
}

func (s *ServiceGenPlugin) ChangeToTplData(genPluginStruct *model.SysGenPlugin) (dataList []tplData, err error) {
	var basePath = global.GqaConfig.System.GenPluginPath
	tplFileList, err := s.GetAllTplFile(basePath, nil)
	if err != nil {
		return nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	for _, value := range tplFileList {
		temp, err := template.ParseFiles(value)
		if err != nil {
			return nil, err
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
			dataList[index].genGoFilePath = filepath.Join(genPath, "readme.txt")
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
	}
	return dataList, err
}

func (s *ServiceGenPlugin) GetAllTplFile(basePath string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(basePath)
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

func (s *ServiceGenPlugin) GenDataNew(genPluginStruct *model.SysGenPlugin, dataList []tplData) (dataListNew []tplData, fileListNew []string, makeDirListNew []string) {
	for _, value := range dataList {
		if value.locationPath == "template/gqaplugintemplate/plugins/index.vue.tpl" {
			for _, mo := range genPluginStruct.PluginModel {
				var ggfp = "gqa-gen-plugin\\plugins\\" + genPluginStruct.PluginCode + "\\" + mo.ModelName
				dataListNew = append(dataListNew, tplData{
					template:      value.template,
					locationPath:  "template/gqaplugintemplate/plugins/index.vue.tpl",
					genGoFilePath: ggfp + "\\index.vue",
				})
				fileListNew = append(fileListNew, ggfp+"\\index.vue")
				makeDirListNew = append(makeDirListNew, ggfp)
			}
		} else if value.locationPath == "template/gqaplugintemplate/plugins/modules/recordDetail.vue.tpl" {
			for _, mo := range genPluginStruct.PluginModel {
				var ggfp = "gqa-gen-plugin\\plugins\\" + genPluginStruct.PluginCode + "\\" + mo.ModelName
				dataListNew = append(dataListNew, tplData{
					template:      value.template,
					locationPath:  "template/gqaplugintemplate/plugins/modules/recordDetail.vue.tpl",
					genGoFilePath: ggfp + "\\modules\\recordDetail.vue",
				})
				fileListNew = append(fileListNew, ggfp+"\\modules\\recordDetail.vue")
				makeDirListNew = append(makeDirListNew, ggfp+"\\modules")
			}
		} else {
			dataListNew = append(dataListNew, value)
			fileListNew = append(fileListNew, value.genGoFilePath)
			if value.locationPath != "template/gqaplugintemplate/readme.txt.tpl" {
				if lastSeparator := strings.LastIndex(value.genGoFilePath, string(os.PathSeparator)); lastSeparator != -1 {
					makeDirListNew = append(makeDirListNew, value.genGoFilePath[:lastSeparator])
				}
			}
		}
	}
	return dataListNew, fileListNew, makeDirListNew
}

func (s *ServiceGenPlugin) GenFrontendAndBackendFile(genPluginStruct *model.SysGenPlugin, dataListNew []tplData) (err error) {
	var dataListNewFrontend []tplData
	var dataListNewBackend []tplData
	for _, value := range dataListNew {
		if strings.HasPrefix(value.locationPath, "template/gqaplugintemplate/plugins/") {
			dataListNewFrontend = append(dataListNewFrontend, value)
		} else {
			dataListNewBackend = append(dataListNewBackend, value)
		}
	}
	for _, value := range dataListNewBackend {
		f, err := os.OpenFile(value.genGoFilePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			_ = f.Close()
			return err
		}
		if err = value.template.Execute(f, genPluginStruct); err != nil {
			_ = f.Close()
			return err
		}
		_ = f.Close()
	}

	for i, value := range dataListNewFrontend {
		f, err := os.OpenFile(value.genGoFilePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			_ = f.Close()
			return err
		}
		type dataStruct struct {
			PluginCode  string
			PluginModel interface{}
		}
		var ds = dataStruct{
			PluginCode:  genPluginStruct.PluginCode,
			PluginModel: genPluginStruct.PluginModel[i/2],
		}
		if err = value.template.Execute(f, ds); err != nil {
			_ = f.Close()
			return err
		}
		_ = f.Close()
	}
	return err
}
