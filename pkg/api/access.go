// @BeeOverwrite YES
// @BeeGenerateTime 20200820_230345
package api

import (
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/trans"
)

func AccessList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("client"); v != "" {
		query["client"] = v
	}

	if v := c.Query("authorize"); v != "" {
		query["authorize"] = v
	}

	if v := c.Query("previous"); v != "" {
		query["previous"] = v
	}

	if v := c.Query("accessToken"); v != "" {
		query["accessToken"] = v
	}

	if v := c.Query("refreshToken"); v != "" {
		query["refreshToken"] = v
	}

	if v := c.Query("expiresIn"); v != "" {
		query["expiresIn"] = v
	}

	if v := c.Query("scope"); v != "" {
		query["scope"] = v
	}

	if v := c.Query("redirectUri"); v != "" {
		query["redirectUri"] = v
	}

	if v := c.Query("extra"); v != "" {
		query["extra"] = v
	}

	if v := c.Query("ctime"); v != "" {
		query["ctime"] = v
	}

	total, list := model.AccessListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func AccessInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.AccessInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func AccessCreate(c *core.Context) {
	req := &model.Access{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.AccessCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func AccessDelete(c *core.Context) {
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

	err = model.AccessDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func AccessUpdate(c *core.Context) {
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

	err = model.AccessUpdate(mus.Db, id, reqJson)
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
