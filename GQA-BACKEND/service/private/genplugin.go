package private

import (
	"errors"
	elDir "github.com/Junvary/erleng/dir"
	elPath "github.com/Junvary/erleng/path"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type ServiceGenPlugin struct{}

type tplData struct {
	template     *template.Template
	templatePath string
	genFilePath  string
}

func (s *ServiceGenPlugin) GetGenPluginList(getDataList model.RequestGetGenPluginList) (err error, data interface{}, total int64) {
	pageSize := getDataList.PageSize
	offset := getDataList.PageSize * (getDataList.Page - 1)
	db := global.GqaDb.Model(&model.SysGenPluginList{})
	var dataList []model.SysGenPluginList
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(getDataList.SortBy, getDataList.Desc)).Find(&dataList).Error
	return err, dataList, total
}

func (s *ServiceGenPlugin) GenPlugin(genPluginStruct *model.SysGenPlugin) (err error) {
	// 创建生成插件目录
	var genToPath = global.GqaConfig.System.GenPluginToPath
	if err = elPath.CreatePath(genToPath); err != nil {
		return err
	}
	// 临时目录名
	pluginCodeTimeDir := genPluginStruct.PluginCode + "-" + time.Now().Format("20060102150405")
	// 按模板准备对应 PluginCode 的 tplData文件List
	tplDataList, err := s.PrepareTplData(genPluginStruct, pluginCodeTimeDir)
	if err != nil {
		return err
	}
	// 转换前后台文件和需要创建的目录
	dataListNew, genFileList, makeDirList := s.GenData(genPluginStruct, tplDataList, pluginCodeTimeDir)
	if err = elPath.CreatePath(makeDirList...); err != nil {
		return err
	}
	// 生成前后端文件
	if err = s.GenFrontendAndBackendFile(genPluginStruct, dataListNew); err != nil {
		return err
	}
	defer func() { // 移除中间文件
		if err := os.RemoveAll(path.Join(genToPath, pluginCodeTimeDir)); err != nil {
			return
		}
	}()
	zipPath := genToPath + "/gqa-gen-" + pluginCodeTimeDir + ".zip"
	if err = utils.ZipFiles(zipPath, genFileList, "gqagen"+string(os.PathSeparator)+"plugins", ""); err != nil {
		return err
	}
	toGenRecord := model.SysGenPluginList{
		PluginSort: genPluginStruct.PluginSort,
		PluginCode: genPluginStruct.PluginCode,
		PluginName: genPluginStruct.PluginName,
		PluginFile: zipPath,
	}
	err = global.GqaDb.Create(&toGenRecord).Error
	return err
}

func (s *ServiceGenPlugin) PrepareTplData(genPluginStruct *model.SysGenPlugin, pluginCodeTimeDir string) (dataList []tplData, err error) {
	var genToPath = global.GqaConfig.System.GenPluginToPath
	var templatePath = global.GqaConfig.System.GenPluginTemplatePath
	tplFileList, err := elDir.GetAllFilesInDirWithSuffix(templatePath, ".tpl", nil)
	if err != nil {
		return nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	for _, value := range tplFileList {
		temp, err := template.ParseFiles(value)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, tplData{templatePath: value, template: temp})
	}
	for index, value := range dataList {
		// tplInTemplatePath: template/genplugin/gqaplugin/api/privateapi/privateapi.go.tpl 去掉 template/genplugin/ 前缀
		// 得到具体的 gqaplugin 或者 plugins开头的路径或者readme.txt.tpl
		tplInTemplatePath := strings.TrimPrefix(value.templatePath, templatePath+"/")
		// 如果是说明文档，使 tplData中的genFilePath 变成 gqagen/plugins/[PluginCode]/readme.txt
		if tplInTemplatePath == "readme.txt.tpl" {
			//创建到打包文件夹的第一层目录中
			dataList[index].genFilePath = filepath.Join(genToPath, pluginCodeTimeDir, "readme.txt")
			//结束本次循环，执行下一次
			continue
		}
		// 如果是其他，使 tplData中的genFilePath 变成 gqagen/plugins/[PluginCode]/...
		if lastSeparator := strings.LastIndex(tplInTemplatePath, "/"); lastSeparator != -1 {
			// trueFileName 类似 privateapi.go
			trueFileName := strings.TrimSuffix(tplInTemplatePath[lastSeparator+1:], ".tpl")
			firstDotIndex := strings.Index(trueFileName, ".")
			if firstDotIndex != -1 {
				// 将此文件的 genFilePath 填为: gqagen\plugins\[PluginCode]\gqaplugin\api\privateapi\privateapi.go
				dataList[index].genFilePath = filepath.Join(genToPath, pluginCodeTimeDir, tplInTemplatePath[:lastSeparator], trueFileName)
			}
		}
	}
	return dataList, err
}

func (s *ServiceGenPlugin) GenData(genPluginStruct *model.SysGenPlugin, tplDataList []tplData, pluginCodeTimeDir string) (dataListNew []tplData, genFileList []string, makeDirList []string) {
	for _, value := range tplDataList {
		if value.templatePath == "template/genplugin/plugins/index.vue.tpl" {
			for _, pm := range genPluginStruct.PluginModel {
				lastIndex := strings.LastIndex(value.genFilePath, string(os.PathSeparator))
				var pmGen = value.genFilePath[:lastIndex] + string(os.PathSeparator) + pm.ModelName
				dataListNew = append(dataListNew, tplData{
					template:     value.template,
					templatePath: value.templatePath,
					genFilePath:  pmGen + value.genFilePath[lastIndex:],
				})
				genFileList = append(genFileList, pmGen+value.genFilePath[lastIndex:])
				makeDirList = append(makeDirList, pmGen)
			}
		} else if value.templatePath == "template/genplugin/plugins/modules/recordDetail.vue.tpl" {
			for _, pm := range genPluginStruct.PluginModel {
				lastIndex := strings.LastIndex(value.genFilePath, string(os.PathSeparator))
				var pmGen1 = value.genFilePath[:lastIndex]
				var pmGen2 = value.genFilePath[lastIndex:]
				lastIndex2 := strings.LastIndex(pmGen1, string(os.PathSeparator))
				var pmGen3 = pmGen1[:lastIndex2]
				var pmGen4 = pmGen1[lastIndex2:]
				var pmGen = pmGen3 + string(os.PathSeparator) + pm.ModelName
				dataListNew = append(dataListNew, tplData{
					template:     value.template,
					templatePath: value.templatePath,
					genFilePath:  pmGen + pmGen4 + pmGen2,
				})
				genFileList = append(genFileList, pmGen+pmGen4+pmGen2)
				makeDirList = append(makeDirList, pmGen+pmGen4)
			}
		} else {
			// go的情况
			dataListNew = append(dataListNew, value)
			genFileList = append(genFileList, value.genFilePath)
			// readme的情况
			if value.templatePath != "template/genplugin/readme.txt.tpl" {
				if lastSeparator := strings.LastIndex(value.genFilePath, string(os.PathSeparator)); lastSeparator != -1 {
					makeDirList = append(makeDirList, value.genFilePath[:lastSeparator])
				}
			}
		}
	}
	return dataListNew, genFileList, makeDirList
}

func (s *ServiceGenPlugin) GenFrontendAndBackendFile(genPluginStruct *model.SysGenPlugin, dataListNew []tplData) (err error) {
	var dataListNewFrontend []tplData
	var dataListNewBackend []tplData
	// 将前后端的文件分开
	for _, value := range dataListNew {
		if strings.HasPrefix(value.templatePath, "template/genplugin/plugins/") {
			dataListNewFrontend = append(dataListNewFrontend, value)
		} else {
			dataListNewBackend = append(dataListNewBackend, value)
		}
	}
	for _, value := range dataListNewBackend {
		f, err := os.OpenFile(value.genFilePath, os.O_CREATE|os.O_WRONLY, 0o755)
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
		f, err := os.OpenFile(value.genFilePath, os.O_CREATE|os.O_WRONLY, 0o755)
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

func (s *ServiceGenPlugin) DeleteGenPluginById(id uint) (err error) {
	var data model.SysGenPluginList
	if data.Stable == "yesNo_yes" {
		return errors.New(utils.GqaI18n("StableCantDo") + data.PluginCode)
	}
	if err = global.GqaDb.Where("id = ?", id).First(&data).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&data).Error
	_ = os.RemoveAll(data.PluginFile)
	return err
}

func (s *ServiceGenPlugin) DownloadGenPluginById(id uint) (err error, plugin model.SysGenPluginList) {
	var data model.SysGenPluginList
	err = global.GqaDb.Where("id = ?", id).First(&data).Error
	return err, data
}
