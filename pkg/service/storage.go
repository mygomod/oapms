package service

import (
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"oapms/pkg/model"
	"oapms/pkg/mus"
	"time"
)

// storage implements interface "github.com/RangelReale/osin".storage and interface "github.com/felipeweb/osin-mysql/storage".storage
type storage struct {
}

// New returns a new mysql storage instance.
func NewStorage() *storage {
	return &storage{}
}

// Clone the storage if needed. For example, using mgo, you can clone the session with session.Clone
// to avoid concurrent access problems.
// This is to avoid cloning the connection at each method access.
// Can return itself if not a problem.
func (s *storage) Clone() osin.Storage {
	return s
}

// Close the resources the storage potentially holds (using Clone for example)
func (s *storage) Close() {
}

// GetClient loads the client by id
func (s *storage) GetClient(clientId string) (client osin.Client, err error) {
	app, err := model.AppInfoX(mus.Db, model.Conds{"client_id": clientId})
	if err != nil {
		return
	}
	c := osin.DefaultClient{
		Id:          app.ClientId,
		Secret:      app.Secret,
		RedirectUri: app.RedirectUri,
		UserData:    app.Extra,
	}
	return &c, nil
}

// UpdateClient updates the client (identified by it's id) and replaces the values with the values of client.
func (s *storage) UpdateClient(c osin.Client) error {
	err := model.AppUpdateX(mus.Db, model.Conds{"client_id": c.GetId()}, model.Ups{
		"secret":       c.GetSecret(),
		"redirect_uri": c.GetRedirectUri(),
		"extra":        c.GetUserData(),
	})
	return err
}

// CreateClient stores the client in the database and returns an error, if something went wrong.
func (s *storage) CreateClient(c osin.Client) error {

	create := model.App{
		ClientId:    c.GetId(),
		Secret:      c.GetSecret(),
		RedirectUri: c.GetRedirectUri(),
		Extra:       cast.ToString(c.GetUserData()),
	}

	err := model.AppCreate(mus.Db, &create)
	return err
}

// RemoveClient removes a client (identified by id) from the database. Returns an error if something went wrong.
func (s *storage) RemoveClient(clientId string) (err error) {
	err = model.AppDeleteX(mus.Db, model.Conds{"client_id": clientId})
	return
}

// SaveAuthorize saves authorize data.
func (s *storage) SaveAuthorize(data *osin.AuthorizeData) (err error) {
	obj := model.Authorize{
		Client:      data.Client.GetId(),
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectUri: data.RedirectUri,
		State:       data.State,
		Ctime:       data.CreatedAt.Unix(),
		Extra:       cast.ToString(data.UserData),
	}

	tx := mus.Db.Begin()

	err = model.AuthorizeCreate(tx, &obj)
	if err != nil {
		tx.Rollback()
		return
	}

	err = s.AddExpireAtData(tx, data.Code, data.ExpireAt())
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (s *storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	var data osin.AuthorizeData

	info, err := model.AuthorizeInfoX(mus.Db, model.Conds{"code": code})
	if err != nil {
		return nil, err
	}

	data = osin.AuthorizeData{
		Code:        info.Code,
		ExpiresIn:   info.ExpiresIn,
		Scope:       info.Scope,
		RedirectUri: info.RedirectUri,
		State:       info.State,
		CreatedAt:   time.Unix(info.Ctime, 0),
		UserData:    info.Extra,
	}
	c, err := s.GetClient(info.Client)
	if err != nil {
		return nil, err
	}

	if data.ExpireAt().Before(time.Now()) {
		return nil, errors.New(fmt.Sprintf("Token expired at %s.", data.ExpireAt().String()))
	}

	data.Client = c
	return &data, nil
}

// RemoveAuthorize revokes or deletes the authorization code.
func (s *storage) RemoveAuthorize(code string) (err error) {
	err = model.AuthorizeDeleteX(mus.Db, model.Conds{"code": code})
	if err != nil {
		return
	}

	if err = s.RemoveExpireAtData(code); err != nil {
		return err
	}
	return nil
}

// SaveAccess writes AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (s *storage) SaveAccess(data *osin.AccessData) (err error) {
	prev := ""
	authorizeData := &osin.AuthorizeData{}

	if data.AccessData != nil {
		prev = data.AccessData.AccessToken
	}

	if data.AuthorizeData != nil {
		authorizeData = data.AuthorizeData
	}

	extra := cast.ToString(data.UserData)

	tx := mus.Db.Begin()

	if data.RefreshToken != "" {
		if err := s.saveRefresh(tx, data.RefreshToken, data.AccessToken); err != nil {
			tx.Rollback()
			return err
		}
	}

	if data.Client == nil {
		return errors.New("data.Client must not be nil")
	}

	obj := model.Access{
		Client:       data.Client.GetId(),
		Authorize:    authorizeData.Code,
		Previous:     prev,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
		ExpiresIn:    int(data.ExpiresIn),
		Scope:        data.Scope,
		RedirectUri:  data.RedirectUri,
		Ctime:        data.CreatedAt.Unix(),
		Extra:        extra,
	}

	err = model.AccessCreate(tx, &obj)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = App.AddCallNo(tx, data.Client.GetId())
	if err != nil {
		tx.Rollback()
		return
	}
	err = s.AddExpireAtData(tx, data.AccessToken, data.ExpireAt())
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return nil
}

// LoadAccess retrieves access data by token. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *storage) LoadAccess(code string) (*osin.AccessData, error) {
	var result osin.AccessData

	info, err := model.AccessInfoX(mus.Db, model.Conds{"access_token": code})
	if err != nil {
		return nil, err
	}

	result.AccessToken = info.AccessToken
	result.RefreshToken = info.RefreshToken
	result.ExpiresIn = int32(info.ExpiresIn)
	result.Scope = info.Scope
	result.RedirectUri = info.RedirectUri
	result.CreatedAt = time.Unix(info.Ctime, 0)
	result.UserData = info.Extra
	client, err := s.GetClient(info.Client)
	if err != nil {
		return nil, err
	}

	result.Client = client
	result.AuthorizeData, _ = s.LoadAuthorize(info.Authorize)
	prevAccess, _ := s.LoadAccess(info.Previous)
	result.AccessData = prevAccess
	return &result, nil
}

// RemoveAccess revokes or deletes an AccessData.
func (s *storage) RemoveAccess(code string) (err error) {
	err = model.AccessDeleteX(mus.Db, model.Conds{"access_token": code})
	if err != nil {
		return
	}
	err = s.RemoveExpireAtData(code)
	return
}

// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *storage) LoadRefresh(code string) (*osin.AccessData, error) {
	info, err := model.RefreshInfoX(mus.Db, model.Conds{"token": code})
	if err != nil {
		return nil, err
	}
	return s.LoadAccess(info.Access)
}

// RemoveRefresh revokes or deletes refresh AccessData.
func (s *storage) RemoveRefresh(code string) (err error) {
	err = model.RefreshDeleteX(mus.Db, model.Conds{"token": code})
	return
}

// CreateClientWithInformation Makes easy to create a osin.DefaultClient
func (s *storage) CreateClientWithInformation(id string, secret string, redirectURI string, userData interface{}) osin.Client {
	return &osin.DefaultClient{
		Id:          id,
		Secret:      secret,
		RedirectUri: redirectURI,
		UserData:    userData,
	}
}

func (s *storage) saveRefresh(tx *gorm.DB, refresh, access string) (err error) {
	obj := model.Refresh{
		Token:  refresh,
		Access: access,
	}

	err = model.RefreshCreate(tx, &obj)
	return
}

// AddExpireAtData add info in expires table
func (s *storage) AddExpireAtData(tx *gorm.DB, code string, expireAt time.Time) (err error) {
	obj := model.Expires{
		Token:     code,
		ExpiresAt: expireAt.Unix(),
	}
	err = model.ExpiresCreate(tx, &obj)
	return
}

// RemoveExpireAtData remove info in expires table
func (s *storage) RemoveExpireAtData(code string) (err error) {
	err = model.ExpiresDeleteX(mus.Db, model.Conds{"token": code})
	return
}
