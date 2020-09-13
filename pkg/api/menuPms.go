// @BeeOverwrite YES
// @BeeGenerateTime 20200902_214035
package api

import (
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/trans"
)

func MenuPmsList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("pmsCode"); v != "" {
		query["pmsCode"] = v
	}

	if v := c.Query("key"); v != "" {
		query["key"] = v
	}

	if v := c.Query("appId"); v != "" {
		query["appId"] = v
	}

	if v := c.Query("ctime"); v != "" {
		query["ctime"] = v
	}

	if v := c.Query("utime"); v != "" {
		query["utime"] = v
	}

	total, list := model.MenuPmsListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func MenuPmsInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.MenuPmsInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func MenuPmsCreate(c *core.Context) {
	req := &model.MenuPms{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.MenuPmsCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func MenuPmsDelete(c *core.Context) {
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

	err = model.MenuPmsDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func MenuPmsUpdate(c *core.Context) {
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

	err = model.MenuPmsUpdate(mus.Db, id, reqJson)
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
