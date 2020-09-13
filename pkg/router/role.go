// @BeeOverwrite NO
// @BeeGenerateTime 20200831_101837
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitRole(r gin.IRoutes) {
	r.GET("/api/admin/role/menus", core.Handle(api.RoleMenus))
	r.GET("/api/admin/role/info", core.Handle(api.RoleInfo))
	core.RegisterUrl(r, "get", "/api/admin/role/all", api.RoleAll).WithCustomInfo("show", "/oa/user")
	r.GET("/api/admin/role/list", core.Handle(api.RoleList))
	r.POST("/api/admin/role/create", core.Handle(api.RoleCreate))
	r.POST("/api/admin/role/update", core.Handle(api.RoleUpdate))
	r.POST("/api/admin/role/delete", core.Handle(api.RoleDelete))
}
