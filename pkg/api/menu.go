// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230345
package api

import (
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/trans"
)

func MenuList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}
	if v := c.Query("appId"); v != "" {
		query["app_id"] = v
	}

	if v := c.Query("pid"); v != "" {
		query["pid"] = v
	}

	if v := c.Query("name"); v != "" {
		query["name"] = v
	}

	if v := c.Query("url"); v != "" {
		query["url"] = v
	}

	if v := c.Query("pmsCode"); v != "" {
		query["pms_code"] = v
	}

	total, list := model.MenuListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func MenuTree(c *core.Context) {
	aidStr := c.Query("aid")
	aid := cast.ToInt(aidStr)
	if aid == 0 {
		aid = 1
	}

	list, err := model.MenuList(model.Conds{"app_id": aid})
	if err != nil {
		c.JSONErrTips("get menu tree err", err)
		return
	}

	c.JSONOK(model.Menus(list).ToTree())
}

func MenuInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.MenuInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func MenuCreate(c *core.Context) {
	req := &model.Menu{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.MenuCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func MenuDelete(c *core.Context) {
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

	err = model.MenuDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func MenuUpdate(c *core.Context) {
	req := model.Menu{}
	err := c.Bind(&req)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	if req.Id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err = model.MenuUpdate(mus.Db, req.Id, model.Ups{
		"path": req.Path,
		"icon": req.Icon,
		"name": req.Name,
	})
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
