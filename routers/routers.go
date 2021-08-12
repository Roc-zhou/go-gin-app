package routers

import (
	"roc/go-gin-app/routers/api/user"
	"roc/go-gin-app/routers/hanlder"

	"github.com/gin-gonic/gin"
)

func InItRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(hanlder.Recover)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/user/register", user.UserRegister)
		apiv1.POST("/user/login", user.UserLogin)
	}

	return r
}
