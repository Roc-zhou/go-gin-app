package model

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UserClaims struct {
	UserId int64
}
