package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oapms/pkg/model/typex"
)

func RegisterUrl(r gin.IRoutes, method string, path string, h HandlerFunc) typex.Url {
	u := typex.Url{
		Method: method,
		Path:   path,
		Key:    fmt.Sprintf("%s@%s", method, path),
	}

	switch method {
	case "get":
		u.Store()
		r.GET(path, Handle(h))
	case "post":
		u.Store()
		r.POST(path, Handle(h))
	case "put":
		u.Store()
		r.PUT(path, Handle(h))
	case "delete":
		u.Store()
		r.DELETE(path, Handle(h))
	default:
		panic("not support type, type:" + method)
	}
	return u
}
