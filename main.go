package main

import (
	"roc/go-gin-app/common"
	"roc/go-gin-app/routers"

	"github.com/spf13/viper"
)

func main() {
	// 读取配置文件
	common.InitConfig()
	// 数据库
	/*db := common.InitDb()
	defer db.Close()
	var user []model.User
	s := db.Find(&user)*/

	initRoutes := routers.InItRouter()
	port := viper.GetString("server.port")
	initRoutes.Run(":" + port)
}
