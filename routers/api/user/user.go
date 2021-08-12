package user

import (
	"fmt"
	"roc/go-gin-app/common"
	"roc/go-gin-app/util"

	"github.com/gin-gonic/gin"
)

type User struct {
	common.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// 注册
func UserRegister(c *gin.Context) {
	json := User{}
	c.BindJSON(&json)

	db := common.GetDB()
	user1 := User{Name: json.Name, Password: json.Password, Phone: json.Phone}
	fmt.Println(user1)
	db.Debug().Create(&user1)

	// var user User
	// s := db.Where("name = ?", "111").First(&user)
	// fmt.Println(s.RowsAffected)

	util.Success(c, nil)
}

// 登陆
func UserLogin(c *gin.Context) {
	util.Success(c, nil)
}
