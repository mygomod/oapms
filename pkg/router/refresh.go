// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitRefresh(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/refresh/info", api.RefreshInfo)
	core.RegisterUrl(r, "get", "/api/admin/refresh/list", api.RefreshList)
	core.RegisterUrl(r, "post", "/api/admin/refresh/create", api.RefreshCreate)
	core.RegisterUrl(r, "post", "/api/admin/refresh/update", api.RefreshUpdate)
	core.RegisterUrl(r, "post", "/api/admin/refresh/delete", api.RefreshDelete)
}
