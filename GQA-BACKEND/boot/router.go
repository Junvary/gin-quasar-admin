package boot

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/middleware"
	"gin-quasar-admin/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	var Router = gin.Default()

	// 为文件提供静态资源路径 头像和文件 无须鉴权
	Router.StaticFS(global.GqaConfig.Upload.AvatarUrl, http.Dir(global.GqaConfig.Upload.AvatarSavePath))
	Router.StaticFS(global.GqaConfig.Upload.FileUrl, http.Dir(global.GqaConfig.Upload.FileSavePath))

	/*
		公共路由分组：以 public 开头，路由内部无须再次分组，无须鉴权。
	*/
	PublicGroup := Router.Group("public")
	// router目录下的public文件夹，提取为本分组
	routerPublic := router.GroupRouterApp.RouterPublic
	{
		// 检查数据库是否需要初始化，个性化配置和初始化数据库
		routerPublic.InitRouterCheckAndInitDb(PublicGroup)
		// 验证码
		routerPublic.InitRouterCaptcha(PublicGroup)
		// 登录
		routerPublic.InitRouterLogin(PublicGroup)
		// 所有字典列表
		routerPublic.InitRouterDictDetail(PublicGroup)
	}
	/*
		鉴权路由分组：这里以空分组，路由内部按实绩情况分组，需要鉴权。
	*/
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JwtHandler()).Use(middleware.CasbinHandler())
	routerSystem := router.GroupRouterApp.RouterSystem
	{
		routerSystem.InitRouterMenu(PrivateGroup)
		routerSystem.InitRouterUser(PrivateGroup)
		routerSystem.InitRouterRole(PrivateGroup)
		routerSystem.InitRouterDept(PrivateGroup)
		routerSystem.InitRouterDict(PrivateGroup)
		routerSystem.InitRouterApi(PrivateGroup)
		routerSystem.InitRouterUpload(PrivateGroup)
	}

	return Router
}
