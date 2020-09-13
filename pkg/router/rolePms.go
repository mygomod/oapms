// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitRolePms(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/rolePms/info", api.RolePmsInfo)
	core.RegisterUrl(r, "get", "/api/admin/rolePms/list", api.RolePmsList)
	core.RegisterUrl(r, "post", "/api/admin/rolePms/create", api.RolePmsCreate)
	core.RegisterUrl(r, "post", "/api/admin/rolePms/update", api.RolePmsUpdate)
	core.RegisterUrl(r, "post", "/api/admin/rolePms/delete", api.RolePmsDelete)
}
