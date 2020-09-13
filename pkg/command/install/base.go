package install

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http/httptest"
	"oapms/pkg/api"
	"oapms/pkg/router/core"
)

func MockData() {
	router := gin.New()
	urlUserCreate := "/user/create"
	urlAppCreate := "/app/create"
	urlRoleCreate := "/role/create"
	urlCasbinRuleCreate := "/casbinRule/create"
	router.POST(urlUserCreate, core.Handle(api.UserCreate))
	router.POST(urlAppCreate, core.Handle(api.AppCreate))
	router.POST(urlRoleCreate, core.Handle(api.RoleCreate))
	router.POST(urlCasbinRuleCreate, core.Handle(api.CasbinRuleCreate))
	server := &mock{
		router: router,
	}
	server.mockUser(urlUserCreate)
	server.mockApp(urlAppCreate)
	server.mockMenu()
	server.mockRole(urlRoleCreate)
	server.mockCasbin(urlCasbinRuleCreate)
}

type mock struct {
	router *gin.Engine
	db     *gorm.DB
}

func (m *mock) PostForm(uri string, param interface{}) []byte {
	postByte, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	// 构造post请求
	req := httptest.NewRequest("POST", uri, bytes.NewReader(postByte))
	req.Header.Set("Content-Type", "application/json")

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	m.router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

func (m *mock) Get(uri string) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	m.router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
