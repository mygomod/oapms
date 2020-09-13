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

func RolePmsList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("roleId"); v != "" {
		query["roleId"] = v
	}

	if v := c.Query("pmsId"); v != "" {
		query["pmsId"] = v
	}

	total, list := model.RolePmsListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func RolePmsInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.RolePmsInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func RolePmsCreate(c *core.Context) {
	req := &model.RolePms{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.RolePmsCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func RolePmsDelete(c *core.Context) {
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

	err = model.RolePmsDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func RolePmsUpdate(c *core.Context) {
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

	err = model.RolePmsUpdate(mus.Db, id, reqJson)
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
