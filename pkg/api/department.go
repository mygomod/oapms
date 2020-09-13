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

func DepartmentList(c *core.Context) {
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

	if v := c.Query("pid"); v != "" {
		query["pid"] = v
	}

	if v := c.Query("orderNum"); v != "" {
		query["orderNum"] = v
	}

	if v := c.Query("extendField"); v != "" {
		query["extendField"] = v
	}

	if v := c.Query("intro"); v != "" {
		query["intro"] = v
	}

	if v := c.Query("createdAt"); v != "" {
		query["createdAt"] = v
	}

	if v := c.Query("updatedAt"); v != "" {
		query["updatedAt"] = v
	}

	total, list := model.DepartmentListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func DepartmentInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.DepartmentInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func DepartmentCreate(c *core.Context) {
	req := &model.Department{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.DepartmentCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func DepartmentDelete(c *core.Context) {
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

	err = model.DepartmentDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func DepartmentUpdate(c *core.Context) {
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

	err = model.DepartmentUpdate(mus.Db, id, reqJson)
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
