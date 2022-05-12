<h1>Gin-Quasar-Admin<sup>v2</sup></h1>

<div align=center>
<img src="https://i.loli.net/2020/12/14/cnJoF9r1BXY7Da5.png" width=150" height="150" />
<h1>Gin-Quasar-Admin<sup>v2</sup></h1>
<img src="https://img.shields.io/badge/Quasar-2.6.6-brightgreen"/>                          
<img src="https://img.shields.io/badge/Vue-3.2.33-brightgreen"/>                          
<img src="https://img.shields.io/badge/Gin-1.7.7-brightgreen"/>                              
<img src="https://img.shields.io/badge/Go-1.18.1-brightgreen"/>                  
<img src="https://img.shields.io/badge/License-MIT-brightgreen"/>                                                                 </div>                   

#### 项目使用 Go语言、Gin框架、Vue3(script-setup语法糖)、Quasar2搭建，通过插件形式进行个性化逻辑开发，通过网页端快速配置出个性化网站。V2版本是v1版本的完全重写升级版本，代码更加精简，逻辑更加清晰，欢迎使用！

<div align=center>项目在不断完善中，欢迎Clone、Fork、Issue、PR。如果你感觉不错，麻烦给个小小的 Star 鼓励一下！</div>

***

### 主要功能：

1. 动态配置前台：大图、Logo、Favicon、主次标题、描述、登录页插件等。
2. 动态配置后台：初始密码、验证码、JWT、图片/文件上传允许大小、允许后缀、保存路径等。
3. 用户可以同时拥有多个角色、多个部门。
4. 角色：菜单权限、Api授权、部门数据权限。
5. 部门管理、菜单管理、字典管理、Api管理、角色管理、用户管理、前后台配置管理、日志管理、消息管理、插件生成。
6. i18n、聊天室、打印功能。

### 仓库位置

github：https://github.com/Junvary/gin-quasar-admin

gitee：https://gitee.com/junvary/gin-quasar-admin

### 插件列表

| 序号 |      插件名      |                       仓库                       |                    描述/特点                     |
| :--: | :--------------: | :----------------------------------------------: | :----------------------------------------------: |
|  1   |    系统开发科    |   https://github.com/Junvary/gqa-plugin-xtkfk    | 包含登录页插件和常见增删改查的比较完善的插件样例 |
|  2   |     考勤统计     | https://github.com/Junvary/gqa-plugin-attendance |           在插件中连接并查询其他数据库           |
|  3   | 工会固定资产折旧 |  https://github.com/Junvary/gqa-plugin-AssetGh   |                打印功能、导出功能                |

### 插件开发与接入方式

1. 开发：后端gqaplugin目录、前端plugins目录建立对应文件，开发完成可提取成单独仓库。
2. 接入：clone或下载插件仓库代码，gqaplugin目录中文件夹放入GQA-BACKEND/gqaplugin目录（后端），plugins目录中文件夹放入GQA-FRONTEND/src/pugins目录即可。注：后端可通过github直接导入。

### 项目运行：

> 前端：
>
> ```
> // 进入前端文件夹：GQA-FRONTEND后安装依赖：
> yarn
> // 运行项目：
> yarn run dev
> // 打包：
> yarn run build
> ```

> 后端：
>
> ```go
> # 安装好Go语言环境，后进入GQA-BACKEND文件夹：
> go run main.go
> 
> # 打包（Linux）：
> set GOARCH=amd64
> set GOOS=linux
> go build main.go
> ```
>

### 项目截图（Todo）：

我是一个占位符

### 交流群

QQ群号：342045316

### 鸣谢

<a href="https://www.jetbrains.com/?from=gin-quasar-admin"><img src="https://goframe.org/download/thumbnails/1114119/jetbrains.png" height="120" alt="JetBrains"/></a>

 

### License

Copyright (c) 2021-time.Now()    Junvary

[MIT License](https://github.com/Junvary/gin-quasar-admin/blob/main/LICENSE)
