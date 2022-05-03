package boot

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	var r = gin.Default()
	r.Use(middleware.Cors())
	StaticFS(r)
	//公共路由分组：以 public 开头，路由内部无须再次分组，无须鉴权。
	PublicGroup := r.Group("public")
	RouterPublic(PublicGroup)

	//鉴权路由分组：这里以空分组，路由内部按实绩情况分组，需要鉴权。
	PrivateGroup := r.Group("")
	PrivateGroup.Use(middleware.JwtHandler()).Use(middleware.RoleApiHandler())
	RouterPrivate(PrivateGroup)
	RouterPlugin(PublicGroup, PrivateGroup)
	return r
}

func StaticFS(r *gin.Engine) {
	r.StaticFS("/"+global.GqaServeAvatar, http.Dir(global.GqaServeUpload+"/"+global.GqaServeAvatar))
	r.StaticFS("/"+global.GqaServeFile, http.Dir(global.GqaServeUpload+"/"+global.GqaServeFile))
	r.StaticFS("/"+global.GqaServeLogo, http.Dir(global.GqaServeUpload+"/"+global.GqaServeLogo))
	r.StaticFS("/"+global.GqaServeBanner, http.Dir(global.GqaServeUpload+"/"+global.GqaServeBanner))
}

func RouterPublic(PublicGroup *gin.RouterGroup) {
	routerPublic := router.GqaRouter.RouterPublic
	{
		routerPublic.RouterCaptcha.InitRouterCaptcha(PublicGroup)
		routerPublic.RouterConfigBackend.InitRouterConfigBackend(PublicGroup)
		routerPublic.RouterConfigFrontend.InitRouterConfigFrontend(PublicGroup)
		routerPublic.RouterDb.InitRouterDb(PublicGroup)
		routerPublic.RouterDict.InitRouterDict(PublicGroup)
		routerPublic.RouterLogin.InitRouterLogin(PublicGroup)
		routerPublic.RouterWebsocket.InitRouterWebSocket(PublicGroup)
	}
}

func RouterPrivate(PrivateGroup *gin.RouterGroup) {
	routerPrivate := router.GqaRouter.RouterPrivate
	{
		routerPrivate.RouterApi.InitRouterRouterApi(PrivateGroup)
		routerPrivate.RouterConfigBackend.InitRouterRouterConfigBackend(PrivateGroup)
		routerPrivate.RouterConfigFrontend.InitRouterRouterConfigFrontend(PrivateGroup)
		routerPrivate.RouterDept.InitRouterRouterDept(PrivateGroup)
		routerPrivate.RouterDict.InitRouterRouterDict(PrivateGroup)
		routerPrivate.RouterLog.InitRouterRouterLog(PrivateGroup)
		routerPrivate.RouterMenu.InitRouterRouterMenu(PrivateGroup)
		routerPrivate.RouterNoteTodo.InitRouterRouterNoteTodo(PrivateGroup)
		routerPrivate.RouterNotice.InitRouterRouterNotice(PrivateGroup)
		routerPrivate.RouterRole.InitRouterRouterRole(PrivateGroup)
		routerPrivate.RouterUpload.InitRouterRouterUpload(PrivateGroup)
		routerPrivate.RouterUser.InitRouterRouterUser(PrivateGroup)
	}
}

func RouterPlugin(PublicGroup, PrivateGroup *gin.RouterGroup) {
	gqaplugin.RegisterPluginRouter(PublicGroup, PrivateGroup)
}
