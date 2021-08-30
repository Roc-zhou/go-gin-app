package model

type User struct {
	Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
