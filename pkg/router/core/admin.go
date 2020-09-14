package core

import (
	"fmt"
	"github.com/dgryski/dgoogauth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"oapms/pkg/service"
	"time"
)

// 后台取用户
func AdminLoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从session中获取用户信息
		user, ok := AdminSessionUser(c)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 302,
				"data": RedirectLoginURL,
				"msg":  "user not login",
			})
			c.Abort()
			return
		}

		if user.Uid == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 302,
				"data": RedirectLoginURL,
				"msg":  "user not login",
			})
			c.Abort()
			return
		}

		c.Set(AdminContextKey, user)
		c.Next()
	}
}

func AdminLoginRedirectRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从session中获取用户信息
		user, ok := AdminSessionUser(c)
		if !ok {
			c.Redirect(http.StatusFound, RedirectLoginURL)
			c.Abort()
			return
		}

		if user.Uid == 0 {
			c.Redirect(http.StatusFound, RedirectLoginURL)
			c.Abort()
			return
		}

		c.Set(AdminContextKey, user)
		c.Next()
	}
}

// 后台取用户
func AdminSessionUser(c *gin.Context) (*model.User, bool) {
	resp, flag := sessions.Default(c).Get(AdminSessionKey).(*model.User)
	return resp, flag
}

// 后台取用户
func AdminContextUser(c *gin.Context) *model.User {
	resp := &model.User{}
	respI, flag := c.Get(AdminContextKey)
	if flag {
		resp = respI.(*model.User)
	}
	return resp
}

// Authed 鉴权通过
func (c *Context) IsAuthenticated() bool {
	if user, ok := AdminSessionUser(c.Context); ok && user.Uid > 0 {
		return true
	}
	return false
}

func (c *Context) AuthDirect() {
	c.JSON(http.StatusOK, gin.H{
		"code": 302,
		"data": RedirectLoginURL,
		"msg":  "user not login",
	})
}

func (c *Context) AdminUser() *model.User {
	return AdminContextUser(c.Context)
}

// 后台 Uid 返回uid
func (c *Context) AdminUid() int {
	return AdminContextUser(c.Context).Uid
}

func (c *Context) AdminName() string {
	return AdminContextUser(c.Context).Nickname
}

// UpdateSession updates the User object stored in the session. This is useful incase a change
// is made to the user model that needs to persist across requests.
func (c *Context) LoginByUid(uid int, mfa string) error {
	user, err := model.UserInfo(mus.Db, uid)
	if err != nil {
		return err
	}

	fmt.Println("aaaaa", viper.GetBool("oauth.mfa"))
	// google 验证器
	if viper.GetBool("oauth.mfa") {
		if mfa == "" {
			return errors.New("用户mfa为空")
		}

		userSecret, err := service.User.GetSecret(uid)
		if err != nil {
			return err
		}

		otpc := &dgoogauth.OTPConfig{
			Secret:      userSecret.Secret,
			WindowSize:  3,
			HotpCounter: 0,
		}
		val, err := otpc.Authenticate(mfa)
		if err != nil {
			return err
		}
		if !val {
			return errors.New("not authenticate")
		}
	}

	err = model.UserUpdate(mus.Db, uid, model.Ups{
		"last_login_time": time.Now().Unix(),
		"last_login_ip":   c.ClientIP(),
	})
	if err != nil {
		return err
	}

	s := sessions.Default(c.Context)
	s.Options(sessions.Options{
		Path:     "/",
		MaxAge:   24 * 3600,
		Secure:   false,
		HttpOnly: true,
	})
	s.Set(AdminSessionKey, &user)
	return s.Save()
}

// Logout will clear out the session and call the Logout() user function.
func (c *Context) Logout() error {
	s := sessions.Default(c.Context)
	s.Options(sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	})

	s.Delete(AdminSessionKey)
	return s.Save()
}
