“{{.PluginName}}” “{{.PluginCode}}” 插件模板由 Gin-Quasar-Admin 生成，使用方法如下：
1. 后台：将 gqaplugin 文件夹中的内容拷贝至：GQA-BACKEND/gqaplugin/
2. 前台：将 Plugin 文件夹中的内容拷贝至：GQA-FRONTEND/src/pages/Plugin/
3. 前往 GQA-BACKEND/gqaplugin/gqa_plugin.go 文件
4. 找到 PluginList 切片，并引入 {{.PluginCode}}.Plugin{{.PluginCode}},
5. 后续开发