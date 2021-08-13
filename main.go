package main

import (
	"roc/go-gin-app/common"
	"roc/go-gin-app/routers"

	"github.com/spf13/viper"
)

func main() {
	common.Init()
	// exp := time.Now().Add(time.Duration(10) * time.Second).Unix() // 当前时间增加24小时
	// fmt.Println(exp)
	// customClaims := &module.CustomClaims{
	// 	UserId: 1,
	// 	StandardClaims: jwt.StandardClaims{
	// 签发时间
	// 		IssuedAt:  time.Now().Unix(),
	// 		ExpiresAt: exp, // 过期时间，
	// 	},
	// }
	// token, err := customClaims.GetToken()
	// if err != nil {
	// 	fmt.Printf("create token error: %s", err)
	// }
	// fmt.Println(token)

	// s, err := module.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTYyODU4NjQ1MX0.-KB4kTUpQ3RaevFkvxY7lz7todkX-57xBSb7UE6sx4I")
	// if err != nil {
	// 	fmt.Printf("parse token error: %s", err)
	// }
	// fmt.Println(s)

	initRoutes := routers.InItRouter()
	port := viper.GetString("server.port")
	initRoutes.Run(":" + port)

}
