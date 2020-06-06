package util

import (
	"net/http"

	"roc/go-gin-app/app/e"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode, code int, data gin.H, status string) {
	c.JSON(httpCode, gin.H{
		"code":   code,
		"status": status,
		"body":   data,
	})
}

func Success(c *gin.Context, data gin.H) {
	Response(c, http.StatusOK, e.SUCCESS, data, "ok")
}
func Fail(c *gin.Context) {
	Response(c, http.StatusInternalServerError, e.ERROR, e.GetMsg(c, e.ERROR), "fail")
}
