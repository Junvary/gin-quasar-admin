{{.PluginName}}({{.PluginCode}}) 插件模板由 Gin-Quasar-Admin 生成，使用方法如下：
1. 拷贝后台：将 gqaplugin 文件夹中的内容拷贝至：GQA-BACKEND/gqaplugin/
2. 拷贝前台：将 plugins 文件夹中的内容拷贝至：GQA-FRONTEND/src/plugins/
3. 引入插件：找到 GQA-BACKEND/gqaplugin/gqaplugin.go 文件，找到 PluginList 切片，并引入 {{.PluginCode}}.Plugin{{.PluginCode}}
4. 开始后续开发...