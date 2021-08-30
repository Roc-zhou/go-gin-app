package user

import (
	"roc/go-gin-app/common"
	"roc/go-gin-app/model"
	"roc/go-gin-app/util"

	"github.com/Roc-zhou/go-util-package/encrypt"
	"github.com/gin-gonic/gin"
)

const SALT = "asdasdqw"

// 注册
func UserRegister(c *gin.Context) {
	json := model.User{}
	c.BindJSON(&json)
	db := common.GetDB()
	user1 := model.User{Name: json.Name, Password: encrypt.Md5(encrypt.Md5(json.Password) + SALT), Phone: json.Phone}
	if result := db.Debug().Create(&user1); result.Error != nil {
		panic(result.Error)
	}
	util.Success(c, nil)
}

// 登陆
func UserLogin(c *gin.Context) {
	util.Success(c, nil)
}
