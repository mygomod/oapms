package service

import (
	"crypto/rand"
	"encoding/base32"
	"go.uber.org/zap"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"time"
)

type user struct{}

func InitUser() *user {
	return &user{}
}

func (*user) Create(nickname string, pwd string, email string, state int, ip string) (err error) {
	var pwdHash string
	pwdHash, err = Authorize.Hash(pwd)
	if err != nil {
		return
	}

	user := model.User{
		Nickname:    nickname,
		Password:    pwdHash,
		Email:       email,
		State:       state,
		Ctime:       time.Now().Unix(),
		Utime:       time.Now().Unix(),
		LastLoginIp: ip,
	}
	err = model.UserCreate(mus.Db, &user)
	return
}

func (*user) Update(uid int, nickname string, pwd string) (err error) {
	var pwdHash string
	pwdHash, err = Authorize.Hash(pwd)
	if err != nil {
		mus.Logger.Debug("update user hash error", zap.String("err", err.Error()))
		return
	}

	err = model.UserUpdateX(mus.Db, model.Conds{"uid": uid}, model.Ups{
		"nickname": nickname,
		"password": pwdHash,
	})

	if err != nil {
		mus.Logger.Error("update user create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

func (*user) GetUserByNicknamePwd(nickname string, pwd string, clientIp string) (userInfo model.User, err error) {
	userInfo, err = model.UserInfoX(mus.Db, model.Conds{"nickname": nickname})
	if err != nil {
		mus.Logger.Error("GetUserByNicknamePwd ERROR", zap.Error(err))
		return
	}

	err = Authorize.Verify(userInfo.Password, pwd)
	if err != nil {
		mus.Logger.Error("verify error1", zap.Error(err))
		return
	}
	return
}

// {"name":"Serati Ma","avatar":"https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png","userid":"00000001","email":"antdesign@alipay.com","signature":"海纳百川，有容乃大","title":"交互专家","group":"蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED","tags":[{"key":"0","label":"很有想法的"},{"key":"1","label":"专注设计"},{"key":"2","label":"辣~"},{"key":"3","label":"大长腿"},{"key":"4","label":"川妹子"},{"key":"5","label":"海纳百川"}],"notifyCount":12,"country":"China","geographic":{"province":{"label":"浙江省","key":"330000"},"city":{"label":"杭州市","key":"330100"}},"address":"西湖区工专路 77 号","phone":"0752-268888888"}
func (*user) Info(uid interface{}) (resp model.UserInfoResp, err error) {

	info, err := model.UserInfoX(mus.Db, model.Conds{"uid": uid})
	if err != nil {
		mus.Logger.Error("getUserInfoErr", zap.Error(err))
		return
	}

	resp.Name = info.Nickname
	resp.Uid = info.Uid
	resp.Avatar = "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png"
	resp.Tags = make([]struct {
		Key   string `json:"key"`
		Label string `json:"label"`
	}, 0)
	resp.Tags = append(resp.Tags, struct {
		Key   string `json:"key"`
		Label string `json:"label"`
	}{
		Key:   "0",
		Label: "很有想法的",
	})
	resp.Address = "湖北武汉"
	return
}

func (*user) GetSecret(uid int) (userSecret model.UserSecret, err error) {
	user := model.User{}
	mus.Db.Select("nickname").Where("uid = ?", uid).Find(&user)
	mus.Db.Select("uid,secret,is_bind").Where("uid = ?", uid).Find(&userSecret)

	if userSecret.Uid > 0 {
		userSecret.Nickname = user.Nickname
		return
	}

	secret := make([]byte, 10)
	_, err = rand.Read(secret)
	if err != nil {
		return
	}
	userSecret = model.UserSecret{
		Uid:    uid,
		Secret: base32.StdEncoding.EncodeToString(secret),
		IsBind: 0,
	}
	err = model.UserSecretCreate(mus.Db, &userSecret)
	userSecret.Nickname = user.Nickname
	return
}
