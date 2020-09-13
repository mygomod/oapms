package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/model/typex"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/service"
)

func InitRouter() *gin.Engine {
	r := mus.Gin
	adminGrp := r.Group("", mus.Session)

	adminGrp.POST("/api/admin/oauth/login", core.Handle(api.OauthLogin))

	adminGrp.POST("/api/v1/oauth/token", core.Handle(api.OauthToken))
	adminGrp.GET("/api/v1/oauth/info", core.Handle(api.OauthInfo))
	adminGrp.GET("/api/v1/oauth/user", core.Handle(api.OauthUser))
	adminGrp.GET("/api/v1/user/showGoogleCode", core.Handle(api.UserShowGoogleCode))
	adminGrp.GET("/api/admin/account/googleCode", core.Handle(api.AccountGoogleCode))
	adminGrp.GET("/api/admin/output/url", func(context *gin.Context) {
		context.JSON(200, typex.OutputUrl())
	})

	// must login
	adminGrp.Use(core.AdminLoginRequired())
	adminGrp.GET("/api/admin/oauth/redirect", core.Handle(api.OauthRedirect))
	adminGrp.GET("/api/admin/oauth/logout", core.Handle(api.OauthLogout))
	adminGrp.GET("/api/admin/oauth/user", core.Handle(api.OauthAdminUser))
	adminGrp.GET("/api/admin/account/app", core.Handle(api.AccountApp))
	adminGrp.GET("/api/admin/account/menu", core.Handle(api.AccountMenu))

	adminGrp.Use(core.PermCheck)

	InitUser(adminGrp)
	InitApp(adminGrp)
	InitMenu(adminGrp)
	InitRole(adminGrp)

	service.PmsMenu.AutoMenu(1, typex.OutputUrl())

	return r
}
