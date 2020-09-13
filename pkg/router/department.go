// @BeeOverwrite YES
// @BeeGenerateTime 20200902_160324
package router

import (
	"github.com/gin-gonic/gin"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func InitDepartment(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/admin/department/info", api.DepartmentInfo)
	core.RegisterUrl(r, "get", "/api/admin/department/list", api.DepartmentList)
	core.RegisterUrl(r, "post", "/api/admin/department/create", api.DepartmentCreate)
	core.RegisterUrl(r, "post", "/api/admin/department/update", api.DepartmentUpdate)
	core.RegisterUrl(r, "post", "/api/admin/department/delete", api.DepartmentDelete)
}
