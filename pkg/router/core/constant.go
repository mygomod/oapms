package core

import (
	"encoding/gob"
	"oapms/pkg/model"
)

const AdminSessionKey = "geek/session/admin"
const AdminContextKey = "geek/context/admin"
const RedirectLoginURL = "/user/login"

func init() {
	gob.Register(&model.User{})
}
