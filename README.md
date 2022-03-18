# 欢迎使用Gin-Quasar-Admin！
<div align=center>
<img src="https://i.loli.net/2020/12/14/cnJoF9r1BXY7Da5.png" width=200" height="200" />
<h1>Gin-Quasar-Admin</h1>
</div>
<div align=center>
<img src="https://img.shields.io/badge/quasar-2.6.1-brightgreen"/>
<img src="https://img.shields.io/badge/vue-3.2.31-brightgreen"/>
<img src="https://img.shields.io/badge/gin-1.7.7-brightgreen"/>
<img src="https://img.shields.io/badge/golang-1.17.2-brightgreen"/>
</div>






```
升级新版本Quasar后，如遇 Cannot find module 'autoprefixer' 错误，请手动删除yarn.lock文件和node_modules文件夹后，重新执行yarn安装依赖！
```




#### Gin-Quasar-Admin 使用 Go+Gin+Vue3+Quasar2 搭建，可以通过简单的配置快速搭建个性化网站，并通过插件形式（可生成模板）开发个性化逻辑。项目实现和集成了众多功能，如用户管理、部门管理、角色管理、数据权限、Websocket、聊天室、消息管理、菜单管理、字典管理、i18n等。是Quasar、Go的入门和学习佳品。

### 主要功能：

1. 动态配置网站浏览器标签页图标。
2. 动态配置前台：主标题、次标题、网站描述、登录页插件等。
3. 动态配置后台：初始密码、验证码、JWT、图片/文件上传允许大小、允许后缀等。
4. 用户可以同时拥有多个角色。
5. 用户可以同时拥有多个部门。
6. 角色授权菜单权限。
7. 角色授权API权限。
8. 角色授权数据权限。
9. 部门管理。
10. 字典管理。
11. 用户管理。
12. 菜单管理。
13. i18n。
14. 聊天室。
15. 消息管理。
16. 待处理消息和待办便签提醒。
16. 插件模板生成。

#### 项目在不断完善中，欢迎clone和fork试用。如果你感觉不错，麻烦给个小小的 Star 鼓励一下！

github：https://github.com/Junvary/gin-quasar-admin

gitee：https://gitee.com/junvary/gin-quasar-admin

***

### 项目运行：

> 前端：
>
> ```
> // 安装Quasar-cli:
> yarn global add @quasar/cli
> 
> // 进入前端文件夹：GQA-FRONTEND后安装依赖：
> yarn
> 
> // 运行项目：
> quasar dev
> 
> // 打包：
> quasar build
> ```



> 后端：
>
> ```go
> # 安装好Go语言环境，后进入GQA-BACKEND文件夹：
> go run main.go
> 
> # 打包：
> set GOARCH=amd64
> set GOOS=linux
> go build main.go
> ```
>
> 

### 项目截图：

<div align=center>
    <img src="https://github.com/Junvary/gin-quasar-admin/blob/dev/img/1.png" /><br/><br/>
    <img src="https://s2.loli.net/2021/12/06/vNYn7UpcHOowdlV.png" /><br/><br/>
    <img src="https://github.com/Junvary/gin-quasar-admin/blob/dev/img/2.png" /><br/><br/>
    <img src="https://s2.loli.net/2021/12/06/5lmnMOfixvhb3Q7.png" /><br/><br/>
    <img src="https://github.com/Junvary/gin-quasar-admin/blob/dev/img/3.png" /><br/><br/>
    <img src="https://s2.loli.net/2021/12/06/BRKo4mLf1OQsDFP.png" /><br/><br/>
    <img src="https://github.com/Junvary/gin-quasar-admin/blob/dev/img/4.png" /><br/><br/>
    <img src="https://s2.loli.net/2021/12/06/xeSPZN8q1lVLYgj.png" /><br/><br/>
    <img src="https://github.com/Junvary/gin-quasar-admin/blob/dev/img/5.png" /><br/><br/>
    <img src="https://s2.loli.net/2021/12/06/IYxVW7PEmyRtM2J.png" /><br/><br/>
</div>






### 最后：

任何项目、框架都各有千秋，如果有了自己的一些想法，那么就付诸于实际吧。

生命在于折腾，乐趣在于想法的实现。

鸣谢 https://github.com/flipped-aurora/gin-vue-admin 

 
