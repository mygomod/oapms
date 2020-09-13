package model

type UserInfoResp struct {
	Uid        int    `json:"uid"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Avatar     string `json:"avatar"`
	Country    string `json:"country"`
	Geographic struct {
		City struct {
			Key   string `json:"key"`
			Label string `json:"label"`
		} `json:"city"`
		Province struct {
			Key   string `json:"key"`
			Label string `json:"label"`
		} `json:"province"`
	} `json:"geographic"`
	Group       string `json:"group"`
	NotifyCount int    `json:"notifyCount"`
	Phone       string `json:"phone"`
	Signature   string `json:"signature"`
	Tags        []struct {
		Key   string `json:"key"`
		Label string `json:"label"`
	} `json:"tags"`
}

type RespOauthLogin struct {
	CurrentAuthority string `json:"currentAuthority"`
}

type ReqOauthLogin struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Mfa      string `json:"mfa"`
	Params   struct {
		Redirect     string `json:"redirect"`     // redirect by ant design
		RedirectUri  string `json:"redirect_uri"` // redirect by backend
		ClientId     string `json:"client_id"`
		ResponseType string `json:"response_type"`
	} `form:"params" binding:"required"`
}
