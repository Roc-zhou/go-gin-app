package main

import (
	"roc/go-gin-app/common"

	"github.com/gin-gonic/gin"
)

func main() {
	// 读取配置文件
	common.InitConfig()
	// 数据库
	db := common.InitDb()
	defer db.Close()

	r := gin.Default()
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Sex      int    `json:"sex"`
	}

	var user []User
	db.Find(&user)

	var adduser = User{Name: "nihao11", Sex: 2}
	db.Create(&adduser)

	//rows, err := db.Raw("SELECT id,name,sex FROM user").Rows()
	//if err != nil {
	//	panic("rows error : " + err.Error())
	//}
	//defer rows.Close()
	//var data = map[string]interface{}{"name": 12}
	//for rows.Next() {
	//	//rows.Scan(&result)
	//	//data = append(data, rows.Scan(&user))
	//	fmt.Println(rows.Scan(&user.Id, &user.Name, &user.Sex))
	//}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": user,
		})
	})
	r.Run(":3386")
}
