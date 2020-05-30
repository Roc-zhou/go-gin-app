package util

import (
	"net/http"

	"roc/go-gin-app/app/e"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode, code int, data gin.H, msg string) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Success(c *gin.Context, data gin.H) {
	Response(c, http.StatusOK, e.SUCCESS, data, e.GetMsg(e.SUCCESS))
}
func Fail(c *gin.Context, data gin.H) {
	Response(c, http.StatusInternalServerError, e.ERROR, data, e.GetMsg(e.ERROR))
}
