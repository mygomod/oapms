// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitPms(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/pms/info", api.PmsInfo)
	core.RegisterUrl(r, "get", "/api/admin/pms/list", api.PmsList)
	core.RegisterUrl(r, "post", "/api/admin/pms/create", api.PmsCreate)
	core.RegisterUrl(r, "post", "/api/admin/pms/update", api.PmsUpdate)
	core.RegisterUrl(r, "post", "/api/admin/pms/delete", api.PmsDelete)
}
