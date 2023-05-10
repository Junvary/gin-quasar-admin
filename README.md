<h1>Gin-Quasar-Admin<sup>v2</sup></h1>

[English](README.en.md) | 简体中文 | [Wiki](https://github.com/Junvary/gin-quasar-admin/wiki)

<div align=center>
<img src="https://i.loli.net/2020/12/14/cnJoF9r1BXY7Da5.png" width=150" height="150" />
<h1>Gin-Quasar-Admin<sup>v2</sup></h1>
<img src="https://img.shields.io/badge/Vue-3.2.45-brightgreen"/> 
<img src="https://img.shields.io/badge/Quasar-2.11.10-brightgreen"/>                          
<img src="https://img.shields.io/badge/Go-1.20.2-brightgreen"/>                          
<img src="https://img.shields.io/badge/Gin-1.8.1-brightgreen"/>                              
<img src="https://img.shields.io/badge/Gorm-1.24.0-brightgreen"/>                  
<img src="https://img.shields.io/badge/License-MIT-brightgreen"/>                                                                 </div>










#### 基于Quasar2、Vue3、Vite、 Go、Gin、Gorm搭建的全功能管理系统。

#### 可通过插件形式（和主仓隔离）进行个性化逻辑开发，可通过网页快速配置出个性化网站。

#### 项目在不断完善中，欢迎Clone、Fork、Issue、PR(to dev)。

#### 如果你感觉不错，麻烦给个小小的 Star 鼓励一下！

***

### 交流群

QQ群号：342045316

### 项目截图：

![white1.png](https://s2.loli.net/2022/12/01/kEbBwuLi37VlGcF.png)

![dark1.png](https://s2.loli.net/2022/12/01/feywBOXFDRgk9rY.png)

![white2.png](https://s2.loli.net/2022/12/01/oNSz8dYDqFZCRxI.png)

![dark2.png](https://s2.loli.net/2022/12/01/AJL7zjm9RiG6fQr.png)


### 在线体验

1. 地址：http://81.68.159.232/
2. 用户名：admin
3. 密码：123456
4. 为了他人方便，请不要修改密码等信息。

### 主要功能：

1. 动态配置前台：大图、Logo、Favicon、主次标题、描述、登录页插件等。
2. 动态配置后台：初始密码、验证码、JWT、图片/文件上传允许大小、允许后缀、保存路径等。
3. 用户可以同时拥有多个角色、多个部门。
4. 角色：菜单权限、Api授权、部门数据权限。
5. 部门管理、菜单管理、字典管理、Api管理、角色管理、用户管理、前后台配置管理、日志管理、消息管理、插件生成、在线用户。
6. i18n、聊天室、打印功能、踢出用户。

### 仓库位置

github：https://github.com/Junvary/gin-quasar-admin

gitee：https://gitee.com/junvary/gin-quasar-admin

### 插件开发与接入方式

1. 开发：后端gqaplugin目录、前端plugins目录建立对应文件，开发完成可提取成单独仓库（或通过插件生成相关目录手动移入）。
2. 接入：clone或下载插件仓库代码，gqaplugin目录中文件夹放入GQA-BACKEND/gqaplugin目录（后端），plugins目录中文件夹放入GQA-FRONTEND/src/pugins目录即可。注：后端可通过github直接导入，部分插件需按插件说明引入。

### 项目运行：

> 前端：
> 
> ```
> // 进入前端文件夹：GQA-FRONTEND后安装依赖：
> yarn
> // 运行项目：
> yarn dev
> // 打包：
> yarn build
> ```

> 后端：
>
> ```go
> # 安装好Go语言环境，后进入GQA-BACKEND文件夹：
> go mod tidy
> go run main.go
> 
> # 打包（Linux）：
> set GOARCH=amd64
> set GOOS=linux
> go build main.go
> ```

### 鸣谢

<a href="https://www.jetbrains.com/?from=gin-quasar-admin"><img src="https://goframe.org/download/thumbnails/1114119/jetbrains.png" height="120" alt="JetBrains"/></a>

### License

Copyright (c) 2021-time.Now()    Junvary

[MIT License](https://github.com/Junvary/gin-quasar-admin/blob/main/LICENSE)
