// @BeeOverwrite NO
// @BeeGenerateTime 20200822_105621
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitMenu(r gin.IRoutes) {
	r.GET("/api/admin/menu/info", core.Handle(api.MenuInfo))
	r.GET("/api/admin/menu/list", core.Handle(api.MenuList))
	r.GET("/api/admin/menu/tree", core.Handle(api.MenuTree))
	r.POST("/api/admin/menu/create", core.Handle(api.MenuCreate))
	r.POST("/api/admin/menu/update", core.Handle(api.MenuUpdate))
	r.POST("/api/admin/menu/delete", core.Handle(api.MenuDelete))
}
