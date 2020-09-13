// @BeeOverwrite NO
// @BeeGenerateTime 20200820_195417
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/service"
	"oapms/pkg/trans"
)

func RoleAll(c *core.Context) {
	aid := cast.ToInt(c.Query("aid"))
	list, _ := model.RoleList(model.Conds{
		"app_id": aid,
	})
	c.JSONOK(list)
}

func RoleList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}
	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("name"); v != "" {
		query["name"] = v
	}

	if v := c.Query("appId"); v != "" {
		query["app_id"] = v
	}

	total, list := model.RoleListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func RoleInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.RoleInfo(mus.Db, reqId)

	c.JSONOK(info)
}

type ReqRoleCreateOrUpdate struct {
	Id      int                 `json:"id"`
	Name    string              `json:"name"`
	AppId   int                 `json:"appId"`
	Intro   string              `json:"intro"`
	Menus   []ReqRoleCreateMenu `json:"menus"`
	DataPms []int               `json:"dataPms"`
}

type ReqRoleCreateMenu struct {
	MenuId  int   `json:"menuId"`
	Actions []int `json:"actions"`
}

func RoleCreate(c *core.Context) {
	req := &ReqRoleCreateOrUpdate{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	menuIds := make(model.MenuIdsJson, 0)
	for _, value := range req.Menus {
		menuIds = append(menuIds, value.MenuId)
		if len(value.Actions) > 0 {
			for _, action := range value.Actions {
				menuIds = append(menuIds, action)
			}
		}
	}

	roleData := &model.Role{
		Name:    req.Name,
		AppId:   req.AppId,
		Intro:   req.Intro,
		MenuIds: menuIds,
	}

	err := model.RoleCreate(mus.Db, roleData)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}

	dataPermCount := len(req.DataPms)
	if len(menuIds) > 0 {
		err = service.Pms.AssignMenuPms(roleData.Id, menuIds, dataPermCount)
		if err != nil {
			c.JSONErrTips("授权菜单、接口权限失败", err)
			return
		}
	}

	// insert data permissions
	if dataPermCount > 0 {
		err = service.Pms.AssignDataPms(roleData.Id, menuIds)
		if err != nil {
			c.JSONErrTips("授权数据权限失败", err)
			return
		}
	}

	c.JSONOK(req)
}

func RoleDelete(c *core.Context) {
	reqJson := make(map[string]interface{}, 0)
	err := c.Bind(&reqJson)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	id := cast.ToInt(reqJson["id"])
	if id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err = model.RoleDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func RoleUpdate(c *core.Context) {
	req := &ReqRoleCreateOrUpdate{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	menuIds := make(model.MenuIdsJson, 0)
	for _, value := range req.Menus {
		menuIds = append(menuIds, value.MenuId)
		if len(value.Actions) > 0 {
			for _, action := range value.Actions {
				menuIds = append(menuIds, action)
			}
		}
	}

	if req.Id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err := model.RoleUpdate(mus.Db, req.Id, gin.H{
		"name":     req.Name,
		"app_id":   req.AppId,
		"intro":    req.Intro,
		"menu_ids": menuIds,
	})
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}

	err = service.Pms.AssignMenuPms(req.Id, menuIds, len(req.DataPms))
	if err != nil {
		c.JSONErrTips("授权菜单、接口权限失败", err)
		return
	}

	err = service.Pms.AssignDataPms(req.Id, req.DataPms)
	if err != nil {
		c.JSONErrTips("授权数据权限失败", err)
		return
	}
	c.JSONOK()
}

func RoleMenus(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}
	info, _ := model.RoleInfo(mus.Db, reqId)
	menus := model.MenuAllListByIds(info.MenuIds)
	c.JSONOK(menus.ToCheckArr())
}
