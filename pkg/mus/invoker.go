package mus

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mygomod/muses/pkg/cache/redis"
	mmysql "github.com/mygomod/muses/pkg/database/mysql"
	"github.com/mygomod/muses/pkg/logger"
	musgin "github.com/mygomod/muses/pkg/server/gin"
	"github.com/mygomod/muses/pkg/session/ginsession"
)

var (
	Cfg     musgin.Cfg
	Logger  *logger.Client
	Gin     *gin.Engine
	Db      *gorm.DB
	Session gin.HandlerFunc
	Redis   *redis.Client
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
	return nil

}
