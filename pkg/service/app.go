package service

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"time"
)

type app struct{}

func InitApp() *app {
	return &app{}
}

func (*app) Create(name, redirectUri string) (err error) {
	create := model.App{
		Name:        name,
		Secret:      getRandomString(32),
		RedirectUri: redirectUri,
		CallNo:      0,
		State:       1,
		Ctime:       time.Now().Unix(),
		Utime:       time.Now().Unix(),
	}

	err = model.AppCreate(mus.Db, &create)
	return
}

// 调用次数
func (*app) AddCallNo(tx *gorm.DB, clientId string) (err error) {
	_, err = model.AppInfoX(tx, model.Conds{
		"client_id": clientId,
	})
	if err != nil {
		return
	}

	err = tx.Model(model.App{}).Where("client_id = ?", clientId).Updates(model.Ups{
		"call_no": gorm.Expr("call_no+?", 1),
	}).Error
	return
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
