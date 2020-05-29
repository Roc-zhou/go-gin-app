package routers

import (
	"roc/go-gin-app/util"

	"github.com/gin-gonic/gin"
)

func InItRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/demo", func(c *gin.Context) {
			util.Success(c, gin.H{"name": "123"})
			return
		})
	}

	return r
}
