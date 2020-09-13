// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitExpires(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/expires/info", api.ExpiresInfo)
	core.RegisterUrl(r, "get", "/api/admin/expires/list", api.ExpiresList)
	core.RegisterUrl(r, "post", "/api/admin/expires/create", api.ExpiresCreate)
	core.RegisterUrl(r, "post", "/api/admin/expires/update", api.ExpiresUpdate)
	core.RegisterUrl(r, "post", "/api/admin/expires/delete", api.ExpiresDelete)
}
