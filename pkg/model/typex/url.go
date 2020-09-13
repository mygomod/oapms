package typex

import (
	"fmt"
	"sync"
)

var storeUrlMap sync.Map
var codeNameMap = map[string]string{
	"edit":   "编辑",
	"delete": "删除",
	"show":   "查看",
}

type Url struct {
	Name   string   `json:"name"`
	Code   string   `json:"code"`
	Menus  []string `json:"menus"`
	Method string   `json:"method"`
	Path   string   `json:"path"`
	Key    string   `json:"key"`
}

func (u Url) String() string {
	return u.Key
}

// common
func (u Url) WithInfo(name string, code string, menus ...string) {
	u.Name = name
	u.Code = code
	u.Menus = menus
	u.Store()
}

func (u Url) WithCustomInfo(code string, menu string) {
	var flag bool
	u.Name, flag = codeNameMap[code]
	if !flag {
		panic("error custom info, code is " + code)
	}
	u.Code = fmt.Sprintf("%s:%s", menu, code)
	u.Menus = []string{menu}
	u.Store()
}

func (u Url) Store() {
	storeUrlMap.Store(u.String(), u)
}

func OutputUrl() []Url {
	urls := make([]Url, 0)
	storeUrlMap.Range(func(key, value interface{}) bool {
		u, ok := value.(Url)
		if !ok {
			return true
		}
		urls = append(urls, u)
		return true
	})
	return urls
}
