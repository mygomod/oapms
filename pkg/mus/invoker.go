package mus

import (
	dingtalk "github.com/bullteam/go-dingtalk/src"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mygomod/muses/pkg/cache/redis"
	mmysql "github.com/mygomod/muses/pkg/database/mysql"
	"github.com/mygomod/muses/pkg/logger"
	musgin "github.com/mygomod/muses/pkg/server/gin"
	"github.com/mygomod/muses/pkg/session/ginsession"
	"github.com/spf13/viper"
)

var (
	Cfg            musgin.Cfg
	Logger         *logger.Client
	Gin            *gin.Engine
	Db             *gorm.DB
	Session        gin.HandlerFunc
	Redis          *redis.Client
	DingTalkClient *dingtalk.DingTalkClient
)

// Init 初始化muses相关容器
func Init() error {
	Cfg = musgin.Config()
	Db = mmysql.Caller("oapms")
	if Db == nil {
		panic("db nil")
	}
	Redis = redis.Caller("oapms")
	if Redis == nil {
		panic("redis nil")
	}
	Logger = logger.Caller("system")
	Gin = musgin.Caller()
	Session = ginsession.Caller()
	if Session == nil {
		panic("session nil")
	}
	DingTalkClient = dingtalk.NewDingTalkCompanyClient(&dingtalk.DTConfig{
		AppKey:    viper.GetString("dingtalk.appKey"),
		AppSecret: viper.GetString("dingtalk.appSecret"),
		CachePath: viper.GetString("dingtalk.cachePath"),
	})
	return nil

}
