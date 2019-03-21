package models

//easyjson:json
type UserShort struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}

//easyjson:json
type UserDetail struct {
	UserShort
	Nickname string `json:"nickname"`
}

//easyjson:json
type UserDetailList []UserDetail
