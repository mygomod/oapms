package api

import (
	"encoding/json"
	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/router/core"
	"oapms/pkg/service"
	"strconv"
	"strings"
)

// {status: "error", type: "account", currentAuthority: "guest"}
func OauthLogin(c *core.Context) {
	// 如果已经登录
	respView := model.RespOauthLogin{
		CurrentAuthority: "admin",
	}

	// 如果已经登录
	reqView := model.ReqOauthLogin{}

	err := c.Bind(&reqView)
	if err != nil {
		mus.Logger.Error("json marshal request body error", zap.Error(err))
		c.JSONErrTips("login params error", err)
		return
	}

	if c.IsAuthenticated() && reqView.Params.RedirectUri != "" {
		user, ok := core.AdminSessionUser(c.Context)
		if !ok {
			c.JSONErrTips("用户不存在", err)
			return
		}
		oauthServer(c, reqView, user)
		return
	}

	if c.IsAuthenticated() && reqView.Params.RedirectUri == "" {
		c.JSONOK(respView)
		return
	}

	userInfo, err := service.User.GetUserByNicknamePwd(reqView.Nickname, reqView.Password, c.ClientIP())
	if err != nil {
		mus.Logger.Error("pwd error", zap.Error(err))
		c.JSONErrTips("密码错误", err)
		return
	}

	if reqView.Params.RedirectUri == "" {
		err = c.LoginByUid(userInfo.Uid, reqView.Mfa)
		if err != nil {
			c.JSONErrTips(err.Error(), err)
			return

		}
		c.JSONOK(respView)
		return
	}
	//c.String(200,"fuck")

	// redirect url exist
	oauthServer(c, reqView, &userInfo)

	return
}

func oauthServer(c *core.Context, reqView model.ReqOauthLogin, u *model.User) {
	server := service.Oauth2Server.GetServer()
	resp := server.NewResponse()
	r := c.Request

	r.Form = make(url.Values, 0)
	r.Form.Set("client_id", reqView.Params.ClientId)
	r.Form.Set("redirect_uri", reqView.Params.RedirectUri)
	r.Form.Set("response_type", reqView.Params.ResponseType)

	defer resp.Close()

	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
		ar.Authorized = true
		out, _ := json.Marshal(u)
		ar.UserData = string(out)
		server.FinishAuthorizeRequest(resp, r, ar)
	}

	if resp.IsError && resp.InternalError != nil {
		mus.Logger.Error("oauth2 error", zap.Error(resp.InternalError))
		c.JSONErrTips(resp.InternalError.Error(), resp.InternalError)
		return
	}

	err := c.LoginByUid(u.Uid, reqView.Mfa)
	if err != nil {
		c.JSONErrTips("login err", err)
		return

	}

	if resp.Type != osin.REDIRECT {
		c.JSONErrTips("login type err", nil)
		return
	}

	// Output redirect with parameters
	redirectUri, err := resp.GetRedirectUrl()
	if err != nil {
		mus.Logger.Error("get redirect url error", zap.Error(err))
		c.JSONErrTips("get redirect url err", err)
		return
	}

	_, err = mus.Redis.Set("uid_redirect_"+strconv.Itoa(u.Uid), redirectUri, 30)
	if err != nil {
		c.JSONErrTips("系统错误", err)
		return
	}

	c.JSONResult(301, "redirect", gin.H{
		"redirect_uri": "/api/admin/oauth/redirect",
	})
	return
}

func OauthRedirect(c *core.Context) {
	rUrl, err := mus.Redis.GetString("uid_redirect_" + strconv.Itoa(c.AdminUid()))
	if err != nil {
		c.JSONErrTips("系统错误", err)
		return
	}
	c.Redirect(302, rUrl)
	return
}

func OauthToken(c *core.Context) {
	server := service.Oauth2Server.GetServer()
	resp := server.NewResponse()
	defer resp.Close()
	r := c.Request
	w := c.Writer

	if ar := server.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		server.FinishAccessRequest(resp, r, ar)
	}

	osin.OutputJSON(resp, w, r)
	return
}

func OauthInfo(c *core.Context) {
	server := service.Oauth2Server.GetServer()
	resp := server.NewResponse()
	r := c.Request
	w := c.Writer
	defer resp.Close()

	if ir := server.HandleInfoRequest(resp, r); ir != nil {
		server.FinishInfoRequest(resp, r, ir)
	}
	osin.OutputJSON(resp, w, r)
	return
}

func OauthUser(c *core.Context) {
	//  Authorization:[Askuy qYBgP93GSkqvQCZTEGzODw]
	auth := c.Request.Header.Get("Authorization")
	auths := strings.Split(auth, " ")
	if len(auths) < 2 {
		c.JSONErrTips("parse auth error", nil)
		return
	}

	info, err := model.AccessInfoX(mus.Db, model.Conds{"access_token": auths[1]})
	if err != nil {
		c.JSONErrTips("get auth error", nil)
		return
	}
	c.JSONResultRaw(0, "ok", []byte(info.Extra))

	return
}

func OauthAdminUser(c *core.Context) {
	if !c.IsAuthenticated() {
		c.JSONErrTips("no auth", nil)
		return
	}

	info, err := model.UserInfo(mus.Db, c.AdminUid())
	if err != nil {
		c.JSONErrTips("no user", err)
		return
	}
	c.JSONOK(info)
	return
}

func OauthLogout(c *core.Context) {
	err := c.Logout()
	if err != nil {
		c.JSONErrTips("logout err", err)
		return
	}
	c.JSONOK()
	return
}
