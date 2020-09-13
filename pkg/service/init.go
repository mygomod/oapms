package service

var (
	Storage      *storage
	Oauth2Server *oauth2Server
	User         *user
	App          *app
	Authorize    *authorize
	Pms          *pms
	PmsMenu      *pmsMenu
)

func Init() error {
	Storage = NewStorage()
	Oauth2Server = NewOauth2Server()
	User = InitUser()
	App = InitApp()
	Authorize = InitAuthorize()
	Pms = InitPms()
	PmsMenu = InitPmsMenu()
	InitMailer()
	return nil

}
