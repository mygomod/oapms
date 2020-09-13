// @BeeOverwrite NO
// @BeeGenerateTime 20200822_105621
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitApp(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/app/info", api.AppInfo).WithCustomInfo("show", "/oa/app")
	core.RegisterUrl(r, "get", "/api/admin/app/select", api.AppSelect).WithCustomInfo("show", "/oa/app")
	core.RegisterUrl(r, "get", "/api/admin/app/selectArr", api.AppSelectArr).WithCustomInfo("show", "/oa/app")
	core.RegisterUrl(r, "get", "/api/admin/app/list", api.AppList).WithCustomInfo("show", "/oa/app")
	core.RegisterUrl(r, "post", "/api/admin/app/create", api.AppCreate).WithCustomInfo("edit", "/oa/app")
	core.RegisterUrl(r, "post", "/api/admin/app/update", api.AppUpdate).WithCustomInfo("edit", "/oa/app")
	core.RegisterUrl(r, "post", "/api/admin/app/delete", api.AppDelete).WithCustomInfo("delete", "/oa/app")
}
