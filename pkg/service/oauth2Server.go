package service

import "github.com/RangelReale/osin"

type oauth2Server struct {
	server *osin.Server
}

func NewOauth2Server() *oauth2Server {
	// ex.NewTestStorage implements the "osin.Storage" interface
	os := &oauth2Server{
		osin.NewServer(&osin.ServerConfig{
			AuthorizationExpiration:   250,
			AccessExpiration:          3600,
			TokenType:                 "askuy",
			AllowedAuthorizeTypes:     osin.AllowedAuthorizeType{osin.CODE},
			AllowedAccessTypes:        osin.AllowedAccessType{osin.AUTHORIZATION_CODE},
			ErrorStatusCode:           200,
			AllowClientSecretInParams: false,
			AllowGetAccessRequest:     false,
			RetainTokenAfterRefresh:   false,
		}, Storage),
	}
	return os
}

func (os *oauth2Server) GetServer() *osin.Server {
	return os.server
}
