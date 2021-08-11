package user

import (
	"roc/go-gin-app/util"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 注册
func UserRegister(c *gin.Context) {
	json := User{}
	c.BindJSON(&json)
	util.Success(c, nil)
}

// 登陆
func UserLogin(c *gin.Context) {
	util.Success(c, nil)
}
