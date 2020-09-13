// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230345
package api

import (
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/service"
	"oapms/pkg/trans"
)

func PmsList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("appId"); v != "" {
		query["appId"] = v
	}

	if v := c.Query("pid"); v != "" {
		query["pid"] = v
	}

	if v := c.Query("name"); v != "" {
		query["name"] = v
	}

	if v := c.Query("pmsCode"); v != "" {
		query["pmsCode"] = v
	}

	if v := c.Query("pmsRule"); v != "" {
		query["pmsRule"] = v
	}

	if v := c.Query("pmsType"); v != "" {
		query["pmsType"] = v
	}

	if v := c.Query("orderNum"); v != "" {
		query["orderNum"] = v
	}

	if v := c.Query("intro"); v != "" {
		query["intro"] = v
	}

	if v := c.Query("ctime"); v != "" {
		query["ctime"] = v
	}

	if v := c.Query("utime"); v != "" {
		query["utime"] = v
	}

	total, list := model.PmsListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func PmsInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.PmsInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func PmsCreate(c *core.Context) {
	req := &model.Pms{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.PmsCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func PmsDelete(c *core.Context) {
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

	err = model.PmsDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func PmsUpdate(c *core.Context) {
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

	err = model.PmsUpdate(mus.Db, id, reqJson)
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}

func PmsRoles(c *core.Context) {
	uid := c.Query("uid")
	c.JSONOK(service.Pms.GetAllRoles(uid))
}
