// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitAccess(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/access/info", api.AccessInfo)
	core.RegisterUrl(r, "get", "/api/admin/access/list", api.AccessList)
	core.RegisterUrl(r, "post", "/api/admin/access/create", api.AccessCreate)
	core.RegisterUrl(r, "post", "/api/admin/access/update", api.AccessUpdate)
	core.RegisterUrl(r, "post", "/api/admin/access/delete", api.AccessDelete)
}
