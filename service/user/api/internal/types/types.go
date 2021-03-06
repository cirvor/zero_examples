// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type SearchReq struct {
	Number string `form:"number"`
}

type SearchReply struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
