package model

type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginInfo struct {
	Id    int
	Token string
}
