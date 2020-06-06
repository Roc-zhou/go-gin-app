package e

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetMsg(c *gin.Context, code int) gin.H {
	lan := c.Request.Header.Get("lan")
	fmt.Println(lan)
	var msg string
	if lan == "en" {
		msg = errorMsgEn[code]
	} else {
		msg = errorMsgCn[code]
	}
	return gin.H{
		"errorMsg": msg,
	}
}
