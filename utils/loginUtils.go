package utils

type Credentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}
