package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oapms/pkg/mus"
	"oapms/pkg/service"
	"strings"
)

// Ignored permissions
var ignoredPerms = map[string]bool{}

func PermCheck(c *gin.Context) {
	return
	route := strings.Split(c.Request.URL.RequestURI(), "?")[0]
	for _, p := range c.Params {
		route = strings.Replace(route, "/"+p.Value, "/:"+p.Key, 1)
	}

	// 'delete@/v1/users/:id'
	route = strings.ToLower(c.Request.Method) + "@" + route
	uid := fmt.Sprintf("%#v", AdminContextUser(c).Uid)
	if _, ok := ignoredPerms[route]; ok {
		c.Next()
		return
	}
	if !service.Pms.CheckPermission(uid, 1, route) {
		mus.Logger.Warn(fmt.Sprintf("No permission for %s", route))
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  route + ", no pms",
		})
		c.Abort()
		return
	}
	mus.Logger.Info(fmt.Sprintf("Pass permission check for %s", route))
	c.Next()
}
