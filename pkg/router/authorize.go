// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitAuthorize(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/authorize/info", api.AuthorizeInfo)
	core.RegisterUrl(r, "get", "/api/admin/authorize/list", api.AuthorizeList)
	core.RegisterUrl(r, "post", "/api/admin/authorize/create", api.AuthorizeCreate)
	core.RegisterUrl(r, "post", "/api/admin/authorize/update", api.AuthorizeUpdate)
	core.RegisterUrl(r, "post", "/api/admin/authorize/delete", api.AuthorizeDelete)
}
