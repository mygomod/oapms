// @BeeOverwrite NO
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitUser(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/user/info", api.UserInfo).WithCustomInfo("show", "/oa/user")
	core.RegisterUrl(r, "get", "/api/admin/user/list", api.UserList).WithCustomInfo("show", "/oa/user")
	core.RegisterUrl(r, "get", "/api/admin/user/getRole", api.UserGetRole).WithCustomInfo("show", "/oa/user")
	core.RegisterUrl(r, "post", "/api/admin/user/sendGoogleCode", api.UserSendGoogleCode).WithCustomInfo("edit", "/oa/user")
	core.RegisterUrl(r, "post", "/api/admin/user/setRole", api.UserSetRole).WithCustomInfo("edit", "/oa/user")
	core.RegisterUrl(r, "post", "/api/admin/user/create", api.UserCreate).WithCustomInfo("edit", "/oa/user")
	core.RegisterUrl(r, "post", "/api/admin/user/update", api.UserUpdate).WithCustomInfo("edit", "/oa/user")
	core.RegisterUrl(r, "post", "/api/admin/user/delete", api.UserDelete).WithCustomInfo("delete", "/oa/user")
}
