// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitCasbinRule(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/casbinRule/info", api.CasbinRuleInfo)
	core.RegisterUrl(r, "get", "/api/admin/casbinRule/list", api.CasbinRuleList)
	core.RegisterUrl(r, "post", "/api/admin/casbinRule/create", api.CasbinRuleCreate)
	core.RegisterUrl(r, "post", "/api/admin/casbinRule/update", api.CasbinRuleUpdate)
	core.RegisterUrl(r, "post", "/api/admin/casbinRule/delete", api.CasbinRuleDelete)
}
