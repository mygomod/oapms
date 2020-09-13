// @BeeOverwrite YES
// @BeeGenerateTime 20200902_214035
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitMenuPms(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/menuPms/info", api.MenuPmsInfo)
	core.RegisterUrl(r, "get", "/api/admin/menuPms/list", api.MenuPmsList)
	core.RegisterUrl(r, "post", "/api/admin/menuPms/create", api.MenuPmsCreate)
	core.RegisterUrl(r, "post", "/api/admin/menuPms/update", api.MenuPmsUpdate)
	core.RegisterUrl(r, "post", "/api/admin/menuPms/delete", api.MenuPmsDelete)
}
