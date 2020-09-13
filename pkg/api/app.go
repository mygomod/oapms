// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230055
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/random"
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/trans"
)

func AppList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("aid"); v != "" {
		query["aid"] = v
	}

	if v := c.Query("clientId"); v != "" {
		query["clientId"] = v
	}

	if v := c.Query("name"); v != "" {
		query["name"] = v
	}

	total, list := model.AppListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

type appSelect struct {
	Text string `json:"text"`
}

func AppSelect(c *core.Context) {
	list, err := model.AppList(model.Conds{})
	if err != nil {
		c.JSONErrTips("get app select list err", err)
		return
	}

	output := make(map[int]appSelect)
	for _, value := range list {
		output[value.Aid] = appSelect{
			Text: value.Name,
		}
	}
	c.JSONOK(output)
}

type appSelectArr struct {
	Key   int    `json:"key"`
	Title string `json:"title"`
}

func AppSelectArr(c *core.Context) {
	list, err := model.AppList(model.Conds{})
	if err != nil {
		c.JSONErrTips("get app select list err", err)
		return
	}

	output := make([]appSelectArr, 0)

	for _, value := range list {
		output = append(output, appSelectArr{
			Key:   value.Aid,
			Title: value.Name,
		})
	}
	c.JSONOK(output)
}

func AppInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("aid"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.AppInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func AppCreate(c *core.Context) {
	req := &model.App{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}
	// todo check params
	appCreate := &model.App{
		ClientId:    random.String(16, random.Alphanumeric),
		Name:        req.Name,
		Secret:      random.String(32, random.Alphanumeric),
		RedirectUri: req.RedirectUri,
		Extra:       req.Extra,
		State:       req.State,
		Url:         req.Url,
	}

	err := model.AppCreate(mus.Db, appCreate)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func AppDelete(c *core.Context) {
	reqJson := make(map[string]interface{}, 0)
	err := c.Bind(&reqJson)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	id := cast.ToInt(reqJson["aid"])
	if id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err = model.AppDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func AppUpdate(c *core.Context) {
	req := &model.App{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	if req.Aid == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err := model.AppUpdate(mus.Db, req.Aid, gin.H{
		"name":         req.Name,
		"extra":        req.Extra,
		"state":        req.State,
		"redirect_uri": req.RedirectUri,
	})
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
