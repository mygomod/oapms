// @BeeOverwrite YES
// @BeeGenerateTime 20200820_195417
package api

import (
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/trans"
)

func CasbinRuleList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("p"); v != "" {
		query["p"] = v
	}

	if v := c.Query("v0"); v != "" {
		query["v0"] = v
	}

	if v := c.Query("v1"); v != "" {
		query["v1"] = v
	}

	if v := c.Query("v2"); v != "" {
		query["v2"] = v
	}

	if v := c.Query("v3"); v != "" {
		query["v3"] = v
	}

	if v := c.Query("v4"); v != "" {
		query["v4"] = v
	}

	if v := c.Query("v5"); v != "" {
		query["v5"] = v
	}

	total, list := model.CasbinRuleListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func CasbinRuleInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.CasbinRuleInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func CasbinRuleCreate(c *core.Context) {
	req := &model.CasbinRule{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.CasbinRuleCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func CasbinRuleDelete(c *core.Context) {
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

	err = model.CasbinRuleDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func CasbinRuleUpdate(c *core.Context) {
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

	err = model.CasbinRuleUpdate(mus.Db, id, reqJson)
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
