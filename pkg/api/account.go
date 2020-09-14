package api

import (
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"image/png"
	"net/url"
	"oapms/pkg/router/core"
	"oapms/pkg/service"
	"strconv"
)

func AccountAllPms(c *core.Context) {
	c.JSONOK(service.Pms.GetAllPermissions(strconv.Itoa(c.AdminUid())))
}

func AccountApp(c *core.Context) {
	c.JSONOK(service.Pms.GetRelatedApp(strconv.Itoa(c.AdminUid()), true))
}

func AccountMenu(c *core.Context) {
	c.JSONOK(service.Pms.GetAppMenu(strconv.Itoa(c.AdminUid()), strconv.Itoa(1)).ToTree())
}

// https://github.com/google/google-authenticator/wiki/Key-Uri-Format
func AccountGoogleCode(c *core.Context) {
	userSecretQuery, err := service.User.GetSecret(1)
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

	//URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)
	URL.Path += "/" + url.PathEscape(issuer) + ":" + account
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
