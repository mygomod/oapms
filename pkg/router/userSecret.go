// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitUserSecret(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/userSecret/info", api.UserSecretInfo)
	core.RegisterUrl(r, "get", "/api/admin/userSecret/list", api.UserSecretList)
	core.RegisterUrl(r, "post", "/api/admin/userSecret/create", api.UserSecretCreate)
	core.RegisterUrl(r, "post", "/api/admin/userSecret/update", api.UserSecretUpdate)
	core.RegisterUrl(r, "post", "/api/admin/userSecret/delete", api.UserSecretDelete)
}
