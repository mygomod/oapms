// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230345
package api

import (
	"fmt"
	"github.com/satori/uuid"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"image/png"
	"net/url"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/service"
	"oapms/pkg/trans"
	"strconv"
	"strings"
)

func UserList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("uid"); v != "" {
		query["uid"] = v
	}

	if v := c.Query("nickname"); v != "" {
		query["nickname"] = v
	}

	if v := c.Query("email"); v != "" {
		query["email"] = v
	}

	if v := c.Query("avatar"); v != "" {
		query["avatar"] = v
	}

	if v := c.Query("password"); v != "" {
		query["password"] = v
	}

	if v := c.Query("state"); v != "" {
		query["state"] = v
	}

	if v := c.Query("gender"); v != "" {
		query["gender"] = v
	}

	if v := c.Query("birthday"); v != "" {
		query["birthday"] = v
	}

	if v := c.Query("ctime"); v != "" {
		query["ctime"] = v
	}

	if v := c.Query("utime"); v != "" {
		query["utime"] = v
	}

	if v := c.Query("lastLoginIp"); v != "" {
		query["lastLoginIp"] = v
	}

	if v := c.Query("lastLoginTime"); v != "" {
		query["lastLoginTime"] = v
	}

	total, list := model.UserListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func UserInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("uid"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.UserInfo(mus.Db, reqId)

	c.JSONOK(info)
}

type ReqUserCreate struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	State    int    `json:"state"`
}

func UserCreate(c *core.Context) {
	req := &ReqUserCreate{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := service.User.Create(req.Nickname, req.Password, req.Email, req.State, c.ClientIP())
	if err != nil {
		c.JSONErrTips("create user err", err)
		return
	}
	c.JSONOK()
}

func UserDelete(c *core.Context) {
	reqJson := make(map[string]interface{}, 0)
	err := c.Bind(&reqJson)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	id := cast.ToInt(reqJson["uid"])
	if id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err = model.UserDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func UserUpdate(c *core.Context) {
	reqJson := make(map[string]interface{}, 0)
	err := c.Bind(&reqJson)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	id := cast.ToInt(reqJson["uid"])
	if id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err = model.UserUpdate(mus.Db, id, model.Ups{
		"email":    reqJson["email"],
		"nickname": reqJson["nickname"],
	})
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}

type ReqUserSetRole struct {
	Uid     int   `json:"uid"`
	RoleIds []int `json:"roleIds"`
}

func UserSetRole(c *core.Context) {
	reqJson := ReqUserSetRole{}
	err := c.Bind(&reqJson)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	if reqJson.Uid == 0 || len(reqJson.RoleIds) == 0 {
		c.JSONErrTips("id or role id is error", nil)
		return
	}

	service.Pms.AssignRole(reqJson.Uid, reqJson.RoleIds)
	c.JSONOK()
}

func UserGetRole(c *core.Context) {
	reqId := cast.ToInt(c.Query("uid"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}
	c.JSONOK(service.Pms.GetAllRoleIds(strconv.Itoa(reqId)))
}

type ReqUserSendGoogleCode struct {
	Uid int `json:"uid"`
}

func UserSendGoogleCode(c *core.Context) {
	req := &ReqUserSendGoogleCode{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	if req.Uid == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.UserInfo(mus.Db, req.Uid)

	s := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	_, err := mus.Redis.Set(s, info.Uid, 3600)
	if err != nil {
		c.JSONErrTips("系统错误", err)
		return
	}

	html, err := service.Mailer.ParseTpl("googlecode.html", map[string]interface{}{
		"image": viper.GetString("cdnUrl") + "/api/v1/user/showGoogleCode?token=" + s,
	})
	if err != nil {
		c.JSONErrTips("html error", err)
		return
	}

	fmt.Println(html)

	err = service.Mailer.Send("Google验证器二维码", info.Email, html, "")
	if err != nil {
		c.JSONErrTips("邮件发送失败", err)
		return
	}

	c.JSONOK(info)
}

// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
func UserShowGoogleCode(c *core.Context) {
	token := c.Query("token")
	uid, err := mus.Redis.GetInt(token)
	if err != nil {
		c.JSONErrTips("get token error", err)
		return
	}

	userSecretQuery, err := service.User.GetSecret(uid)
	if err != nil {
		c.JSONErrTips("get user secret error", err)
		return
	}
	account := userSecretQuery.Nickname
	issuer := viper.GetString("oauth.mfaName")
	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		c.JSONErrTips("url parse error", err)
		return
	}

	URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)
	params := url.Values{}
	params.Add("secret", userSecretQuery.Secret)
	params.Add("issuer", issuer)
	URL.RawQuery = params.Encode()
	p, err := qrcode.New(URL.String(), qrcode.Medium)
	if err != nil {
		c.JSONErrTips("qrcode new error", err)
		return
	}

	img := p.Image(256)
	header := c.Writer.Header()
	header.Add("Content-Type", "image/jpeg")
	c.Status(200)
	err = png.Encode(c.Writer, img)
	if err != nil {
		c.JSONErrTips("encode image err", err)
		return
	}
}
