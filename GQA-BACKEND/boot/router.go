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
	r.Use(middleware.Cors()).Use(middleware.I18nHandler())
	StaticFS(r)

	//Public route: starts with "public". There is no need to regroup or authenticate within the route
	PublicGroup := r.Group("public")
	RouterPublic(PublicGroup)

	//Private route：starts with "". The route is grouped according to the actual performance, and authentication is required
	PrivateGroup := r.Group("")
	PrivateGroup.Use(middleware.JwtHandler()).Use(middleware.RoleApiHandler())
	RouterPrivate(PrivateGroup)

	//Register Plugin Route
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
		routerPrivate.RouterGenPlugin.InitRouterGenPlugin(PrivateGroup)
		routerPrivate.RouterLog.InitRouterRouterLog(PrivateGroup)
		routerPrivate.RouterMenu.InitRouterRouterMenu(PrivateGroup)
		routerPrivate.RouterTodo.InitRouterRouterTodo(PrivateGroup)
		routerPrivate.RouterNotice.InitRouterRouterNotice(PrivateGroup)
		routerPrivate.RouterRole.InitRouterRouterRole(PrivateGroup)
		routerPrivate.RouterUpload.InitRouterRouterUpload(PrivateGroup)
		routerPrivate.RouterUser.InitRouterRouterUser(PrivateGroup)
		routerPrivate.RouterUserOnline.InitRouterRouterUserOnline(PrivateGroup)
		routerPrivate.RouterCron.InitRouterRouterCron(PrivateGroup)
	}
}

func RouterPlugin(PublicGroup, PrivateGroup *gin.RouterGroup) {
	gqaplugin.RegisterPluginRouter(PublicGroup, PrivateGroup)
}
