package install

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"oapms/pkg/api"
	"oapms/pkg/model"
	"oapms/pkg/mus"
)

func (m *mock) mockUser(url string) {
	m.PostForm(url, api.ReqUserCreate{
		Nickname: "admin",
		Password: "123456",
	})
}

func (m *mock) mockApp(url string) {
	m.PostForm(url, model.App{
		Name:  "用户权限系统",
		State: 1,
		Url:   viper.GetString("cdnUrl"),
	})
	m.PostForm(url, model.App{
		Name:  "数据中台",
		State: 1,
	})
}

func (m *mock) mockRole(url string) {
	m.PostForm(url, api.ReqRoleCreateOrUpdate{
		Name:  "Oapms管理员",
		AppId: 1,
		Menus: []api.ReqRoleCreateMenu{
			{
				MenuId:  1,
				Actions: nil,
			},
			{
				MenuId:  2,
				Actions: nil,
			},
			{
				MenuId:  3,
				Actions: nil,
			},
			{
				MenuId:  4,
				Actions: nil,
			},
			{
				MenuId:  5,
				Actions: nil,
			},
			{
				MenuId:  6,
				Actions: nil,
			},
			{
				MenuId:  7,
				Actions: nil,
			},
			{
				MenuId:  8,
				Actions: nil,
			},
			{
				MenuId:  9,
				Actions: nil,
			},
			{
				MenuId:  10,
				Actions: nil,
			},
		},
	})

	m.PostForm(url, model.Role{
		Name:  "Oapms菜单管理员",
		AppId: 1,
	})
}

func (m *mock) mockCasbin(url string) {
	file, err := ioutil.ReadFile("./data/mockdata/casbin_rule.json")
	if err != nil {
		panic(err)
	}
	info := make([]model.CasbinRule, 0)
	err = json.Unmarshal(file, &info)
	if err != nil {
		panic(err)
	}
	for _, value := range info {
		m.PostForm(url, value)
	}
}

type menuJson struct {
	AppId int             `json:"appId"`
	Menus model.MenuTrees `json:"menus"`
}

func (m *mock) mockMenu() {
	file, err := ioutil.ReadFile("./data/mockdata/menu_oapms.json")
	if err != nil {
		panic(err)
	}
	info := make([]menuJson, 0)
	err = json.Unmarshal(file, &info)
	if err != nil {
		panic(err)
	}

	for _, value := range info {
		appId := value.AppId
		walk(appId, 0, value.Menus)
	}
}

func walk(appId int, id int, mTree model.MenuTrees) []*model.Menu {
	arr := make([]*model.Menu, 0)
	for _, value := range mTree {
		value.AppId = appId
		value.Pid = id
		value.State = 1
		mus.Db.Create(value)

		for _, pmsExtend := range value.MenuPmsKey {
			extend := model.MenuPms{
				PmsCode: value.PmsCode,
				Key:     pmsExtend,
				AppId:   appId,
			}
			mus.Db.Create(&extend)
		}
		arr = append(arr, value)
		if value.Children != nil {
			childArr := walk(appId, value.Id, *value.Children)
			arr = append(arr, childArr...)
		}
	}
	return arr
}
